package main

import (
	mid "garbage-removal-notification-api/middleware"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(mid.CORSMiddleware()))
	e.POST("/sendMessage", sendMessage)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
