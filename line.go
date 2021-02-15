package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func sendMessage(ctx echo.Context) error {
	return ctx.NoContent(http.StatusNoContent)
}
