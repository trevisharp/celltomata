package services

type CryptoService interface {
	EncryptPassword(rawPassword string) (string, error)
	CheckPassword(rawPassword, hashPassowrd string) bool
}
