package main

import (
	"log"
	"strconv"

	gomail "gopkg.in/gomail.v2"
)

// EMAIL_TITLE constant
const EMAIL_TITLE = "CPI Error Notification"

// EmailPayload entity
type EmailPayload struct {
	To      []string
	Content string
}

// EmailSender type
type EmailSender struct {
	config SMTP
	mailer *gomail.Dialer
}

// SendEmail func
func (s *EmailSender) SendEmail(payload EmailPayload) {
	m := gomail.NewMessage()
	m.SetHeader("From", s.config.From)
	m.SetHeader("To", payload.To...)
	m.SetHeader("Subject", EMAIL_TITLE)
	// set high priority
	m.SetHeader("X-Priority", "1 (Highest)")
	m.SetHeader("X-MSMail-Priority", "High")
	m.SetHeader("Importance", "High")
	m.SetBody("text/html", payload.Content)
	// set mail
	if err := s.mailer.DialAndSend(m); err != nil {
		log.Println(err)
	}
}

// NewSender constructor
func NewSender(config SMTP) *EmailSender {
	port, _ := strconv.ParseInt(config.Port, 10, 32)
	rt := &EmailSender{
		config,
		gomail.NewDialer(config.Server, int(port), config.Username, config.Password),
	}
	return rt
}
