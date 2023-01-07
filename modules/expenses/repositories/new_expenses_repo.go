package repositories

import "github.com/benzdeus/assessment/entities"

func (repo expenseRepo) NewExpenese(expenseRequest entities.ExpenseModel) (entities.ExpenseModel, error) {
	error := repo.db.Create(&expenseRequest).Error

	return expenseRequest, error
}
