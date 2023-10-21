package main

import (
	"log"
	"refactoring/config"
	"refactoring/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %v", err)
	}

	app.Run(cfg)
}
