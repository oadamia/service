package service

import "github.com/labstack/echo/v4"

// New creates and returns service instance
func New() (e *echo.Echo) {
	e = echo.New()
	return
}
