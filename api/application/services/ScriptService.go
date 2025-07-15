package services

import "github.com/trevisharp/celltomata/api/domain/models"

type ScriptService struct{}

func (s *ScriptService) Generate(script string) *models.WorldRule {
	var commands []any

	commands = append(commands, "oi")

	return &models.WorldRule{
		Commands: commands,
	}
}
