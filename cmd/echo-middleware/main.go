package main

import (
	"github.com/labstack/gommon/log"
	app2 "modules/internal/pkg/app"
)

func main() {
	app, err := app2.New()

	if err = app.Run(); err != nil {
		log.Fatal(err)
	}
}
