package profile

import (
	"encoding/json"
	"log"

	"github.com/go-openapi/runtime/middleware"
	"github.com/jacygao/auth/controller/internal/responder"
	"github.com/jacygao/auth/restapi/operations/profile"
	"github.com/jacygao/auth/store"
	"github.com/jacygao/auth/store/data"
)

// Controller defines HTTP handlers.
type Controller struct {
	store store.DB
}

// NewController returns a new instance of Controller.
func NewController(db store.DB) *Controller {
	return &Controller{
		store: db,
	}
}

// Authenticate implements the HTTP handler logic for /profile/authenticate
func (c *Controller) Authenticate(req *profile.AutheticateProfileParams) middleware.Responder {
	if len(req.Body.Token) == 0 {
		return responder.DefaultBadRequest()
	}

	user := &data.User{}
	if err := c.store.FindUserByTok(req.Body.Token, user); err != nil {
		log.Println(err.Error())
		return responder.DefaultServerError()
	}

	// TODO: TouchTok should return new token expiry.
	if err := c.store.TouchTok(user.Email); err != nil {
		log.Println(err.Error())
		return responder.DefaultServerError()
	}

	b, err := json.Marshal(user)
	if err != nil {
		log.Println(err.Error())
		return responder.DefaultServerError()
	}
	return responder.DefaultOK().WithResponse(b)
}

// Get implements the HTTP handler logic for /profile/{id}
func (c *Controller) Get(req *profile.GetProfileParams) middleware.Responder {
	if len(req.ID) == 0 {
		return responder.DefaultBadRequest()
	}

	user := &data.User{}
	if err := c.store.GetUser(req.ID, user); err != nil {
		if !c.store.IsErrNotFound(err) {
			log.Println(err.Error())
			return responder.DefaultServerError()
		}
		responder.DefaultNoContent()
	}

	profile := Profile{
		Email:    user.Email,
		Timezone: user.Timezone,
		IsActive: user.IsActive,
		Created:  user.Created,
		Modified: user.Modified,
	}

	b, err := json.Marshal(profile)
	if err != nil {
		log.Println(err.Error())
		return responder.DefaultServerError()
	}

	return responder.DefaultOK().WithResponse(b)
}
