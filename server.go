package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/benzdeus/assessment/db"
	"github.com/benzdeus/assessment/middlewares"
	"github.com/benzdeus/assessment/modules"
	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("Please use server.go for main file")
	fmt.Println("start at port:", os.Getenv("PORT"))

	dbPG, err := db.New()
	if err != nil {
		log.Fatal(err)
	}

	sqlDB, _ := dbPG.DB()
	defer sqlDB.Close()

	e := echo.New()
	// e.Logger.SetLevel(e.Logger.Level())
	e.Use(middlewares.AuthMiddleware)

	modules.InitHandlers(dbPG, e)

	go func() {
		if err := e.Start(os.Getenv("PORT")); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
