package controller

import (
	"github.com/jacygao/auth/controller/health"
	"github.com/jacygao/auth/controller/profile"
	"github.com/jacygao/auth/controller/user"
)

type Controller struct {
	Health  *health.Controller
	User    *user.Controller
	Profile *profile.Controller
}

func New() *Controller {
	return &Controller{
		Health:  health.NewController(),
		User:    user.NewController(),
		Profile: profile.NewController(),
	}
}
