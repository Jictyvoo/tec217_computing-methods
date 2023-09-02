package main

import (
	"errors"
	"log/slog"
	"math"
	"os"
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
		reportError(err)
		return
	}
	reportResult("Bisection", result, totalInteractions, &buffer)
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
		reportError(err)
		return
	}
	reportResult("FalsePositive", result, totalInteractions, &buffer)
}

func quest03LinearIteration() {
	const (
		x0      = -1
		epsilon = 1e-4
		N       = 1000
	)

	var (
		buffer strings.Builder
		method = methods.LinearIterationMethod{}
	)

	fX := func(x float64) float64 { return math.Pow(x, 2) - (3 * x) + math.Pow(math.E, x) - 2 }
	gX := func(x float64) float64 { return (math.Pow(x, 2) - 2 + math.Pow(math.E, x)) / 3 }

	result, totalInteractions, err := method.Run(
		x0, fX, gX, epsilon, N,
	)

	if err = errors.Join(err, utils.WriteInteractionAsCSV(&buffer, method.InteractionData())); err != nil {
		reportError(err)
		return
	}
	reportResult("LinearIteration", result, totalInteractions, &buffer)
}

func quest04NewtonRaphson() {
	const (
		x0      = -1
		epsilon = 1e-4
		N       = 1000
	)

	var (
		buffer strings.Builder
		method = methods.NewtonRaphsonMethod{}
	)

	fX := func(x float64) float64 { return math.Pow(x, 2) - (3 * x) + math.Pow(math.E, x) - 2 }
	f1X := func(x float64) float64 { return (2 * x) + math.Pow(math.E, x) - 3 }

	result, totalInteractions, err := method.Run(
		x0, fX, f1X, epsilon, N,
	)

	if err = errors.Join(err, utils.WriteInteractionAsCSV(&buffer, method.InteractionData())); err != nil {
		reportError(err)
		return
	}
	reportResult("NewtonRaphson", result, totalInteractions, &buffer)
}

func quest05Secant() {
	const (
		x0, x1  = 0, -1
		epsilon = 1e-4
		N       = 1000
	)

	var (
		buffer strings.Builder
		method = methods.SecantMethod{}
	)

	fX := func(x float64) float64 { return math.Pow(x, 2) - (3 * x) + math.Pow(math.E, x) - 2 }

	result, totalInteractions, err := method.Run(
		x0, x1, fX, epsilon, N,
	)

	if err = errors.Join(err, utils.WriteInteractionAsCSV(&buffer, method.InteractionData())); err != nil {
		reportError(err)
		return
	}
	reportResult("Secant", result, totalInteractions, &buffer)
}

func main() {
	slog.SetDefault(
		slog.New(
			slog.NewTextHandler(
				os.Stdout,
				&slog.HandlerOptions{
					Level: slog.LevelInfo,
				},
			),
		),
	)
	quest01Bisection()
	quest02FalsePositive()
	quest03LinearIteration()
	quest04NewtonRaphson()
	quest05Secant()
}
