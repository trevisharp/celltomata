package infrastructure

import (
	"github.com/trevisharp/celltomata/api/domain/models"
)

type SupabaseUserRepository struct{}

func (s SupabaseUserRepository) Find(username string) (*models.User, error) {
	return nil, nil

}

func (s SupabaseUserRepository) Create(user *models.User) error {
	return nil

}

func (s SupabaseUserRepository) Update(user *models.User) error {
	return nil
}
