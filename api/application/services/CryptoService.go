package services

type CryptoService interface {
	EncryptPassword(rawPassword string) string
}
