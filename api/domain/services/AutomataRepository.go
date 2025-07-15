package services

import "github.com/trevisharp/celltomata/api/domain/models"

type AutomataRepository interface {
	Get(id int) *models.Automata
	Save(automata *models.Automata)
	Delete(id int)
}
