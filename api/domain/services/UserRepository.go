package services

import "github.com/trevisharp/celltomata/domain/models"

type UserRepository interface {
	Find(username string) *models.User
	Create(user *models.User)
	Update(user *models.User)
}
