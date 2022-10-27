package service

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/oadamia/logwrapper"
	"github.com/oadamia/service/validator"
)

var authSecret string

// New creates and returns service instance
func New(secret string) (e *echo.Echo) {
	e = echo.New()
	e.Validator = validator.New()
	e.Logger = logwrapper.Wrapper()
	authSecret = secret
	return
}

// UseJWTMidddleware set JWT middleware for group
func UseJWTMidddleware(g *echo.Group) {
	jwtConfig := middleware.JWTConfig{
		// Claims:     &jwt.UserClaim{},
		SigningKey: []byte(authSecret),
	}
	g.Use(middleware.JWTWithConfig(jwtConfig))
}
