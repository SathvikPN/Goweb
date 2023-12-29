package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SathvikPN/Goweb/router"
	"github.com/SathvikPN/Goweb/services"
)

func main() {
	err := services.InitDB()
	if err != nil {
		log.Fatal("DB Init fail", err)
	}

	r := router.Router()

	fmt.Println("Starting REST server on port 8080 ...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
