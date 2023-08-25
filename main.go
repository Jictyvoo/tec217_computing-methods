package main

import (
	"errors"
	"fmt"
	"log/slog"
	"math"
	"strings"

	"github.com/jictyvoo/tec217_computing-methods/internal/methods"
	"github.com/jictyvoo/tec217_computing-methods/internal/utils"
)

func quest01Bisection() {
	var (
		a, b   float64 = 0.5, 1
		buffer strings.Builder
		method = methods.BisectionMethod{}
	)

	result, totalInteractions, err := method.Run(
		a, b, func(x float64) float64 { return math.Sin(x) - math.Pow(x, 2) }, 0.02, 1000,
	)

	if err = errors.Join(err, utils.WriteInteractionAsCSV(&buffer, method.InteractionData())); err != nil {
		slog.Error("Failed to find root", slog.String("error", err.Error()))
		return
	}
	slog.Info(
		fmt.Sprintf(
			"The value of root is : %.5f. After %d iteractions\n",
			result, totalInteractions,
		),
		slog.String("output-table", buffer.String()),
	)

}

func quest02FalsePositive() {
	const (
		a, b    float64 = 0, 1
		epsilon         = 5e-4
		N               = 1000
	)

	var (
		buffer strings.Builder
		method = methods.FalsePositionMethod{}
	)

	result, totalInteractions, err := method.Run(
		a, b, func(x float64) float64 { return math.Pow(x, 3) - (9 * x) + 3 }, epsilon, N,
	)

	if err = errors.Join(err, utils.WriteInteractionAsCSV(&buffer, method.InteractionData())); err != nil {
		slog.Error("Failed to find root", slog.String("error", err.Error()))
		return
	}
	slog.Info(
		fmt.Sprintf(
			"The value of root is : %.5f. After %d iteractions\n",
			result, totalInteractions,
		),
		slog.String("output-table", buffer.String()),
	)
}

func main() {
	quest01Bisection()
	quest02FalsePositive()
}
