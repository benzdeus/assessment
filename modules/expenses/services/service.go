package services

import "github.com/benzdeus/assessment/entities"

type expeneseService struct {
	repo entities.ExpeneseRepository
}

func NewExpeneseService(repo entities.ExpeneseRepository) *expeneseService {
	return &expeneseService{
		repo,
	}
}
