package database

import (
	"log"

	"github.com/jinzhu/gorm"
)

var Connector *gorm.DB

func Connect(connectionString string) error {
	var err error
	Connector, err = gorm.Open("mysql", connectionString)

	if err != nil {
		return err
	}
	log.Println("Connection was successful")
	return nil
}
