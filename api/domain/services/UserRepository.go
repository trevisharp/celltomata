package services

import (
	"github.com/trevisharp/celltomata/api/domain/models"
)

type UserRepository interface {
	Find(login string) (*models.User, error)
	Get(id int) (*models.User, error)
	Create(user *UserData) (int, error)
	Update(user *models.User) error
}
