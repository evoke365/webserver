package user

import (
	"log"
	"strings"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/jacygao/auth/restapi/operations/user"
	"github.com/jacygao/auth/store"
)

// Controller defines HTTP handlers.
type Controller struct {
	store store.DB
}

// NewController returns a new instance of Controller.
func NewController(s store.DB) *Controller {
	return &Controller{
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
		ActivationCodeExpiry: time.Now().Add(time.Minute * time.Duration(h.conf.VerificationCodeExpiryMinutes)),
	}

	if _, err := c.store.InsertUser(user); err != nil {
		log.Println(err.Error())
		return &DefaultResponderError{}
	}

	// if err := h.callback.OnSignup(user); err != nil {
	// 	respond500(w, err)
	// 	return
	// }

	// if err := h.callback.OnVerify(user.Email, user.ActivationCode); err != nil {
	// 	respond500(w, err)
	// 	return
	// }

	return &DefaultResponderOK{}
}

// FindUser defines the logic of finding a user from a data store.
func (c *Controller) FindUser(req *user.FindUserParams) middleware.Responder {
	if len(req.ID) == 0 {
		return &DefaultResponderBadRequest{}
	}
	user := &store.User{}
	if err := c.store.GetUser(strings.ToLower(req.ID), user); err != nil {
		if !c.store.IsErrNotFound(err) {
			log.Println(err.Error())
			return &DefaultResponderError{}
		}
		return &DefaultResponderNoContent{}
	}
	return &DefaultResponderOK{}
}

func (c *Controller) ForgetPassword(req *user.ForgetPasswordParams) middleware.Responder {
	if len(req.Body.ID) == 0 {
		return &DefaultResponderBadRequest{}
	}

	email := strings.ToLower(param)
	code := encode(6)
	exp := time.Now().Add(time.Minute * time.Duration(h.conf.VerificationCodeExpiryMinutes))
	user, err := h.model.UpdateActiveCode(email, code, exp)
	if err != nil {
		respond500(w, err)
		return
	}

	if err := h.callback.OnVerify(user.Email, user.ActivationCode); err != nil {
		respond500(w, err)
		return
	}

	res := struct {
		Action string `json:"action"`
	}{
		"forget",
	}
	respond200(w, res)
}

func (c *Controller) LoginUser(req *user.LoginUserParams) middleware.Responder {

}

func (c *Controller) NewPassword(req *user.NewPasswordParams) middleware.Responder {

}

func (c *Controller) VerifyUser(req *user.VerifyUserParams) middleware.Responder {

}
