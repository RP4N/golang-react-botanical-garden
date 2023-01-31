package main

import (
	"log"
	"reimanexample/gardenService/api/welcome"
	"reimanexample/gardenService/internal/routing"
)

func main() {
	//os.Setenv("ALLOWED_ORIGIN", "http://localhost:3000") //set this to the frontend url
	engine := routing.Build()
	routing.AddRoute(engine, welcome.Path, welcome.Method, welcome.HandleRequest)

	if err := engine.Run(":8080"); err != nil {
		log.Fatal("error starting local server", err)
	}
	log.Println("Server started")
}
