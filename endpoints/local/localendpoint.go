package main

import (
	"log"
	"os"
	"reimanexample/service/api/welcome"
	"reimanexample/service/internal/routing"
)

func main() {
	engine := routing.Build()
	routing.AddRoute(engine, welcome.Path, welcome.Method, welcome.HandleRequest)
	os.Setenv("ALLOWED_ORIGIN", "http://localhost:3000") //set this to the frontend url
	if err := engine.Run(":8080"); err != nil {
		log.Fatal("error starting local server", err)
	}
	log.Println("Server started")
}
