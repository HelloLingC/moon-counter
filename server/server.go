package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/HelloLingC/moon-counter/common"
	"github.com/HelloLingC/moon-counter/database"
)

type Server struct {
	Config *common.Config
	DB     database.IDatabase
}

func NewInstance(config *common.Config, db database.IDatabase) *Server {
	return &Server{
		Config: config,
		DB:     db,
	}
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
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
	identifier := r.URL.Query().Get("id")
	if identifier == "" {
		http.Error(w, "missing identifier", http.StatusBadRequest)
		return
	}
	count, err := s.DB.AddCounter(identifier)
	if err != nil {
		common.SilentError("SQL err when adding:", err)
		http.Error(w, "DB Error", http.StatusServiceUnavailable)
		return
	}
	svg := BuildCounterImg(fmt.Sprintf("%d", count))
	w.Header().Set("Content-Type", "image/svg+xml")
	fmt.Fprint(w, svg)
}

func (s Server) textHndl(w http.ResponseWriter, r *http.Request) {
	origin := r.Header.Get("Origin")
	rmOrigin := checkOrigin(w, origin, s.Config.Hostnames)
	if rmOrigin == "" {
		// Didn't pass the origin check
		http.Error(w, "access blocked", http.StatusForbidden)
		return
	}
	count, err := s.DB.AddCounter(rmOrigin)
	if err != nil {
		common.SilentError("SQL err when adding", err)
		http.Error(w, "DB error", http.StatusServiceUnavailable)
		return
	}
	// [S] Do NOT send all the allowed origins to the client
	w.Header().Set("Access-Control-Allow-Origin", origin)
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "%d", count)
}

func (s Server) Start() {
	log.Printf("Moon Counter starts running at localhost:%d", s.Config.Port)

	tHndl := http.HandlerFunc(s.textHndl)

	http.HandleFunc("/moon-counter/js", s.jsTextHndl)
	http.Handle("/counter/text", corsMiddleware(tHndl))
	http.HandleFunc("/counter/img", s.imgHndl)

	err := http.ListenAndServe(fmt.Sprintf(":%d", s.Config.Port), nil)
	if err != nil {
		log.Fatal("Error starting server: ", err.Error())
	}
}
