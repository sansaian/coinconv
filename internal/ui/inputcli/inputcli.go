package inputcli

import (
	"fmt"
	"github.com/sansaian/coinconv/internal/entities"
	"os"
	"strconv"
)

const (
	limitAmount  = 1000000000000.0
	limitCLIArgs = 3
)

type InputCLI struct {
}

func New() *InputCLI {
	return &InputCLI{}
}

func (cli *InputCLI) GetData() (*entities.InputData, error) {

	args := os.Args[1:]
	if len(args) < limitCLIArgs {
		return nil, fmt.Errorf("insufficient number of arguments: want=%v, got=%v", limitCLIArgs, len(args))
	}

	amount, err := parseAmount(args[0])
	if err != nil {
		return nil, fmt.Errorf("can't parse amount: %w", err)
	}

	from, err := validateCurrency(args[1])
	if err != nil {
		return nil, fmt.Errorf("can't parse currency from %w", err)
	}
	to, err := validateCurrency(args[2])
	if err != nil {
		return nil, fmt.Errorf("can't parse currency from: %w", err)
	}

	return &entities.InputData{Amount: amount, From: from, To: to}, nil
}

func parseAmount(amount string) (float64, error) {
	parsed, err := strconv.ParseFloat(amount, 10)
	if err != nil {
		return 0.0, fmt.Errorf("bad amount value: %w", err)
	}
	if parsed > limitAmount {
		return 0.0, fmt.Errorf("amount must be less than or equal to 1000000000000: %f", limitAmount)
	}
	return parsed, nil
}

func validateCurrency(currency string) (string, error) {
	if currency == "" || currency == " " {
		return "", fmt.Errorf("empty currency")
	}
	return currency, nil
}
