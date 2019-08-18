package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/julienschmidt/httprouter"
)

// Handler defines Handler instance and its dependencies.
type Handler struct {
	conf   Config
	model  Model
	mailer Mailer
}

// NewHandler returns a new Handler instance.
func NewHandler(c Config, model Model, mailer Mailer) *Handler {
	return &Handler{
		conf:   c,
		model:  model,
		mailer: mailer,
	}
}

// Health handles endpoint /health.
func (h *Handler) Health(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Auth service is up and running")
}

// User handles endpoint /user/find/:id
// If user is found, return status code 200 with response body 1
// If user is not found, return status code 200 with response body 0
func (h *Handler) User(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	intercept(w, r)
	param := ps.ByName("id")
	if len(param) == 0 {
		respond400(w)
		return
	}
	user := &User{}
	if err := h.model.GetUser(strings.ToLower(param), user); err != nil {
		if !h.model.IsErrNotFound(err) {
			respond500(w, err)
			return
		}
		respond200(w, 0)
		return
	}
	respond200(w, 1)
	return
}

func (h *Handler) Forget(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	intercept(w, r)
	param := ps.ByName("id")
	if len(param) == 0 {
		respond400(w)
		return
	}

	email := strings.ToLower(param)
	code := encode(6)
	exp := time.Now().Add(time.Minute * time.Duration(h.conf.VerificationCodeExpiryMinutes))
	if err := h.model.UpdateActiveCode(email, code, exp); err != nil {
		respond500(w, err)
		return
	}

	go h.mailer.SendVerificationEmail(email, code)

	res := struct {
		Action string `json:"action"`
	}{
		"forget",
	}
	respond200(w, res)
}

func (h *Handler) SetPassword(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	intercept(w, r)
	obj := struct {
		Email    string `json:"email"`
		Token    string `json:"token"`
		Password string `json:"password"`
	}{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&obj); err != nil {
		respond500(w, err)
		return
	}

	if err := h.model.UpdatePassword(obj.Email, obj.Token, getMD5Hash(obj.Password)); err != nil {
		respond500(w, err)
		return
	}

	respond200(w, 1)
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

	user := &User{}
	if err := h.model.GetUser(obj.Email, user); err != nil {
		respond500(w, err)
		return
	}
	if user == nil {
		respond404(w)
		return
	}
	if !user.IsActive {
		code := encode(6)
		exp := time.Now().Add(time.Minute * time.Duration(h.conf.VerificationCodeExpiryMinutes))
		if err := h.model.UpdateActiveCode(obj.Email, code, exp); err != nil {
			respond500(w, err)
			return
		}

		go h.mailer.SendVerificationEmail(user.Email, code)

		res := struct {
			Action string `json:"action"`
		}{
			"verify",
		}
		respond200(w, res)
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
	user.IsActive = false
	user.ActivationCode = code
	user.ActivationCodeExpiry = time.Now().Add(time.Minute * time.Duration(h.conf.VerificationCodeExpiryMinutes))

	if _, err := h.model.InsertUser(user); err != nil {
		respond500(w, err)
		return
	}

	go h.mailer.SendVerificationEmail(user.Email, code)

	respond200(w, 1)
	return
}

// Verify handles endpoint /user/verify.
func (h *Handler) Verify(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	intercept(w, r)

	obj := struct {
		Email          string `json:"email"`
		ActivationCode string `json:"code"`
	}{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&obj); err != nil {
		respond500(w, err)
		return
	}

	user := &User{}
	if err := h.model.VerifyUser(obj.Email, obj.ActivationCode, user); err != nil {
		respond500(w, err)
		return
	}

	if user != nil {
		// mark user active
		if err := h.model.ActivateUser(user.Email); err != nil {
			respond500(w, err)
			return
		}
		res := struct {
			Email string `json:"email"`
			Token string `json:"token"`
		}{
			user.Email,
			user.Token,
		}
		respond200(w, res)
		return
	}

	respond404(w)
	return
}

// Auth handles endpoint /authenticate/:token
func (h *Handler) Authenticate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// TODO: implement specific interceptor for cors control
	intercept(w, r)
	param := ps.ByName("token")
	if len(param) == 0 {
		respond404(w)
		return
	}

	user := &User{}
	if err := h.model.FindUserByTok(param, user); err != nil {
		respond400(w)
		return
	}

	// TODO: TouchTok should return new token expiry.
	if err := h.model.TouchTok(user.Email); err != nil {
		respond500(w, err)
		return
	}

	respond200(w, user)
	return
}

// Profile handles endpoint /profile/:id
func (h *Handler) Profile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	intercept(w, r)
	param := ps.ByName("id")
	if len(param) == 0 {
		respond404(w)
		return
	}

	user := &User{}
	if err := h.model.GetUser(param, user); err != nil {
		if !h.model.IsErrNotFound(err) {
			respond500(w, err)
			return
		}
		respond404(w)
	}

	profile := struct {
		Email    string    `bson:"email"`
		Timezone int       `bson:"timezone"`
		IsActive bool      `bson:"is_active"`
		Created  time.Time `bson:"created,omitempty"`
		Modified time.Time `bson:"modified,omitempty"`
	}{
		user.Email,
		user.Timezone,
		user.IsActive,
		user.Created,
		user.Modified,
	}

	respond200(w, profile)
	return
}
