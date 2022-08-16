package outputcli

import (
	"fmt"
	"github.com/sansaian/coinconv/internal/entities"
)

type OutputCLI struct {
}

func New() *OutputCLI {
	return &OutputCLI{}
}

func (cli *OutputCLI) PrintData(result *entities.ConvertingResult) error {
	fmt.Println("Result =", result.Result)
	return nil
}
