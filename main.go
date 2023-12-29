package main

import (
	"log"

	"github.com/SathvikPN/Goweb/services"
)

func main() {
	err := services.InitDB()
	if err != nil {
		log.Fatal("DB Init fail", err)
	}
}
