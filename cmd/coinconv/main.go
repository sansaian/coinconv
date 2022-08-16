package main

import (
	"github.com/sansaian/coinconv/config"
	"github.com/sansaian/coinconv/internal/app"
	"log"
)

func main() {
	cfg := config.NewFromENV()
	if !cfg.IsValid() {
		log.Fatalf("config has empty required fields")
	}

	app.Run(cfg)
}
