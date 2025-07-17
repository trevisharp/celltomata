package services

import "github.com/trevisharp/celltomata/api/domain/models"

type UserRepository interface {
	Find(login string) (*models.User, error)
	Create(user *models.User) error
	Update(user *models.User) error
}
