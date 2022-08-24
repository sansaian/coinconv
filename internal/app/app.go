package app

import (
	"github.com/sansaian/coinconv/internal/ui/inputcli"
	"github.com/sansaian/coinconv/internal/ui/outputcli"
	"github.com/sansaian/coinconv/internal/usecase"
	"github.com/sirupsen/logrus"
)

func Run(log *logrus.Logger, useCaseConverter usecase.UseCaseConverter) {
	input := inputcli.New()
	output := outputcli.New()
	result, err := useCaseConverter.Convert(input)
	if err != nil {
		log.WithError(err)
		return
	}
	err = output.PrintData(result)
	if err != nil {
		log.WithError(err)
		return
	}
}
