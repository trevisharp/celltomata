package services

type EmailService interface {
	Send(origin, password, destination, message string) error
}
