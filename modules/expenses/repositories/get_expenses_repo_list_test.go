package repositories_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/benzdeus/assessment/db"
	"github.com/benzdeus/assessment/middlewares"
	"github.com/benzdeus/assessment/modules"
	"github.com/benzdeus/assessment/modules/expenses/repositories"
	"github.com/labstack/echo/v4"
)

func TestGetExpensesList(t *testing.T) {
	dbPG, err := db.New()
	if err != nil {
		log.Fatal(err)
	}

	sqlDB, _ := dbPG.DB()
	defer sqlDB.Close()

	e := echo.New()
	e.Use(middlewares.AuthMiddleware)

	modules.InitHandlers(dbPG, e)

	go func() {
		if err := e.Start(os.Getenv("PORT")); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	repo := repositories.NewExpeneseRepo(dbPG)

	res, err := repo.GetExpeneseList()
	if err != nil {
		t.Error(err)
	}

	log.Printf("%#+v", res)

	defer e.Close()
}
