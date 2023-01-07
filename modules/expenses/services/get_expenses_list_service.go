package services

import "github.com/benzdeus/assessment/entities"

func (service expeneseService) GetExpeneseList() ([]entities.ExpenseResponse, error) {
	return service.repo.GetExpeneseList()
}
