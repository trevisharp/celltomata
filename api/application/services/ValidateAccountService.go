package services

type ValidateAccountService struct {
	EmailService EmailService
}

func (s ValidateAccountService) SendEmail(username, password string) error {
	return nil
}
