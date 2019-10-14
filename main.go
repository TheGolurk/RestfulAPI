package main

import (
	"fmt"

	"github.com/TheGolurk/RestfulAPI/routes"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	routes.StartRoutes(e)

	err := e.Start("192.168.0.4:5000")
	if err != nil {
		fmt.Printf("Error, no se pudo ejecutar el servidor: %v", err)
	}
}
