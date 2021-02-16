package middleware

import (
	"os"
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// CORSMiddleware init cors config
func CORSMiddleware() middleware.CORSConfig {
	origin := os.Getenv("ORIGIN")
	config := middleware.CORSConfig{
		AllowOrigins: strings.Split(origin, "|"),
		AllowMethods: []string{echo.POST},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}
	return config
}
