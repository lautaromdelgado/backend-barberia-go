package main

import (
	"backend-barberia-go/internal/config"
	"log"
	"os"
)

func main() {
	// Get Config
	if err := config.GetConfig(); err != nil {
		log.Fatalf("Error loading config: %v", err)
		os.Exit(1)
	}
}
