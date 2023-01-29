package main

import (
	"log"
	"reimanexample/service/api/welcome"
	"reimanexample/service/internal/routing"
)

func main() {
	engine := routing.Build()
	routing.AddRoute(engine, welcome.Path, welcome.Method, welcome.HandleRequest)
	if err := engine.Run(":8080"); err != nil {
		log.Fatal("error starting local server", err)
	}
	log.Println("Server started")
}
