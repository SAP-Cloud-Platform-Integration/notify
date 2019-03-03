package main

import (
	"strconv"

	"github.com/sirupsen/logrus"

	"gopkg.in/gomail.v2"
)

// EmailPayload entity
type EmailPayload struct {
	From    string
	To      []string
	Title   string
	Content string
}

type EmailSender struct {
	config SMTP
	mailer *gomail.Dialer
}

func (s *EmailSender) SendEmail(payload EmailPayload) {
	m := gomail.NewMessage()
	m.SetHeader("From", payload.From)
	m.SetHeader("To", payload.To...)
	m.SetHeader("Subject", payload.Title)
	m.SetBody("text/html", payload.Content)
	if err := s.mailer.DialAndSend(m); err != nil {
		logrus.Error(err)
	}
}

func NewSender(config SMTP) *EmailSender {
	port, _ := strconv.ParseInt(config.Port, 10, 32)
	rt := &EmailSender{
		config,
		gomail.NewDialer(config.Server, int(port), config.Username, config.Password),
	}
	return rt

}
