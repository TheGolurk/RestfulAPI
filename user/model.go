package user

import "github.com/jinzhu/gorm"

// User es la estructura para el usuario
type User struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int32  `json:"age"`
}

// var storage Storage

// // Storage es un mapa del modelo
// type Storage map[string]*Model

// Add crea un nuevo usuario
// func Add(m *Model) *Model {
// 	db := database.Init()
// 	db.NewRecord(m)
// 	db.Create(m)
// 	return m
// }

// // GetAll Retorna todos los usuarios
// func (s Storage) GetAll() *Model {
// 	return m
// }
