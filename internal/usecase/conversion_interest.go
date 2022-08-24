package usecase

import (
	"github.com/sansaian/coinconv/config"
	"github.com/sansaian/coinconv/internal/entities"
)

type UseCaseInterest struct {
	cfg              *config.Config
	usecaseconverter UseCaseConverter
}

func NewWithInterest(cfg *config.Config, usecaseconverter UseCaseConverter) *UseCaseInterest {
	return &UseCaseInterest{
		cfg:              cfg,
		usecaseconverter: usecaseconverter,
	}
}

func (u *UseCaseInterest) Convert(input InputReader) (*entities.ConvertingResult, error) {
	withOutInterest, err := u.usecaseconverter.Convert(input)
	if err != nil {
		return nil, err
	}
	result := withOutInterest.Result + (withOutInterest.Result * u.cfg.Interest / 100)

	return &entities.ConvertingResult{Result: result}, nil
}
