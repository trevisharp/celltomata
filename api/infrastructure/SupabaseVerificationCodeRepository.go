package infrastructure

import (
	"github.com/trevisharp/celltomata/api/application/services"
	"github.com/trevisharp/celltomata/api/domain/models"
)

type SupabaseVerificationCodeRepository struct {
	generator services.CodeGenerator
}

func (s SupabaseVerificationCodeRepository) CreateNew(userId int) (string, error) {
	code, err := s.generator.NewCode(32)
	if err != nil {
		return "", err
	}
	data := models.VerificationCode{
		UserID: userId,
		Code:   code,
	}
	_, err = SupabasePost[models.VerificationCode, models.VerificationCode]("VerificationCode", data)
	if err != nil {
		return "", err
	}
	return code, nil
}

func (s SupabaseVerificationCodeRepository) Get(code string) (*models.VerificationCode, error) {
	verCodes, err := SupabaseGet[models.VerificationCode]("VerificationCode", "Code=eq."+code)
	if err != nil {
		return nil, err
	}

	if len(*verCodes) == 0 {
		return nil, nil
	}

	verificationCode := (*verCodes)[0]
	return &verificationCode, nil
}

func (s SupabaseVerificationCodeRepository) Delete(code string) error {
	err := SupabaseDelete("VerificationCode", "Code=eq."+code)
	return err
}
