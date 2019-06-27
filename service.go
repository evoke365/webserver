// Package auth contains a service that handles authentication.
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

// Config defines configuration variables.
type Config struct {
	HTTPPort              int
}

// Service defines Auth Service instance structure and dependecies.
type Service struct {
	conf    Config
	mailer  mail.Mailer
	handler *Handler
}

// Start starts the Auth Service.
func (s *Service) Start() {
	router := httprouter.New()
	router.GET("/health", s.handler.Health)
	router.GET("/redirect/:code", s.handler.Redirect)
	router.GET("/user/find/:id", s.handler.User)
	router.POST("/user/login", s.handler.Login)
	router.POST("/user/signup", s.handler.Signup)

	log.Printf("HTTP Server listenning on port %d", s.conf.HTTPPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", s.conf.HTTPPort), router))
}

// Stop contains processes for a graceful shutdown.
func (s *Service) Stop() {
	// Close everything
	log.Println("Shutting down...")
}

// WithMongoDB registers Mongo DB in the Auth Service instance.
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

// WithMemoryDB registers In-memory db in the Auth Service instance.
func (s *Service) WithMemoryDB() (*Service, error) {
	model := NewMemoryDB()
	if s.mailer == nil {
		return nil, errors.New("you must call withMailer first")
	}
	s.handler = NewHandler(model, s.mailer)
	return s, nil
}

// WithMailer registers a mail service in the Auth Service instance.
func (s *Service) WithMailer(ms mail.Mailer) *Service {
	s.mailer = ms
	return s
}

// NewService initialises a new Auth Service instance.
func NewService(c Config) *Service {
	return &Service{
		conf: c,
	}
}
