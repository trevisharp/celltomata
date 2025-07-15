package services

import "github.com/trevisharp/celltomata/api/domain/models"

type UserRepository interface {
	Find(username string) *models.User
	Create(user *models.User)
	Update(user *models.User)
}
