// package auth starts a service that handles authentication
// wrapping http api and provides /auth endpoint
package auth

import (
	"fmt"
	"log"
	"net/http"

	"gopkg.in/mgo.v2"

	"github.com/jacygao/mail"
	"github.com/julienschmidt/httprouter"
)

type Config struct {
	AdminEmail  string
	HttpPort    int
	RedirectUri string
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
	router.POST("/user/signup", s.handler.Signup)

	log.Printf("HTTP Server listenning on port %d", s.conf.HttpPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", s.conf.HttpPort), router))
}

func (s *Service) Stop() {
	// Close everything
	log.Println("Shutting down...")
}

func (s *Service) WithMongoDB(session *mgo.Session, dbName, collection string) error {
	model, err := NewMongoDB(session, dbName, collection)
	if err != nil {
		return err
	}
	mailer := mail.NewService(mail.Config{})
	s.handler = NewHandler(model, mailer)
	return nil
}

func NewService(c Config) *Service {
	return &Service{
		conf: c,
	}
}
