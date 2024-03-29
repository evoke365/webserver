// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/evoke365/webserver/restapi/operations/health"
	"github.com/evoke365/webserver/restapi/operations/note"
	"github.com/evoke365/webserver/restapi/operations/profile"
	"github.com/evoke365/webserver/restapi/operations/user"
)

// NewEvoke365NetOpenAPISpecAPI creates a new Evoke365NetOpenAPISpec instance
func NewEvoke365NetOpenAPISpecAPI(spec *loads.Document) *Evoke365NetOpenAPISpecAPI {
	return &Evoke365NetOpenAPISpecAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),

		NoteAddNoteHandler: note.AddNoteHandlerFunc(func(params note.AddNoteParams) middleware.Responder {
			return middleware.NotImplemented("operation note.AddNote has not yet been implemented")
		}),
		UserSignupUserHandler: user.SignupUserHandlerFunc(func(params user.SignupUserParams) middleware.Responder {
			return middleware.NotImplemented("operation user.SignupUser has not yet been implemented")
		}),
		ProfileAutheticateProfileHandler: profile.AutheticateProfileHandlerFunc(func(params profile.AutheticateProfileParams) middleware.Responder {
			return middleware.NotImplemented("operation profile.AutheticateProfile has not yet been implemented")
		}),
		NoteDeleteNoteHandler: note.DeleteNoteHandlerFunc(func(params note.DeleteNoteParams) middleware.Responder {
			return middleware.NotImplemented("operation note.DeleteNote has not yet been implemented")
		}),
		UserFindUserHandler: user.FindUserHandlerFunc(func(params user.FindUserParams) middleware.Responder {
			return middleware.NotImplemented("operation user.FindUser has not yet been implemented")
		}),
		UserForgetPasswordHandler: user.ForgetPasswordHandlerFunc(func(params user.ForgetPasswordParams) middleware.Responder {
			return middleware.NotImplemented("operation user.ForgetPassword has not yet been implemented")
		}),
		NoteGetNotesHandler: note.GetNotesHandlerFunc(func(params note.GetNotesParams) middleware.Responder {
			return middleware.NotImplemented("operation note.GetNotes has not yet been implemented")
		}),
		ProfileGetProfileHandler: profile.GetProfileHandlerFunc(func(params profile.GetProfileParams) middleware.Responder {
			return middleware.NotImplemented("operation profile.GetProfile has not yet been implemented")
		}),
		HealthHealthzHandler: health.HealthzHandlerFunc(func(params health.HealthzParams) middleware.Responder {
			return middleware.NotImplemented("operation health.Healthz has not yet been implemented")
		}),
		UserLoginUserHandler: user.LoginUserHandlerFunc(func(params user.LoginUserParams) middleware.Responder {
			return middleware.NotImplemented("operation user.LoginUser has not yet been implemented")
		}),
		UserNewPasswordHandler: user.NewPasswordHandlerFunc(func(params user.NewPasswordParams) middleware.Responder {
			return middleware.NotImplemented("operation user.NewPassword has not yet been implemented")
		}),
		NoteUpdateNoteHandler: note.UpdateNoteHandlerFunc(func(params note.UpdateNoteParams) middleware.Responder {
			return middleware.NotImplemented("operation note.UpdateNote has not yet been implemented")
		}),
		UserVerifyUserHandler: user.VerifyUserHandlerFunc(func(params user.VerifyUserParams) middleware.Responder {
			return middleware.NotImplemented("operation user.VerifyUser has not yet been implemented")
		}),
	}
}

/*Evoke365NetOpenAPISpecAPI Reusable authentication server from evoke365.net */
type Evoke365NetOpenAPISpecAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// NoteAddNoteHandler sets the operation handler for the add note operation
	NoteAddNoteHandler note.AddNoteHandler
	// UserSignupUserHandler sets the operation handler for the signup user operation
	UserSignupUserHandler user.SignupUserHandler
	// ProfileAutheticateProfileHandler sets the operation handler for the autheticate profile operation
	ProfileAutheticateProfileHandler profile.AutheticateProfileHandler
	// NoteDeleteNoteHandler sets the operation handler for the delete note operation
	NoteDeleteNoteHandler note.DeleteNoteHandler
	// UserFindUserHandler sets the operation handler for the find user operation
	UserFindUserHandler user.FindUserHandler
	// UserForgetPasswordHandler sets the operation handler for the forget password operation
	UserForgetPasswordHandler user.ForgetPasswordHandler
	// NoteGetNotesHandler sets the operation handler for the get notes operation
	NoteGetNotesHandler note.GetNotesHandler
	// ProfileGetProfileHandler sets the operation handler for the get profile operation
	ProfileGetProfileHandler profile.GetProfileHandler
	// HealthHealthzHandler sets the operation handler for the healthz operation
	HealthHealthzHandler health.HealthzHandler
	// UserLoginUserHandler sets the operation handler for the login user operation
	UserLoginUserHandler user.LoginUserHandler
	// UserNewPasswordHandler sets the operation handler for the new password operation
	UserNewPasswordHandler user.NewPasswordHandler
	// NoteUpdateNoteHandler sets the operation handler for the update note operation
	NoteUpdateNoteHandler note.UpdateNoteHandler
	// UserVerifyUserHandler sets the operation handler for the verify user operation
	UserVerifyUserHandler user.VerifyUserHandler
	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *Evoke365NetOpenAPISpecAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *Evoke365NetOpenAPISpecAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *Evoke365NetOpenAPISpecAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *Evoke365NetOpenAPISpecAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *Evoke365NetOpenAPISpecAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *Evoke365NetOpenAPISpecAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *Evoke365NetOpenAPISpecAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the Evoke365NetOpenAPISpecAPI
func (o *Evoke365NetOpenAPISpecAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.NoteAddNoteHandler == nil {
		unregistered = append(unregistered, "note.AddNoteHandler")
	}
	if o.UserSignupUserHandler == nil {
		unregistered = append(unregistered, "user.SignupUserHandler")
	}
	if o.ProfileAutheticateProfileHandler == nil {
		unregistered = append(unregistered, "profile.AutheticateProfileHandler")
	}
	if o.NoteDeleteNoteHandler == nil {
		unregistered = append(unregistered, "note.DeleteNoteHandler")
	}
	if o.UserFindUserHandler == nil {
		unregistered = append(unregistered, "user.FindUserHandler")
	}
	if o.UserForgetPasswordHandler == nil {
		unregistered = append(unregistered, "user.ForgetPasswordHandler")
	}
	if o.NoteGetNotesHandler == nil {
		unregistered = append(unregistered, "note.GetNotesHandler")
	}
	if o.ProfileGetProfileHandler == nil {
		unregistered = append(unregistered, "profile.GetProfileHandler")
	}
	if o.HealthHealthzHandler == nil {
		unregistered = append(unregistered, "health.HealthzHandler")
	}
	if o.UserLoginUserHandler == nil {
		unregistered = append(unregistered, "user.LoginUserHandler")
	}
	if o.UserNewPasswordHandler == nil {
		unregistered = append(unregistered, "user.NewPasswordHandler")
	}
	if o.NoteUpdateNoteHandler == nil {
		unregistered = append(unregistered, "note.UpdateNoteHandler")
	}
	if o.UserVerifyUserHandler == nil {
		unregistered = append(unregistered, "user.VerifyUserHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *Evoke365NetOpenAPISpecAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *Evoke365NetOpenAPISpecAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	return nil
}

// Authorizer returns the registered authorizer
func (o *Evoke365NetOpenAPISpecAPI) Authorizer() runtime.Authorizer {
	return nil
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *Evoke365NetOpenAPISpecAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *Evoke365NetOpenAPISpecAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *Evoke365NetOpenAPISpecAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the evoke365 net open API spec API
func (o *Evoke365NetOpenAPISpecAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *Evoke365NetOpenAPISpecAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/note"] = note.NewAddNote(o.context, o.NoteAddNoteHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/user/signup"] = user.NewSignupUser(o.context, o.UserSignupUserHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/profile/authenticate"] = profile.NewAutheticateProfile(o.context, o.ProfileAutheticateProfileHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/note/{id}"] = note.NewDeleteNote(o.context, o.NoteDeleteNoteHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/user/find/{id}"] = user.NewFindUser(o.context, o.UserFindUserHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/user/forget"] = user.NewForgetPassword(o.context, o.UserForgetPasswordHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/notes/{userId}"] = note.NewGetNotes(o.context, o.NoteGetNotesHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/profile/{id}"] = profile.NewGetProfile(o.context, o.ProfileGetProfileHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/healthz"] = health.NewHealthz(o.context, o.HealthHealthzHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/user/login"] = user.NewLoginUser(o.context, o.UserLoginUserHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/user/password"] = user.NewNewPassword(o.context, o.UserNewPasswordHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/note/{id}"] = note.NewUpdateNote(o.context, o.NoteUpdateNoteHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/user/verify"] = user.NewVerifyUser(o.context, o.UserVerifyUserHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *Evoke365NetOpenAPISpecAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *Evoke365NetOpenAPISpecAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *Evoke365NetOpenAPISpecAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *Evoke365NetOpenAPISpecAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *Evoke365NetOpenAPISpecAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
