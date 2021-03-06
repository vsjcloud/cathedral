package main

import (
	"log"
	"os"

	"cathedral/internal/config"
	"cathedral/internal/pheme"
)

func main() {
	cfg, err := config.LoadConfig(os.Getenv("CATHEDRAL_CONFIG"))
	if err != nil {
		log.Panicf("cannot load config: %v", err)
	}
	app, err := pheme.NewCathedral(cfg)
	if err != nil {
		log.Panicf("cannot instantiate Cathedral: %v", err)
	}
	app.Serve()
}
