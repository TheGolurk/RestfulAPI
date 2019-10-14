package main

import (
	"github.com/TheGolurk/RestfulApi/user"
	"github.com/labstack/echo"
)

// StartRoutes Inicializa las rutas
func StartRoutes(e *echo.Echo) {
	e.GET("api/v1/user", user.Create)
}
