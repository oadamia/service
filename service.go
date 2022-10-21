package service

import (
	"github.com/labstack/echo/v4"
	"github.com/oadamia/service/validator"
)

// New creates and returns service instance
func New() (e *echo.Echo) {
	e = echo.New()
	e.Validator = validator.New()
	return
}
