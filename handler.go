package zipdb

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type Handler struct {
	db *DB
}

func NewHandler(db *DB) *Handler {
	return &Handler{
		db: db,
	}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var loc Location
	if err := json.Unmarshal(data, &loc); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if loc, err := h.db.Find(loc.Zip); err == nil && loc != nil {
		http.Error(w, "already exists", http.StatusConflict)
		return
	}

	if err := h.db.Upsert(&loc); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) Read(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	zip := chi.URLParam(r, "zip")
	if _, err := h.db.Find(zip); err != nil {
		if strings.Contains(err.Error(), "record not found") {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var loc Location
	if err := json.Unmarshal(data, &loc); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	loc.ID = zip

	if err := h.db.Upsert(&loc); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	zip := chi.URLParam(r, "zip")
	loc, err := h.db.Find(zip)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if err := h.db.Delete(loc); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
