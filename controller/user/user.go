package user

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/jacygao/auth/restapi/operations/user"
	"github.com/jacygao/auth/store"
)

type Controller struct {
	Store store.DB
}

func NewController(s store.DB) *Controller {
	return &Controller{
		Store: s,
	}
}

func (c *Controller) Signup(req *user.SignupUserParams) middleware.Responder {

}

func (c *Controller) FindUser(req *user.FindUserParams) middleware.Responder {

}

func (c *Controller) ForgetPassword(req *user.ForgetPasswordParams) middleware.Responder {

}

func (c *Controller) LoginUser(req *user.LoginUserParams) middleware.Responder {

}

func (c *Controller) NewPassword(req *user.NewPasswordParams) middleware.Responder {

}

func (c *Controller) VerifyUser(req *user.VerifyUserParams) middleware.Responder {

}
