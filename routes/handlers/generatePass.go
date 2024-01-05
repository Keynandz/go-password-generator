package handlers

import (
	"go-password/repositories"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GeneratePassword(e echo.Context) error {
	limit, _ := strconv.Atoi(e.QueryParam("limit"))
	up, _ := strconv.Atoi(e.QueryParam("up"))
	low, _ := strconv.Atoi(e.QueryParam("low"))
	symbol, _ := strconv.Atoi(e.QueryParam("symbol"))
	digit, _ := strconv.Atoi(e.QueryParam("digit"))

	password, err := repositories.GeneratePassword(limit, up, low, digit, symbol)
	if err != nil {
		return err
	}

	strength := repositories.PasswordStrength(password)

	return e.JSON(http.StatusOK, map[string]interface{}{
		"result":      password,
		"strength":    strength,
		"status_code": http.StatusOK,
	})
}
