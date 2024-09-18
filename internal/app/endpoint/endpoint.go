package endpoint

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type Service interface {
	DaysLeft() int64
}

type Endpoint struct {
	s Service
}

func New(s Service) *Endpoint {
	return &Endpoint{
		s: s,
	}
}

func (e *Endpoint) Status(ctx echo.Context) error {
	days := e.s.DaysLeft()

	resp := fmt.Sprintf("Days until 01.01.2025: %d", days)

	err := ctx.String(http.StatusOK, resp)
	if err != nil {
		return err
	}

	return nil
}
