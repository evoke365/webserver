package auth

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/jacygao/mail"
	"github.com/julienschmidt/httprouter"
)

// Handler defines Handler instance and its dependencies.
type Handler struct {
	conf   Config
	model  Model
	mailer mail.Mailer
}

// NewHandler returns a new Handler instance.
func NewHandler(model Model, mailer mail.Mailer) *Handler {
	return &Handler{
		model:  model,
		mailer: mailer,
	}
}

// Health handles endpoint /health.
func (h *Handler) Health(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Auth service is up and running")
}

// Redirect handles endpoint /user/auth/:code.
func (h *Handler) Auth(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	intercept(w, r)
	param := ps.ByName("code")
	if len(param) == 0 {
		respond404(w)
		return
	}

	mailBytes, err := hex.DecodeString(param)
	if err != nil {
		respond500(w, err)
		return
	}

	// TODO: check url code expiry
	// TODO: checking and formatting uri string

	res := struct {
		Email string `json:"email"`
	}{
		Email: string(mailBytes),
	}
	respond200(w, res)
}

// User handles endpoint /user/find/:id
func (h *Handler) User(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	intercept(w, r)
	param := ps.ByName("id")
	if len(param) == 0 {
		respond404(w)
		return
	}

	var user *User
	if err := h.model.GetUser(param, user); err != nil {
		if !h.model.IsErrNotFound(err) {
			respond500(w, err)
			return
		}
	}
	if user != nil {
		respond200(w, user.Email)
		return
	}

	respond404(w)
	return

}

// Login handles endpoint /user/login.
func (h *Handler) Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	intercept(w, r)
	obj := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&obj); err != nil {
		respond500(w, err)
		return
	}

	var user *User
	if err := h.model.GetUser(obj.Email, user); err != nil {
		respond500(w, err)
		return
	}
	if user == nil {
		respond404(w)
		return
	}
	if user.Password == getMD5Hash(obj.Password) {
		res := struct {
			Token string `json:"token"`
		}{
			user.Token,
		}
		respond200(w, res)
		return
	}
	respond401(w)
	return
}

// Signup handles endpoint /user/signup.
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
		return
	}

	code := encode(6)

	user := &User{}
	user.Email = strings.ToLower(obj.Email)
	user.Password = getMD5Hash(obj.Password)
	user.Timezone = obj.Timezone
	user.isActive = false
	user.ActivationCode = code
	user.ActivationCodeExpiry = time.Now().Add(time.Minute * 10)

	if _, err := h.model.InsertUser(user); err != nil {
		respond500(w, err)
		return
	}

	if err := h.mailer.Send(user.Email, code); err != nil {
		respond500(w, err)
		return
	}

	respond200(w, "")
	return
}

// TODO: implement config driven CORS.
func intercept(w http.ResponseWriter, req *http.Request) {
	log.Printf("Incoming request: %v", req)
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

func encode(max int) string {
	var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}

	return string(b)
}