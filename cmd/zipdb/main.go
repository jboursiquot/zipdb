package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/jboursiquot/zipdb"
)

type handler struct {
	db *zipdb.DB
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	zip := chi.URLParam(r, "zip")
	loc, err := h.db.Find(zip)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, http.StatusText(500), 500)
		return
	}
	render.JSON(w, r, loc)
}

func main() {
	db, err := zipdb.NewDB(cfg.DBFile)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	h := &handler{db: db}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/{zip:[0-9]+}", h.ServeHTTP)

	hostPort := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)

	log.Info("Listening on %s...", "addr", hostPort)
	if err := http.ListenAndServe(hostPort, r); err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
}
