package repositories

import "github.com/benzdeus/assessment/entities"

func (repo expenseRepo) UpdateExpenese(id string, expenseRequest entities.ExpenseModel) (entities.ExpenseModel, error) {
	expenseRequest.ID = id
	error := repo.db.Where("id=?", id).Updates(&expenseRequest).Error

	return expenseRequest, error
}
