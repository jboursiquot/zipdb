package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/jboursiquot/zipdb"
)

type handler struct {
	locations map[string]zipdb.Location
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	zip := chi.URLParam(r, "zip")
	if loc, found := h.locations[zip]; found {
		render.JSON(w, r, loc)
		return
	}
	http.Error(w, http.StatusText(404), 404)
}

func main() {
	locations, err := zipdb.LoadLocations(cfg.DataFile)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	h := &handler{locations: locations}

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
