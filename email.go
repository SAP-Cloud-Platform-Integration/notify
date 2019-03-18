package main

import (
	"log"
	"strconv"

	gomail "gopkg.in/gomail.v2"
)

const EMAIL_TITLE = "CPI Error Notification"

// EmailPayload entity
type EmailPayload struct {
	To      []string
	Content string
}

type EmailSender struct {
	config SMTP
	mailer *gomail.Dialer
}

func (s *EmailSender) SendEmail(payload EmailPayload) {
	m := gomail.NewMessage()
	m.SetHeader("From", s.config.Username)
	m.SetHeader("To", payload.To...)
	m.SetHeader("Subject", EMAIL_TITLE)
	m.SetBody("text/html", payload.Content)
	if err := s.mailer.DialAndSend(m); err != nil {
		log.Println(err)
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
