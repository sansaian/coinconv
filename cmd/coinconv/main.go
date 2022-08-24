package main

import (
	"github.com/sansaian/coinconv/config"
	"github.com/sansaian/coinconv/internal/app"
	"github.com/sansaian/coinconv/internal/usecase"
	"github.com/sansaian/coinconv/pkg/coinmarketcap"
	"github.com/sansaian/coinconv/pkg/logger"
	"log"
	"net/http"
)

func main() {
	cfg := config.NewFromENV()
	if !cfg.IsValid() {
		log.Fatalf("config has empty required fields")
	}
	log, err := logger.NewWithConfig(cfg.Loglevel)
	if err != nil {
		log.Fatalf("failed initialize logger")
	}
	client := http.DefaultClient
	converter := coinmarketcap.New(cfg.CoinMarket, client)

	useConv := usecase.New(converter)
	useConvWithInterest := usecase.NewWithInterest(cfg, useConv)
	app.Run(log, useConvWithInterest)
}
