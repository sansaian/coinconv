package usecase

import (
	"fmt"
	"github.com/sansaian/coinconv/internal/entities"
)

type UseCase struct {
	converter Converter
}

func New(converter Converter) *UseCase {
	return &UseCase{
		converter: converter,
	}
}

func (u *UseCase) Convert(input InputReader) (*entities.ConvertingResult, error) {

	data, err := input.GetData()
	if err != nil {
		return nil, fmt.Errorf("failed to get correct data for convert %w", err)
	}
	return u.converter.GetConvertingPrice(data)
}
