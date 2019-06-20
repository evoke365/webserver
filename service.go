// package auth starts a service that handles authentication
// wrapping http api and provides /auth endpoint
package auth

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"gopkg.in/mgo.v2"

	"github.com/jacygao/mail"
	"github.com/julienschmidt/httprouter"
)

type Config struct {
	HttpPort              int
	AdminEmail            string
	RedirectUri           string
	PasswordEmailFilePath string
	PasswordEmailFileName string
}

type Service struct {
	conf    Config
	mailer  *mail.Service
	handler *Handler
}

func (s *Service) Start() {
	router := httprouter.New()
	router.GET("/health", s.handler.Health)
	router.GET("/redirect/:code", s.handler.Redirect)
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

func (s *Service) WithMongoDB(session *mgo.Session, dbName, collection string) (*Service, error) {
	model, err := NewMongoDB(session, dbName, collection)
	if err != nil {
		return nil, err
	}
	if s.mailer == nil {
		return nil, errors.New("you must call withMailer first")
	}
	s.handler = NewHandler(model, s.mailer)
	return s, nil
}

func (s *Service) WithMailer(ms *mail.Service) *Service {
	s.mailer = ms
	return s
}

func NewService(c Config) *Service {
	return &Service{
		conf: c,
	}
}
