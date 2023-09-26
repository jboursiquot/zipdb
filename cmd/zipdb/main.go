package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jboursiquot/zipdb"
)

func main() {
	db, err := zipdb.NewDB(cfg.DBFile)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	h := zipdb.NewHandler(db)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Route("/", func(r chi.Router) {
		r.Route("/{zip:[0-9]+}", func(r chi.Router) {
			r.Get("/", h.Read)      // GET /12345
			r.Put("/", h.Update)    // PUT /12345
			r.Delete("/", h.Delete) // DELETE /12345
		})
	})
	r.Post("/", http.HandlerFunc(h.Create))

	hostPort := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)

	log.Info("Listening on %s...", "addr", hostPort)
	if err := http.ListenAndServe(hostPort, r); err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
}
