package user

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/TheGolurk/RestfulAPI/configuration"
	"github.com/TheGolurk/RestfulAPI/response"
	"github.com/labstack/echo"
)

// Create crea un nuevo usuario
func Create(c echo.Context) error {
	u := &User{}
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

	db := configuration.GetConnection()
	defer db.Close()

	err = db.Create(&u).Error
	if err != nil {
		r := response.Model{
			Code:    "500",
			Message: "Error al crear",
			Data:    err,
		}

		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().WriteHeader(http.StatusInternalServerError)
		return json.NewEncoder(c.Response()).Encode(r)
	}

	r := response.Model{
		Code:    "201",
		Message: "Creado Correctamente",
		Data:    u,
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusCreated)
	return json.NewEncoder(c.Response()).Encode(r)
}

// GetAll Obtiene todos los datos
func GetAll(c echo.Context) error {
	users := []User{}
	db := configuration.GetConnection()
	defer db.Close()

	us := db.Find(&users)

	r := response.Model{
		Code:    "200",
		Message: "Consultado Correctamente",
		Data:    us,
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(r)

}

// Delete elimina un usuario por su id
func Delete(c echo.Context) error {

	var usuario User
	id := c.QueryParam("id")

	db := configuration.GetConnection()
	defer db.Close()

	db.First(&usuario, id)
	db.Delete(&usuario)

	users := []User{}
	db.Find(&users)

	r := response.Model{
		Code:    "202",
		Message: "Eliminado Correctamente",
		Data:    &users,
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusAccepted)
	return json.NewEncoder(c.Response()).Encode(r)
}

// Update actualiza los campos
func Update(c echo.Context) error {
	// var usuario User

	i := c.QueryParam("id")
	fn := c.QueryParam("firstname")
	ltn := c.QueryParam("lastname")
	ag, err := strconv.Atoi(c.QueryParam("age"))
	if err != nil {
		log.Fatal("Error al convertir el parametro ", err)
		return err
	}

	db := configuration.GetConnection()
	defer db.Close()

	// db.Find(&usuario, i)

	db.Model(&User{}).Where("ID = ? ", i).Updates(User{FirstName: fn, LastName: ltn, Age: ag})

	users := []User{}
	db.Find(&users)

	r := response.Model{
		Code:    "202",
		Message: "Actualizado Correctamente",
		Data:    &users,
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusAccepted)
	return json.NewEncoder(c.Response()).Encode(r)

}

// Get trae un solo usuario por su ID
func Get(c echo.Context) error {
	id := c.QueryParam("id")

	db := configuration.GetConnection()
	defer db.Close()

	consulta := db.First(&User{}, id)

	r := response.Model{
		Code:    "200",
		Message: "Consultado correctamente",
		Data:    consulta,
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusAccepted)
	return json.NewEncoder(c.Response()).Encode(r)

}
