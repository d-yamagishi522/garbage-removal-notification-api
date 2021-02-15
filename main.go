package main

import (
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.POST("/sendMessage", sendMessage)

	e.Logger.Fatal(e.Start(":8080"))
}
