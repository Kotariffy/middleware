package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/status", Handler)

	err := e.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

func Handler(ctx echo.Context) error {

	d := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)

	duration := time.Until(d)

	str := fmt.Sprintf("Days until 01.01.2025: %d", int64(duration.Hours())/24)

	err := ctx.String(http.StatusOK, str)
	if err != nil {
		return err
	}

	return nil
}
