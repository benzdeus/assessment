package services

import "github.com/benzdeus/assessment/entities"

func (service expeneseService) GetExpenese(id string) (entities.ExpenseResponse, error) {
	return service.repo.GetExpenese(id)
}
