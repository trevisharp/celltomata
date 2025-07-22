package infrastructure

import (
	"fmt"

	"github.com/trevisharp/celltomata/api/domain/models"
	"github.com/trevisharp/celltomata/api/domain/services"
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

func (s SupabaseUserRepository) Get(id int) (*models.User, error) {
	query := fmt.Sprintf("ID.eq=", id)
	users, err := SupabaseGet[models.User]("User", query)
	if err != nil {
		return nil, err
	}

	if len(*users) == 0 {
		return nil, nil
	}

	return &(*users)[0], nil
}

func (s SupabaseUserRepository) Create(user *services.UserData) (int, error) {
	objs, err := SupabasePost[services.UserData, models.User]("User", *user)
	if err != nil {
		return 0, err
	}

	if len(*objs) == 0 {
		return 0, nil
	}

	return (*objs)[0].ID, nil
}

func (s SupabaseUserRepository) Update(user *models.User) error {
	id := user.ID
	return SupabasePatch("User", id, user)
}
