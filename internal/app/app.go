package app

import (
	"github.com/sansaian/coinconv/config"
	"github.com/sansaian/coinconv/internal/ui/inputcli"
	"github.com/sansaian/coinconv/internal/ui/outputcli"
	"github.com/sansaian/coinconv/internal/usecase"
	"github.com/sansaian/coinconv/pkg/coinmarketcap"
	"github.com/sansaian/coinconv/pkg/logger"
	"net/http"
)

func Run(cfg *config.Config) {
	log, err := logger.NewWithConfig(cfg.Loglevel)
	if err != nil {
		log.Fatalf("failed initialize logger")
	}

	input := inputcli.New()
	output := outputcli.New()

	client := http.DefaultClient
	converter := coinmarketcap.New(cfg.CoinMarket, client)

	usecase.Convert(log, converter, input, output)
}
