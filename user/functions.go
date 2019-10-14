package user

import (
	"encoding/json"
	"net/http"

	"github.com/TheGolurk/RestfulApi/response"
	"github.com/labstack/echo"
)

// Create crea un nuevo usuario
func Create(c echo.Context) error {
	u := &Model{}
	err := c.Bind(u)
	if err != nil {
		r := response.Model{
			Code:    "400",
			Message: "Estructura incorrecta",
			Data:    err,
		}

		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().WriteHeader(http.StatusBadRequest)
		return json.NewEncoder(c.Response()).Encode(r)
	}

	data := storage.Create(u)
	r := response.Model{
		Code:    "201",
		Message: "Creado Correctamente",
		Data:    data,
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusCreated)
	return json.NewEncoder(c.Response()).Encode(r)

}
