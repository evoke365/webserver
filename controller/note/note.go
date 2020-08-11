package note

import (
	"encoding/json"
	"log"
	"time"

	"github.com/evoke365/webserver/controller/internal/responder"
	"github.com/evoke365/webserver/store"
	"github.com/evoke365/webserver/store/data"
	"github.com/go-openapi/runtime/middleware"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/evoke365/webserver/restapi/operations/note"
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

func (c *Controller) GetNotes(req *note.GetNotesParams) middleware.Responder {
	ctx := req.HTTPRequest.Context()

	// get userId by Token
	u := &data.User{}
	if err := c.store.GetUser(req.Token.String(), u); err != nil {
		log.Println(err.Error())
		return responder.DefaultUnauthorised()
	}

	notes, err := c.store.GetNotes(ctx, u.Email)
	if err != nil {
		log.Println(err.Error())
		return responder.DefaultServerError()
	}

	b, err := json.Marshal(notes)
	if err != nil {
		log.Println(err.Error())
		return responder.DefaultServerError()
	}

	return responder.DefaultOK().WithResponse(b)
}

func (c *Controller) AddNote(req *note.AddNoteParams) middleware.Responder {
	ctx := req.HTTPRequest.Context()
	u := &data.User{}

	if err := c.store.GetUser(req.Token.String(), u); err != nil {
		log.Println(err.Error())
		return responder.DefaultUnauthorised()

	}

	postNote := &data.Note{
		ID:        primitive.NewObjectID(),
		UserID:    u.Email,
		Keyword:   req.Body.Title,
		Answer:    req.Body.Body,
		Important: true,
		Created:   time.Now(),
		Modified:  time.Now(),
		Deleted:   false,
	}
	if err := c.store.InsertNote(ctx, postNote); err != nil {
		log.Println(err.Error())
		return &responder.ServerError{}
	}

	b, err := json.Marshal(postNote)
	if err != nil {
		log.Println(err.Error())
		return &responder.ServerError{}
	}

	return responder.DefaultOK().WithResponse(b)
}

func (c *Controller) DeleteNote(req *note.DeleteNoteParams) middleware.Responder {
	ctx := req.HTTPRequest.Context()
	u := &data.User{}

	if err := c.store.GetUser(req.Token.String(), u); err != nil {
		log.Println(err.Error())
		return responder.DefaultUnauthorised()
	}

	if err := c.store.DeleteNote(ctx, u.Email, req.ID.String()); err != nil {
		log.Println(err.Error())
		return responder.DefaultServerError()
	}

	return responder.DefaultOK()
}

func (c *Controller) UpdateNote(req *note.UpdateNoteParams) middleware.Responder {
	ctx := req.HTTPRequest.Context()
	u := &data.User{}

	if err := c.store.GetUser(req.Token.String(), u); err != nil {
		log.Println(err.Error())
		return responder.DefaultServerError()
	}

	dataMap := make(map[string]interface{})
	dataMap["important"] = req.Body.IsImportant
	dataMap["modified"] = time.Now()
	if _, err := c.store.UpdateNote(ctx, u.Email, req.ID.String(), dataMap); err != nil {
		log.Println(err)
		return responder.DefaultBadRequest()
	}

	return responder.DefaultOK()
}
