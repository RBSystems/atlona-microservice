package main

import (
	//"log"
	"net/http"

	"github.com/byuoitav/atlona-microservice/handlers"
	"github.com/byuoitav/atlona-microservice/handlersmatrix"
	"github.com/byuoitav/authmiddleware"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	//"time"
)

func main() {
	port := ":8015"
	router := echo.New()
	router.Pre(middleware.RemoveTrailingSlash())
	router.Use(middleware.CORS())

	// Use the `secure` routing group to require authentication
	secure := router.Group("", echo.WrapMiddleware(authmiddleware.Authenticate))

	// Functionality endpoints for Atlona Video over IP Switching
	secure.GET("/:address/input/:input/:output", handlers.SwitchInput)
	secure.GET("/:address/status/:output", handlers.CheckInput)
	secure.GET("/:address/allstatus", handlers.AllInputs)

	// Functionality endpoints for Atlona Standard Switch
	secure.GET("/switch/:address/input/:input/:output", handlersmatrix.SwitchInput)
	secure.GET("/switch/:address/status/:output", handlersmatrix.CheckInput)
	secure.GET("/switch/:address/allstatus", handlersmatrix.AllInputs)

	// Status endpoints

	server := http.Server{
		Addr:           port,
		MaxHeaderBytes: 1024 * 10,
	}

	router.StartServer(&server)
}
