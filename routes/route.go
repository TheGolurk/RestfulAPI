package routes

import (
	"github.com/TheGolurk/RestfulAPI/user"
	"github.com/labstack/echo"
)

// StartRoutes Inicializa las rutas
func StartRoutes(e *echo.Echo) {
	e.GET("api/v1/users", user.GetAll)   //GetAll Users
	e.GET("api/v1/user", user.Get)       //GET one user
	e.POST("api/v1/user", user.Create)   //CREATE
	e.PUT("api/v1/user", user.Update)    //UPDATE
	e.DELETE("api/v1/user", user.Delete) //DELETE
}
