package auth

import (
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
func (h *Handler) Find(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}
