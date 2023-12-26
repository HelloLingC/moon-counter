package server

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"sync"
	"time"

	"crypto/rand"
	"html/template"

	"github.com/golang-jwt/jwt/v5"
)

type AdminPanel struct {
	Enabled   bool
	SecretKey []byte
	Mu        sync.Mutex
	tepl      *template.Template
}

func parseJWT(ts string, skey []byte) (*jwt.Token, error) {
	token, err := jwt.Parse(ts, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected alg: %v", token.Header["alg"])
		}
		return skey, nil
	})

	if err != nil {
		return nil, err
	}
	return token, nil
}

func (a *AdminPanel) Register() {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		log.Fatal("Cannot generate pkey:", err)
	}
	a.SecretKey = bytes
	a.tepl = template.Must(template.ParseFS(tpls, "tpl/*.html"))
}

// Middleware to check whether admin is enabled in the config file
// if not enabled, return 404 to disguise as nothing
func AdminMiddleware(next http.Handler, ad *AdminPanel) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// ad.Mu.Lock()
		// defer ad.Mu.Unlock()
		if ad.Enabled {
			next.ServeHTTP(w, r)
			return
		}
		http.NotFound(w, r)
	})
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		next.ServeHTTP(w, r)
	})
}

func (s *Server) AdminUpdateHndl(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) AdminGetHndl(w http.ResponseWriter, r *http.Request) {

}

func (s Server) AuthHndl(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "post only", http.StatusBadRequest)
		return
	}
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "body parsing error", http.StatusBadRequest)
		return
	}
	passphrase := r.Form.Get("pass")
	if passphrase == "" {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	isGuest := s.Config.AdminCfg.GuestLogin != "" && s.Config.AdminCfg.GuestLogin == passphrase
	if passphrase != s.Config.AdminCfg.Passphrase && !isGuest {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	// token will expire after 12 hrs
	exp := time.Now().Add(time.Hour * 12)
	var sub string
	if isGuest {
		sub = "guest"
	} else {
		sub = "admin"
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "moon-counter",
		"sub": sub,
		"exp": exp.Unix(),
	})
	ts, err := token.SignedString(s.AdminEn.SecretKey)
	if err != nil {
		// Todo: error report
		http.Error(w, "cannot sign the token", http.StatusServiceUnavailable)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	// Its weird enough. http package will automatically handle the time format
	// But jwt package doen't, so we need convert time object to unix time above
	cookie := http.Cookie{
		Name:     "token",
		Value:    ts,
		Expires:  exp,
		HttpOnly: true,
		Secure:   false, // no strict HTTPS
		Path:     "/",
	}
	http.SetCookie(w, &cookie)
	fmt.Fprint(w, "ok")
}

func (s Server) AdminHndl(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err == nil {
		// Token cookie exists
		token, err := parseJWT(cookie.Value, s.AdminEn.SecretKey)
		// Since the secret key will be regenerated eveytime whem server starts
		// the authorized users will meet invaild signature error, because
		// the key is different than before. reauth is required
		if err == nil {
			claims, ok := token.Claims.(jwt.MapClaims)
			if ok && token.Valid {
				sub := claims["sub"]
				if sub == "admin" || sub == "guest" {
					counters, err := s.DB.QueryCounter(0, 15)
					if err != nil {
						http.Error(w, "SQL error: "+err.Error(), http.StatusServiceUnavailable)
						return
					}
					s.AdminEn.tepl.ExecuteTemplate(w, "panel.html", map[string]interface{}{"Sub": sub, "Items": counters})
				}
				return
			}
			return
		}
		// Didn't pass auth
	}
	s.AdminEn.tepl.ExecuteTemplate(w, "index.html", map[string]string{"Url": path.Join(s.Config.Host, s.Config.AdminCfg.Path, "/auth")})
}
