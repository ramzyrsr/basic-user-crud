package router

import (
	"github.com/gorilla/mux"
	"lemonilo.app/controller"
	"lemonilo.app/middleware"
)

func InitaliseHandlers(router *mux.Router) {
	// User Routes
	router.HandleFunc("/user", middleware.SetMiddlewareJSON(controller.CreateUser)).Methods("POST")
	router.HandleFunc("/users", middleware.SetMiddlewareJSON(controller.GetAllUser)).Methods("GET")
	router.HandleFunc("/user/{userId}", middleware.SetMiddlewareJSON(controller.GetUserById)).Methods("GET")
	router.HandleFunc("/user/{userId}", middleware.SetMiddlewareJSON(middleware.SetMiddlewareAuthentication(controller.UpdateUserById))).Methods("PATCH")
	router.HandleFunc("/user/{userId}", middleware.SetMiddlewareJSON(middleware.SetMiddlewareAuthentication(controller.DeletUserById))).Methods("DELETE")

	// Login Route
	router.HandleFunc("/login", controller.Login).Methods("POST")
}
