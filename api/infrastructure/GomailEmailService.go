package infrastructure

import "gopkg.in/gomail.v2"

type GomailEmailService struct{}

func (s GomailEmailService) Send(origin, password, destination, title, message string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", origin)
	m.SetHeader("To", destination)
	m.SetHeader("Subject", title)
	m.SetBody("text/plain", message)

	d := gomail.NewDialer("smtp.gmail.com", 587, origin, password)
	return d.DialAndSend()
}
