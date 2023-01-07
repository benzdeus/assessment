package repositories

import "gorm.io/gorm"

type expenseRepo struct {
	db *gorm.DB
}

func NewExpeneseRepo(db *gorm.DB) *expenseRepo {
	return &expenseRepo{
		db,
	}
}
