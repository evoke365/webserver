// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/jacygao/auth/controller"
	"github.com/jacygao/auth/restapi/operations"
	"github.com/jacygao/auth/restapi/operations/health"
	"github.com/jacygao/auth/restapi/operations/user"
)

//go:generate swagger generate server --target ../../auth --name Evoke365NetOpenAPISpec --spec ../openapi/spec.yaml

func configureFlags(api *operations.Evoke365NetOpenAPISpecAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.Evoke365NetOpenAPISpecAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	controller := controller.New()

	if api.UserSignupUserHandler == nil {
		api.UserSignupUserHandler = user.SignupUserHandlerFunc(func(params user.SignupUserParams) middleware.Responder {
			return middleware.NotImplemented("operation user.SignupUser has not yet been implemented")
		})
	}
	if api.UserFindUserHandler == nil {
		api.UserFindUserHandler = user.FindUserHandlerFunc(func(params user.FindUserParams) middleware.Responder {
			return middleware.NotImplemented("operation user.FindUser has not yet been implemented")
		})
	}
	if api.UserForgetPasswordHandler == nil {
		api.UserForgetPasswordHandler = user.ForgetPasswordHandlerFunc(func(params user.ForgetPasswordParams) middleware.Responder {
			return middleware.NotImplemented("operation user.ForgetPassword has not yet been implemented")
		})
	}
	if api.HealthHealthzHandler == nil {
		api.HealthHealthzHandler = health.HealthzHandlerFunc(func(params health.HealthzParams) middleware.Responder {
			return controller.Health.Healthz()
		})
	}
	if api.UserLoginUserHandler == nil {
		api.UserLoginUserHandler = user.LoginUserHandlerFunc(func(params user.LoginUserParams) middleware.Responder {
			return middleware.NotImplemented("operation user.LoginUser has not yet been implemented")
		})
	}
	if api.UserNewPasswordHandler == nil {
		api.UserNewPasswordHandler = user.NewPasswordHandlerFunc(func(params user.NewPasswordParams) middleware.Responder {
			return middleware.NotImplemented("operation user.NewPassword has not yet been implemented")
		})
	}
	if api.UserVerifyUserHandler == nil {
		api.UserVerifyUserHandler = user.VerifyUserHandlerFunc(func(params user.VerifyUserParams) middleware.Responder {
			return middleware.NotImplemented("operation user.VerifyUser has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
