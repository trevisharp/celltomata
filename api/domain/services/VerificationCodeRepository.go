package services

import "github.com/trevisharp/celltomata/api/domain/models"

type VerificationCodeRepository interface {
	CreateNew(userId int) (string, error)
	Get(code string) (*models.VerificationCode, error)
	Delete(code string) error
}
