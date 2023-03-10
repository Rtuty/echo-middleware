package app

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"modules/internal/app/endpoint"
	"modules/internal/app/service"
	"modules/internal/middleware"
)

type App struct {
	srvc *service.Service
	endp *endpoint.Endpoint
	echo *echo.Echo
}

func New() (*App, error) {
	app := &App{}

	app.srvc = service.New()
	app.endp = endpoint.New(app.srvc)

	app.echo = echo.New()
	app.echo.Use(middleware.RoleChecker)
	app.echo.GET("/status", app.endp.StatusHandler)

	return app, nil
}

func (app *App) Run() error {
	fmt.Println("server running\nhttp://localhost:8080")

	if err := app.echo.Start(":8080"); err != nil {
		log.Fatal(err)
	}

	return nil
}
