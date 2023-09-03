package main

import (
	"errors"
	"log/slog"
	"math"
	"os"
	"strconv"
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

func quest06PhysicsMadness() {
	const (
		x0      = 3.5
		epsilon = 1e-8
		N       = 1000
	)

	var (
		buffer strings.Builder
		method = methods.NewtonRaphsonMethod{}
	)

	fX := func(x float64) float64 { return math.Tan(0.1*x) - (9.2 * math.Pow(math.E, -x)) }
	f1X := func(x float64) float64 {
		return (9.2 * math.Pow(math.E, -x)) + (0.1 * math.Pow(1/math.Cos(0.1*x), 2))
	}

	result, totalInteractions, err := method.Run(
		x0, fX, f1X, epsilon, N,
	)

	if err = errors.Join(err, utils.WriteInteractionAsCSV(&buffer, method.InteractionData())); err != nil {
		reportError(err)
		return
	}
	reportResult("PhysicsMadness(NewtonRaphson)", result, totalInteractions, &buffer)

	{
		const h = 1e-7 // Step size for finite difference

		// Compute the estimate of the derivative at x = 3 using the finite difference
		f3 := fX(3)
		f3Prime := (fX(3+h) - f3) / h

		// Compute the Derivative Estimate at x = 4 Using the Finite Difference
		f4 := fX(4)
		f4Prime := (fX(4+h) - f4) / h

		slog.Info(
			"Derivative in x = 3",
			slog.String("estimate", strconv.FormatFloat(f3Prime, 'f', -1, 32)),
			slog.String("result", strconv.FormatFloat(f1X(3), 'f', -1, 32)),
		)
		slog.Info(
			"Derivative in x = 4",
			slog.String("estimate", strconv.FormatFloat(f4Prime, 'f', -1, 32)),
			slog.String("result", strconv.FormatFloat(f1X(4), 'f', -1, 32)),
		)
	}
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
	quest06PhysicsMadness()
}
