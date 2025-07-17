package services

type ValidateAccountService struct {
	EmailService EmailService
}

func (s ValidateAccountService) SendEmail(username, email, password string) error {
	return nil
}
