package health

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) Healthz() middleware.Responder {
	return &HealthzOK{}
}

type HealthzOK struct {
}

func (h *HealthzOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("Auth service is up and running"))
}
