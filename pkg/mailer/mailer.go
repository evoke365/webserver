package mailer

import (
	"bytes"
	"fmt"
	"html/template"
	"strconv"

	"gopkg.in/gomail.v2"
)

// Config contains required values that defines a mailer client.
type Config struct {
	Hostname string
	Port     string
	Username string
	Password string
}

// Client contains dependencies to initialise a mailer client.
type Client struct {
	conf Config
}

// NewClient returns a new instance of the mailer client.
func NewClient(c Config) *Client {
	return &Client{
		conf: c,
	}
}

// Mail contains required data for sending out an email.
type Mail struct {
	TemplatePath string
	Filename     string
	Sender       string
	Recipient    string
	Data         interface{}
}

// SendVerificationEmail sends a verification email to a recipient.
// TODO: write a geneic method for sneding emails.
func (c *Client) SendVerificationEmail(m Mail) error {
	tplPath := fmt.Sprintf("%s/%s", m.TemplatePath, m.Filename)
	t := template.New(m.Filename)
	t, err := t.ParseFiles(tplPath)
	if err != nil {
		return err
	}

	templateData := struct {
		CODE string
	}{
		CODE: m.Data.(string),
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, templateData); err != nil {
		return err
	}

	msg := gomail.NewMessage()
	msg.SetHeader("From", m.Sender)
	msg.SetHeader("To", m.Recipient)
	msg.SetHeader("Subject", "Verify you account")
	msg.SetBody("text/html", tpl.String())

	port, _ := strconv.Atoi(c.conf.Port)
	d := gomail.NewDialer(c.conf.Hostname, port, c.conf.Username, c.conf.Password)

	if err := d.DialAndSend(msg); err != nil {
		return err
	}

	return nil
}
