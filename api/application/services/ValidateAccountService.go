package services

type ValidateAccountService struct {
	EmailService EmailService
}

func (s ValidateAccountService) SendEmail() error {
	return nil
}
