package server

import (
	"fmt"
	"log"
	"net/http"
	"path"

	"github.com/HelloLingC/moon-counter/common"
	"github.com/HelloLingC/moon-counter/database"
)

type Server struct {
	Config  *common.Config
	DB      database.IDatabase
	AdminEn *AdminPanel
}

func NewInstance(config *common.Config, db database.IDatabase) *Server {
	return &Server{
		Config:  config,
		DB:      db,
		AdminEn: &AdminPanel{Enabled: config.AdminCfg.Enabled},
	}
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		next.ServeHTTP(w, r)
	})
}

func (s Server) jsTextHndl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/javascript")
	fmt.Fprintf(w, JS, s.Config.Host)
}

func (s Server) jsImgHndl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/javascript")
	fmt.Fprintf(w, JS_IMG, s.Config.Host)
}

func (s Server) imgHndl(w http.ResponseWriter, r *http.Request) {
	origin := r.Header.Get("Origin")
	if s.Config.Cors && !checkOrigin(w, origin, s.Config.Hostnames) {
		return
	}
	var identifier string
	if origin != "" {
		// Client is using JS to request here
		identifier = origin
	} else {
		identifier = r.URL.Query().Get("id")
		if identifier == "" {
			http.Error(w, "missing identifier", http.StatusBadRequest)
			return
		}
	}
	if len(identifier) > 100 {
		http.Error(w, "exceeding id arg", http.StatusBadRequest)
		return
	}
	count, err := s.DB.AddCounter(identifier)
	if err != nil {
		common.SilentError("SQL err when adding:", err)
		http.Error(w, "DB Error", http.StatusServiceUnavailable)
		return
	}
	// Todo: digits customization
	svg := BuildCounterImg(fmt.Sprintf("%d", count))
	w.Header().Set("Access-Control-Allow-Origin", origin)
	w.Header().Set("Cache-Control", "max-age=0")
	w.Header().Set("Content-Type", "image/svg+xml")
	fmt.Fprint(w, svg)
}

func (s Server) textHndl(w http.ResponseWriter, r *http.Request) {
	origin := r.Header.Get("Origin")
	if s.Config.Cors && !checkOrigin(w, origin, s.Config.Hostnames) {
		// Didn't pass the origin check
		return
	}
	// Todo; text counter support id argument
	rmOrigin, err := common.StrRemoveProtocol(origin)
	if err != nil {
		http.Error(w, "Invaild origin: missing protocal", http.StatusBadRequest)
		return
	}
	count, err := s.DB.AddCounter(rmOrigin)
	if err != nil {
		common.SilentError("SQL err when adding", err)
		http.Error(w, "DB error", http.StatusServiceUnavailable)
		return
	}
	// Do NOT send all the allowed origins to the client
	w.Header().Set("Access-Control-Allow-Origin", origin)
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "%d", count)
}

func (s Server) Start() {
	log.Printf("MoonCounter starts running at localhost:%d", s.Config.Port)

	tHndl := http.HandlerFunc(s.textHndl)
	iHndl := http.HandlerFunc(s.imgHndl)

	http.HandleFunc("/moon-counter/js", s.jsTextHndl)
	http.HandleFunc("/moon-counter/js/img", s.jsImgHndl)
	http.Handle("/counter/text", corsMiddleware(tHndl))
	http.Handle("/counter/img", corsMiddleware(iHndl))

	if s.Config.AdminCfg.Enabled {
		log.Println("Warning: MoonCounter Admin is enabled")
		if s.Config.AdminCfg.Passphrase == "" {
			log.Fatal("Admin is enabled, but the passphrase is empty")
		}
	}
	adPath := path.Join("/", s.Config.AdminCfg.Path)
	adHndl := http.HandlerFunc(s.AdminHndl)
	adAuthHndl := http.HandlerFunc(s.AuthHndl)

	s.AdminEn.Register()
	http.Handle(adPath, AdminMiddleware(adHndl, s.AdminEn))
	http.Handle(path.Join(adPath, "/auth"), AdminMiddleware(adAuthHndl, s.AdminEn))

	err := http.ListenAndServe(fmt.Sprintf(":%d", s.Config.Port), nil)
	if err != nil {
		log.Fatal("Error starting server: ", err.Error())
	}
}
