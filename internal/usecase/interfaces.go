package usecase

import "github.com/sansaian/coinconv/internal/entities"

type Converter interface {
	GetConvertingPrice(data *entities.InputData) (*entities.ConvertingResult, error)
}

type InputReader interface {
	GetData() (*entities.InputData, error)
}

type DataPrinter interface {
	PrintData(*entities.ConvertingResult) error
}
