package user

import (
	"log"
	"strings"

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
	if len(req.ID) == 0 {
		return &DefaultResponderBadRequest{}
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

// FindUser defines the logic of finding a user from a data store.
func (c *Controller) FindUser(req *user.FindUserParams) middleware.Responder {
	if len(req.ID) == 0 {
		return &DefaultResponderBadRequest{}
	}
	user := &User{}
	if err := h.store.GetUser(strings.ToLower(req.ID), user); err != nil {
		if !h.store.IsErrNotFound(err) {
			log.Println(err.Error())
			return &DefaultResponderError{}
		}
		return &DefaultResponderNoContent{}
	}
	return &DefaultResponderOK{}
}

func (c *Controller) ForgetPassword(req *user.ForgetPasswordParams) middleware.Responder {

}

func (c *Controller) LoginUser(req *user.LoginUserParams) middleware.Responder {

}

func (c *Controller) NewPassword(req *user.NewPasswordParams) middleware.Responder {

}

func (c *Controller) VerifyUser(req *user.VerifyUserParams) middleware.Responder {

}
