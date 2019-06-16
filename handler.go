package auth

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/jacygao/mail"
	"github.com/julienschmidt/httprouter"
)

type Handler struct {
	model  Model
	mailer *mail.Service
}

func NewHandler(model Model, mailer *mail.Service) *Handler {
	return &Handler{
		model:  model,
		mailer: mailer,
	}
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Auth service is up and running")
}

func (h *Handler) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

// Register handles endpoint /user/register
func (h *Handler) Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	intercept(w, r)
	obj := struct {
		Email string `json:"email"`
	}{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&obj); err != nil {
		respond500(w, err)
	}

	var user *User
	if err := h.model.GetUser(obj.Email, user); err != nil {
		if user == nil {
			// send registration email to user
		}
		respond500(w, err)
	}

	if err := respond200(w, 1); err != nil {
		respond500(w, err)
	}
}

func (h *Handler) register(email string) error {
	param := hex.EncodeToString([]byte(email))
	// send email with url

	return nil
}

// login handles endpoint /user/login
func (h *Handler) Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	intercept(w, r)
	obj := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&obj); err != nil {
		respond500(w, err)
	}

	var user *User
	if err := h.model.GetUser(obj.Email, user); err != nil {
		respond500(w, err)
	}
	if user == nil {
		respond404(w)
	}
	if user.Password == getMD5Hash(obj.Password) {
		res := struct {
			Token string `json:"token"`
		}{
			user.Token,
		}
		respond200(w, res)
	}
	respond401(w)
}

// Signup handles endpoint /user/signup
func (h *Handler) Signup(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	intercept(w, r)

	obj := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Timezone int    `json:"timezone"`
	}{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&obj); err != nil {
		respond500(w, err)
	}

	user := &User{}
	user.Email = strings.ToLower(obj.Email)
	user.Password = getMD5Hash(obj.Password)
	user.Timezone = obj.Timezone

	if err := h.model.InsertUser(user); err != nil {
		respond500(w, err)
	}

	respond200(w, "")
}

func intercept(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

func respond200(w http.ResponseWriter, res interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(res); err != nil {
		return err
	}

	return nil
}

func respond500(w http.ResponseWriter, err error) {
	log.Println(err.Error())
	w.WriteHeader(http.StatusInternalServerError)
}

func respond204(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

func respond401(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
}

func respond404(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
}

func getMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
