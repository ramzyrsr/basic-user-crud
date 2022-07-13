package database

import (
	"log"

	"github.com/jinzhu/gorm"
)

var Connector *gorm.DB

func Connect(connectString string) error {
	var err error

	Connector, err = gorm.Open("mysql", connectString)
	if err != nil {
		return err
	}
	log.Println("Connection is successfull")
	return nil
}
