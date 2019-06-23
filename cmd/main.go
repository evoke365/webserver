package main

import (
	"github.com/jacygao/mail"
	"github.com/studybox/auth"
)

type fakeMailer struct {
}

func newFakeMailer() *fakeMailer {
	return &fakeMailer{}
}

func (fm *fakeMailer) Send(to string, msg mail.Message) error {
	return nil
}

func main() {
	conf := auth.Config{
		HTTPPort:    8090,
		RedirectURI: "https://google.com",
	}
	service, err := auth.NewService(conf).WithMailer(newFakeMailer()).WithMemoryDB()
	if err != nil {
		panic(err)
	}

	service.Start()
}
