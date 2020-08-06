package controller

import (
	"github.com/jacygao/auth/controller/health"
	"github.com/jacygao/auth/controller/profile"
	"github.com/jacygao/auth/controller/user"
	"github.com/jacygao/auth/pkg/mailer"
	"github.com/jacygao/auth/store"
)

type Controller struct {
	Health  *health.Controller
	User    *user.Controller
	Profile *profile.Controller
}

func New(db store.DB, m *mailer.Client) *Controller {
	return &Controller{
		Health:  health.NewController(),
		User:    user.NewController(db, m, user.DefaultConfig()),
		Profile: profile.NewController(db),
	}
}
