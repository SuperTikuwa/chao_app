package handler

import (
	"math/rand"

	"github.com/labstack/echo/v4"
)

type RollResponse struct {
	Eyes int `json:"eyes"`
}

func Roll(c echo.Context) error {
	eyes := rand.Intn(6)
	if eyes == 5 {
		eyes = 0
	}

	res := RollResponse{Eyes: eyes}
	return c.JSON(200, res)
}
