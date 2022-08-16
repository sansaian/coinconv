package usecase

import "github.com/sirupsen/logrus"

func Convert(log *logrus.Logger, converter Converter, input InputReader, output DataPrinter) {

	data, err := input.GetData()
	if err != nil {
		log.WithError(err).Errorf("failed to get correct data for convert")
		return
	}
	result, err := converter.GetConvertingPrice(data)
	if err != nil {
		log.WithError(err).Errorf("failed to converted currency")
		return
	}
	err = output.PrintData(result)
	if err != nil {
		log.WithError(err).Errorf("failed to print data")
		return
	}
}
