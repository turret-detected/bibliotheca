package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/avast/retry-go/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/samber/lo"
	"github.com/turret-detected/bibliotheca/service/db"
)

func main() {
	// init connections
	retry.Do(func() error {
		err, ok := lo.TryWithErrorValue(func() error {
			db.Connect()
			return nil
		})
		if !ok {
			return fmt.Errorf("could not connect: %v", err)
		}
		return nil
	}, retry.Attempts(10), retry.Delay(time.Second))
	defer db.Connection.Close(context.Background())

	// TODO do with retry
	// files.Connect()

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)

	// Start server
	e.Logger.Fatal(e.Start(":8051"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
