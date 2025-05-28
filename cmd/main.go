// @title Car Service API
// @version 1.0
// @description API for managing vehicles
// @host localhost:8080
// @BasePath /
package main

import (
	"car-service/internal/app"
	"log"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatalf("app terminated: %v", err)
	}
}
