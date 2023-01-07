package modules

import (
	"github.com/benzdeus/assessment/entities"
	_expensesRepo "github.com/benzdeus/assessment/modules/expenses/repositories"
	_expensesService "github.com/benzdeus/assessment/modules/expenses/services"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Handler struct {
	expensesService entities.ExpeneseService
}

func InitHandlers(db *gorm.DB, engine *echo.Echo) {
	expensesRepo := _expensesRepo.NewExpeneseRepo(db)
	expensesService := _expensesService.NewExpeneseService(expensesRepo)

	handler := Handler{
		expensesService,
	}

	{
		expensesRoute := engine.Group("/expenses")
		expensesRoute.POST("", handler.NewExpenese)
		expensesRoute.GET("/:id", handler.GetExpenese)
	}

}

func (h Handler) NewExpenese(c echo.Context) error {
	expense := new(entities.ExpenseModel)

	if err := c.Bind(&expense); err != nil {
		return c.JSON(400, err)
	}

	expenseRes, err := h.expensesService.NewExpenese(*expense)
	if err != nil {
		return c.JSON(400, err)

	}
	return c.JSON(201, expenseRes)
}

func (h Handler) GetExpenese(c echo.Context) error {

	id := c.Param("id")

	expense, err := h.expensesService.GetExpenese(id)
	if err != nil {
		return c.JSON(400, err)

	}
	return c.JSON(201, expense)
}
