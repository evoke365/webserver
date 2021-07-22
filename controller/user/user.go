package user

import (
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/evoke365/webserver/controller/internal/responder"
	"github.com/evoke365/webserver/pkg/mailer"
	"github.com/evoke365/webserver/restapi/operations/user"
	"github.com/evoke365/webserver/store"
	"github.com/evoke365/webserver/store/data"
	"github.com/go-openapi/runtime/middleware"
)

// Config defines madetory configuration values to initialise a controller instance.
type Config struct {
	VerificationCodeExpiryMinutes int
}

// DefaultConfig returns a config object with default/recommended values.
func DefaultConfig() Config {
	return Config{
		VerificationCodeExpiryMinutes: 15,
	}
}

// Controller defines HTTP handlers.
type Controller struct {
	conf   Config
	store  store.DB
	mailer *mailer.Client
}

// NewController returns a new instance of Controller.
func NewController(s store.DB, m *mailer.Client, c Config) *Controller {
	return &Controller{
		conf:   c,
		store:  s,
		mailer: m,
	}
}

// Signup implements the logic of HTTP Handler for /signup.
func (c *Controller) Signup(req *user.SignupUserParams) middleware.Responder {
	code := randCode(6)

	user := &data.User{
		Email:                strings.ToLower(req.Body.Email),
		Password:             hashMD5(req.Body.Password),
		Timezone:             req.Body.Timezone,
		IsActive:             false,
		ActivationCode:       code,
		ActivationCodeExpiry: time.Now().Add(time.Minute * time.Duration(c.conf.VerificationCodeExpiryMinutes)),
	}

	if _, err := c.store.InsertUser(user); err != nil {
		log.Println(err.Error())
		return responder.DefaultServerError()
	}

	// if err := h.callback.OnSignup(user); err != nil {
	// 	respond500(w, err)
	// 	return
	// }

	// if err := h.callback.OnVerify(user.Email, user.ActivationCode); err != nil {
	// 	respond500(w, err)
	// 	return
	// }

	return responder.DefaultOK()
}

// FindUser defines the logic of finding a user from a data store.
func (c *Controller) FindUser(req *user.FindUserParams) middleware.Responder {
	if len(req.ID) == 0 {
		return responder.DefaultBadRequest()
	}
	user := &data.User{}
	if err := c.store.GetUser(strings.ToLower(req.ID), user); err != nil {
		if !c.store.IsErrNotFound(err) {
			log.Println(err.Error())
			return responder.DefaultServerError()
		}
		return responder.DefaultNoContent()
	}
	return responder.DefaultOK()
}

// ForgetPassword implements the HTTP handler logic for /forget
func (c *Controller) ForgetPassword(req *user.ForgetPasswordParams) middleware.Responder {
	if len(req.Body.ID) == 0 {
		return responder.DefaultBadRequest()
	}

	email := strings.ToLower(req.Body.ID)
	code := randCode(6)
	exp := time.Now().Add(time.Minute * time.Duration(c.conf.VerificationCodeExpiryMinutes))
	if _, err := c.store.UpdateActiveCode(email, code, exp); err != nil {
		log.Println(err.Error())
		return responder.DefaultServerError()
	}

	// if err := h.callback.OnVerify(user.Email, user.ActivationCode); err != nil {
	// 	respond500(w, err)
	// 	return
	// }
	return responder.DefaultOK()
}

// LoginUser implements the HTTP handler logic for /login
func (c *Controller) LoginUser(req *user.LoginUserParams) middleware.Responder {
	user := &data.User{}
	if err := c.store.GetUser(req.Body.Email, user); err != nil {
		log.Println(err.Error())
		return responder.DefaultServerError()
	}
	if user == nil {
		return responder.DefaultBadRequest()
	}

	if !user.IsActive {
		code := randCode(6)
		exp := time.Now().Add(time.Minute * time.Duration(c.conf.VerificationCodeExpiryMinutes))
		if _, err := c.store.UpdateActiveCode(req.Body.Email, code, exp); err != nil {
			log.Println(err.Error())
			return responder.DefaultServerError()
		}

		// if err := h.callback.OnVerify(user.Email, user.ActivationCode); err != nil {
		// 	log.Println(err.Error())
		// 	return responder.DefaultServerError()
		// }

		return responder.DefaultOK()
	}
	if user.Password == hashMD5(req.Body.Password) {
		res := LoginResponse{
			Token: user.Token,
		}

		b, err := json.Marshal(res)
		if err != nil {
			log.Println(err.Error())
			return responder.DefaultServerError()
		}
		return responder.DefaultOK().WithResponse(b)
	}
	return responder.DefaultUnauthorised()
}

// NewPassword implements the HTTP handler logic for /newpassword
func (c *Controller) NewPassword(req *user.NewPasswordParams) middleware.Responder {
	if err := c.store.UpdatePassword(req.Body.Email, req.Body.Token, hashMD5(req.Body.Password)); err != nil {
		log.Println(err.Error())
		return responder.DefaultServerError()
	}
	return responder.DefaultOK()
}

// VerifyUser implements the HTTP handler logic for /verify
func (c *Controller) VerifyUser(req *user.VerifyUserParams) middleware.Responder {
	user := &data.User{}
	if err := c.store.VerifyUser(req.Body.Email, req.Body.Code, user); err != nil {
		log.Println(err.Error())
		return responder.DefaultServerError()
	}
	if user == nil {
		return responder.DefaultBadRequest()
	}

	// mark user active
	if err := c.store.ActivateUser(user.Email); err != nil {
		log.Println(err.Error())
		return responder.DefaultServerError()
	}
	res := VerifyUserResponse{
		user.Email,
		user.Token,
	}
	b, err := json.Marshal(res)
	if err != nil {
		log.Println(err.Error())
		return responder.DefaultServerError()
	}
	return responder.DefaultOK().WithResponse(b)
}
