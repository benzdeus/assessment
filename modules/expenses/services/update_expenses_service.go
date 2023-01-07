package services

import "github.com/benzdeus/assessment/entities"

func (service expeneseService) UpdateExpenese(id string, expeneseRequest entities.ExpenseModel) (entities.ExpenseModel, error) {
	return service.repo.NewExpenese(expeneseRequest)
}
