package infrastructure

import "golang.org/x/crypto/bcrypt"

type BCryptService struct{}

func (s BCryptService) EncryptPassword(rawPassword string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
	return string(bytes), err
}

func (s BCryptService) CheckPassword(rawPassword, hashPassowrd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassowrd), []byte(rawPassword))
	return err != nil
}
