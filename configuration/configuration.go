package configuration

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" //driver para mysql
	"github.com/jinzhu/gorm"
)

// Configuration crea un struct para el json
type Configuration struct {
	Server   string
	Port     string
	User     string
	Password string
	Database string
}

// GetConfiguration obtiene la configuración del json
func GetConfiguration() Configuration {
	var c Configuration
	file, err := os.Open("configuration.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&c)
	if err != nil {
		log.Fatal(err)
	}

	return c
}

// GetConnection obtiene una conexion a la bd
func GetConnection() *gorm.DB {
	c := GetConfiguration()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", c.User, c.Password, c.Server, c.Port, c.Database)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
