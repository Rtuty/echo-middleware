package endpoint

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"modules/internal/app/service"
	"net/http"
)

type Endpoint struct {
	service *service.Service
}

func New(s *service.Service) *Endpoint {
	return &Endpoint{service: s}
}

func (e *Endpoint) StatusHandler(ctx echo.Context) error {
	days := e.service.DaysNumber()

	s := fmt.Sprintf("Number of days: %d", days)

	if err := ctx.String(http.StatusOK, s); err != nil {
		return err
	}

	return nil
}
