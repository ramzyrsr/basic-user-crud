package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql" //Required for MySQL dialect
	"lemonilo.app/database"
	route "lemonilo.app/router"
)

func main() {

	database.InitDB()

	log.Println("Starting the HTTP server on port 8080")

	router := mux.NewRouter().StrictSlash(true)
	route.InitaliseHandlers(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}
