package entities

type ExpenseModel struct {
	ID     string    `json:"id" gorm:"AUTO_INCREMENT;type:int"`
	Title  string    `json:"title"`
	Amount int32     `json:"amount"`
	Note   string    `json:"note"`
	Tags   *[]string `json:"tags" gorm:"type:text[]"`
}

func (e ExpenseModel) TableName() string {
	return "expenses"
}

type ExpeneseRepository interface {
	NewExpenese(expense ExpenseModel) (ExpenseModel, error)
}

type ExpeneseService interface {
	NewExpenese(expense ExpenseModel) (ExpenseModel, error)
}
