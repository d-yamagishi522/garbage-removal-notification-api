package main

import (
	mid "garbage-removal-notification-api/middleware"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(mid.CORSMiddleware()))
	e.POST("/sendMessage", sendMessage)

	e.Logger.Fatal(e.Start(":8080"))
}
