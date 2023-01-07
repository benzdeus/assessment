package repositories

import "github.com/benzdeus/assessment/entities"

func (repo expenseRepo) GetExpeneseList() ([]entities.ExpenseResponse, error) {
	res := []entities.ExpenseResponse{}
	error := repo.db.Find(&res).Error

	return res, error
}
