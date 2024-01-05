package routes

import (
	"github.com/labstack/echo/v4"
	"go-password/routes/handlers"
)

func ApiRoutes(e *echo.Echo) {
	e.GET("/keynandz/api/generate/password/", handlers.GeneratePassword)
}