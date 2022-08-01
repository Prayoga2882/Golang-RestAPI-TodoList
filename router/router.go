package router

import (
	"github.com/gorilla/mux"
	"main/controller"
	"main/middleware"
	"main/model"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/signup", model.SignUp).Methods("POST")
	router.HandleFunc("/api/signin", model.SignIn).Methods("POST")

	router.HandleFunc("/api/todo", middleware.IsAuthorized(controller.FindAllTodo)).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/todo/{id}", middleware.IsAuthorized(controller.FindTodo)).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/todo", middleware.IsAuthorized(controller.CreateTodo)).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/todo/{id}", middleware.IsAuthorized(controller.UpdateTodo)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/todo/{id}", middleware.IsAuthorized(controller.DeleteTodo)).Methods("DELETE", "OPTIONS")

	return router
}
