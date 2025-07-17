package infrastructure

import (
	"fmt"

	"github.com/trevisharp/celltomata/api/domain/models"
)

type SupabaseUserRepository struct{}

func (s SupabaseUserRepository) Find(login string) (*models.User, error) {
	query := fmt.Sprintf("or=(Username.eq.%s,Email.eq.%s)", login, login)
	users, err := SupabaseGet[models.User]("User", query)
	if err != nil {
		return nil, err
	}

	if len(*users) == 0 {
		return nil, nil
	}

	return &(*users)[0], nil
}

func (s SupabaseUserRepository) Create(user *models.User) error {
	return nil

}

func (s SupabaseUserRepository) Update(user *models.User) error {
	return nil
}
