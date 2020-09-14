package controller

import (
	"github.com/evoke365/webserver/controller/health"
	"github.com/evoke365/webserver/controller/note"
	"github.com/evoke365/webserver/controller/profile"
	"github.com/evoke365/webserver/controller/user"
	"github.com/evoke365/webserver/event"
	"github.com/evoke365/webserver/event/bus"
	"github.com/evoke365/webserver/pkg/mailer"
	"github.com/evoke365/webserver/store"
)

type Controller struct {
	Health  *health.Controller
	User    *user.Controller
	Profile *profile.Controller
	Note    *note.Controller
	Event   *event.Controller
}

func New(db store.DB, m *mailer.Client) *Controller {
	return &Controller{
		Health:  health.NewController(),
		User:    user.NewController(db, m, user.DefaultConfig()),
		Profile: profile.NewController(db),
		Note:    note.NewController(db),
		Event: event.NewController(db, bus.NewChannelBus(bus.Config{
			// TODO: config driven
			Topics:                  []string{"user", "reminder"},
			QueueRetryDelayMilliSec: 100,
		})),
	}
}
