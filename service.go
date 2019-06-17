// package auth starts a service that handles authentication
// wrapping http api and provides /auth endpoint
package auth

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Config struct {
	AdminEmail string
	Port       int
}

type Service struct {
	conf    Config
	handler *Handler
}

func (s *Service) Start() {
	router := httprouter.New()
	router.GET("/", s.handler.Index)
	router.GET("/health", s.handler.Health)
	router.POST("/user/register", s.handler.Register)
	router.POST("/user/login", s.handler.Login)
	router.POST("/user", s.handler.Signup)

	log.Printf("HTTP Server listenning on port %d", s.conf.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", s.conf.Port), router))
}

func (s *Service) Stop() {
	// Close everything
	log.Println("Shutting down...")
}

func NewService(c Config, m Model) *Service {
	return &Service{
		conf:    c,
		handler: NewHandler(m),
	}
}
