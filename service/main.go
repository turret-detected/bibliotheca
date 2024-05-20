package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/avast/retry-go/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/turret-detected/bibliotheca/service/db"
)

func main() {
	fmt.Println("starting")
	// init connections
	conn, err := retry.DoWithData(db.Connect, retry.Attempts(10), retry.Delay(time.Second))
	if err != nil {
		panic(err)
	}
	defer conn.Close(context.Background())

	err = db.InitDatabase(conn)
	if err != nil {
		panic(err)
	}

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/hello", hello)

	// JS
	e.Static("/assets", "assets")
	e.File("/", "assets/index.html")

	// Start server
	e.Logger.Fatal(e.Start(":8051"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
