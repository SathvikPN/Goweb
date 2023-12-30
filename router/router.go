package router

import (
	"github.com/SathvikPN/Goweb/middlewares"
	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/newpost", middlewares.CreatePost).Methods("POST")
	router.HandleFunc("/api/post/{id}", middlewares.GetPost).Methods("GET")

	return router
}
