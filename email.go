package main

// EmailPayload entity
type EmailPayload struct {
	From    string
	To      []string
	Title   string
	Content string
}

func sendEmail(config SMTP, payload EmailPayload) {

}
