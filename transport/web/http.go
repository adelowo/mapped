package web

import (
	"net/http"

	"github.com/adelowo/mapped/config"
	"github.com/adelowo/mapped/link"
	"github.com/apex/log"
	"github.com/go-chi/chi"
)

type (
	Server struct {
		DB     link.Repository
		Cfg    *config.Configuration
		Logger log.Interface
	}
)

func (s *Server) Start() {

	mux := chi.NewMux()

	mux.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		r.Body.Close()

		w.Write([]byte("{'status' : 'OK'}"))
	})

	http.ListenAndServe(s.Cfg.HTTPPort, mux)
}
