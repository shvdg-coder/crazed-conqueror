package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

// TODO: Implement actual API server.

// the main is the entry point of the API server.
func main() {
	fmt.Println("Starting up API server...")

	ech := echo.New()
	ech.Logger.SetOutput(os.Stdout)
	ech.Logger.SetLevel(log.DEBUG)
	ech.Use(configureCORS())

	ech.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Echo server is running!")
	})

	ech.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})

	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8080"
	}

	address := ":" + port
	fmt.Printf("http server started on %s\n", address)

	if err := ech.Start(address); err != nil {
		ech.Logger.Fatal("failed to start server: ", err)
	}
}

// configureCORS returns a CORS middleware configuration.
func configureCORS() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.POST, echo.DELETE, echo.OPTIONS, echo.PATCH},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
			echo.HeaderXCSRFToken,
		},
		AllowCredentials: true,
		ExposeHeaders:    []string{echo.HeaderContentLength, echo.HeaderContentType},
	})
}
