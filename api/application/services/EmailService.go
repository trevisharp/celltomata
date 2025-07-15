package services

type EmailService interface {
	Send(origin, password, destination, title, message string) error
}
