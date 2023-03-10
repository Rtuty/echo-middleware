package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"time"
)

func main() {
	s := echo.New()
	s.GET("/status", StatusHandler)

	if err := s.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}

func StatusHandler(ctx echo.Context) error {
	d := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)
	dur := time.Until(d)

	if err := ctx.String(
		http.StatusOK,
		fmt.Sprintf("Number of days: %d", int64(dur.Hours()/24))); //result
	err != nil {
		return err
	}

	return nil
}
