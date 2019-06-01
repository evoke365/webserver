package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Handler struct {
	m *Model
}

func NewHandler(s Store) *Handler {
	return &Handler{
		m: NewModel(s),
	}
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Auth service is up and running")
}

func (h *Handler) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}

// Find finds the user by id.
// returns 0 when found
// returns 1 when not found
func (h *Handler) Find(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	h.intercept(w, r)
	email := ps.ByName("name")
	res, err := h.m.GetUser(email)
	if err != nil {
		h.respond500(w)
	}

	if res == nil {
		h.respond204(w)
	}

	if err := h.respond200(w, res); err != nil {
		h.respond500(w)
	}
}

func (h *Handler) intercept(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

func (h *Handler) respond200(w http.ResponseWriter, res interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(res); err != nil {
		return err
	}

	return nil
}

func (h *Handler) respond500(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
}

func (h *Handler) respond204(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}
