package user

import (
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/jacygao/auth/restapi/operations/user"
	"github.com/jacygao/auth/store"
)

type Config struct {
	VerificationCodeExpiryMinutes int
}

func DefaultConfig() Config {
	return Config{
		VerificationCodeExpiryMinutes: 15,
	}
}

// Controller defines HTTP handlers.
type Controller struct {
	conf  Config
	store store.DB
}

// NewController returns a new instance of Controller.
func NewController(s store.DB, c Config) *Controller {
	return &Controller{
		conf:  c,
		store: s,
	}
}

func (c *Controller) Signup(req *user.SignupUserParams) middleware.Responder {
	code := randCode(6)

	user := &store.User{
		Email:                strings.ToLower(req.Body.Email),
		Password:             hashMD5(req.Body.Password),
		Timezone:             req.Body.Timezone,
		IsActive:             false,
		ActivationCode:       code,
		ActivationCodeExpiry: time.Now().Add(time.Minute * time.Duration(c.conf.VerificationCodeExpiryMinutes)),
	}

	if _, err := c.store.InsertUser(user); err != nil {
		log.Println(err.Error())
		return DefaultServerError()
	}

	// if err := h.callback.OnSignup(user); err != nil {
	// 	respond500(w, err)
	// 	return
	// }

	// if err := h.callback.OnVerify(user.Email, user.ActivationCode); err != nil {
	// 	respond500(w, err)
	// 	return
	// }

	return DefaultOK()
}

// FindUser defines the logic of finding a user from a data store.
func (c *Controller) FindUser(req *user.FindUserParams) middleware.Responder {
	if len(req.ID) == 0 {
		return DefaultBadRequest()
	}
	user := &store.User{}
	if err := c.store.GetUser(strings.ToLower(req.ID), user); err != nil {
		if !c.store.IsErrNotFound(err) {
			log.Println(err.Error())
			return DefaultServerError()
		}
		return DefaultNoContent()
	}
	return DefaultOK()
}

func (c *Controller) ForgetPassword(req *user.ForgetPasswordParams) middleware.Responder {
	if len(req.Body.ID) == 0 {
		return DefaultBadRequest()
	}

	email := strings.ToLower(req.Body.ID)
	code := randCode(6)
	exp := time.Now().Add(time.Minute * time.Duration(c.conf.VerificationCodeExpiryMinutes))
	if _, err := c.store.UpdateActiveCode(email, code, exp); err != nil {
		log.Println(err.Error())
		return DefaultServerError()
	}

	// if err := h.callback.OnVerify(user.Email, user.ActivationCode); err != nil {
	// 	respond500(w, err)
	// 	return
	// }
	return DefaultOK()
}

// LoginUser implements the HTTP handler logic for /login
func (c *Controller) LoginUser(req *user.LoginUserParams) middleware.Responder {
	user := &store.User{}
	if err := c.store.GetUser(req.Body.Email, user); err != nil {
		log.Println(err.Error())
		return DefaultServerError()
	}
	if user == nil {
		return DefaultBadRequest()
	}

	if !user.IsActive {
		code := randCode(6)
		exp := time.Now().Add(time.Minute * time.Duration(c.conf.VerificationCodeExpiryMinutes))
		if _, err := c.store.UpdateActiveCode(req.Body.Email, code, exp); err != nil {
			log.Println(err.Error())
			return DefaultServerError()
		}

		// if err := h.callback.OnVerify(user.Email, user.ActivationCode); err != nil {
		// 	log.Println(err.Error())
		// 	return DefaultServerError()
		// }

		return DefaultOK()
	}
	if user.Password == hashMD5(req.Body.Password) {
		res := LoginResponse{
			Token: user.Token,
		}

		b, err := json.Marshal(res)
		if err != nil {
			log.Println(err.Error())
			return DefaultServerError()
		}
		return DefaultOK().WithResponse(b)
	}
	return DefaultUnauthorised()
}

func (c *Controller) NewPassword(req *user.NewPasswordParams) middleware.Responder {
	if err := c.store.UpdatePassword(req.Body.Email, req.Body.Token, hashMD5(req.Body.Password)); err != nil {
		log.Println(err.Error())
		return DefaultServerError()
	}
	return DefaultOK()
}

func (c *Controller) VerifyUser(req *user.VerifyUserParams) middleware.Responder {
	user := &store.User{}
	if err := c.store.VerifyUser(req.Body.Email, req.Body.Code, user); err != nil {
		log.Println(err.Error())
		return DefaultServerError()
	}
	if user == nil {
		return DefaultBadRequest()
	}

	// mark user active
	if err := c.store.ActivateUser(user.Email); err != nil {
		log.Println(err.Error())
		return DefaultServerError()
	}
	res := VerifyUserResponse{
		user.Email,
		user.Token,
	}
	b, err := json.Marshal(res)
	if err != nil {
		log.Println(err.Error())
		return DefaultServerError()
	}
	return DefaultOK().WithResponse(b)
}
