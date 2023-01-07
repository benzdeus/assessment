package repositories

import "github.com/benzdeus/assessment/entities"

func (repo expenseRepo) GetExpenese(id string) (entities.ExpenseResponse, error) {
	res := entities.ExpenseResponse{}
	error := repo.db.Find(&res, "id=?", id).Error

	return res, error
}
