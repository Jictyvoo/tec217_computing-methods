package main

import (
	"errors"
	"log/slog"
	"math"
	"os"
	"strings"

	"github.com/jictyvoo/tec217_computing-methods/internal/methods"
	"github.com/jictyvoo/tec217_computing-methods/internal/utils"
	"github.com/jictyvoo/tec217_computing-methods/pkg/views"
)

func quest01() {
	var (
		buffer strings.Builder
		method = methods.BisectionMethod{}
		fX     = func(x float64) float64 { return (math.Pow(x, 3) * math.Cos(2*x)) - (3 / (7 * math.Pi)) }
	)

	// Root 1
	{
		const (
			a1, b1  = -1.5, 0
			epsilon = 1e-2
			N       = 9
		)

		result, totalInteractions, err := method.Run(
			a1, b1, fX, epsilon, N,
		)

		if err = errors.Join(err, utils.WriteInteractionAsCSV(&buffer, method.InteractionData())); err != nil {
			views.ReportError(err)
			return
		}
		views.ReportResult("Question 01", result, totalInteractions, &buffer)
		buffer.Reset()
	}

	method.Reset()

	// Root 2
	{
		const (
			a2, b2  = -3, -2
			epsilon = 1e-2
			N       = 9
		)

		result, totalInteractions, err := method.Run(
			a2, b2, fX, epsilon, N,
		)

		if err = errors.Join(err, utils.WriteInteractionAsCSV(&buffer, method.InteractionData())); err != nil {
			views.ReportError(err)
			return
		}
		views.ReportResult("Question 01", result, totalInteractions, &buffer)
		buffer.Reset()
	}
}

func quest02() {
	var (
		buffer strings.Builder
		method = methods.FalsePositionMethod{}
		fX     = func(x float64) float64 { return math.Pow(x, 5) + x - 5 }
	)

	// Root 1
	{
		const (
			a1, b1  = 1, 1.5
			epsilon = 1e-4
			N       = 17
		)

		result, totalInteractions, err := method.Run(
			a1, b1, fX, epsilon, N,
		)

		if err = errors.Join(err, utils.WriteInteractionAsCSV(&buffer, method.InteractionData())); err != nil {
			views.ReportError(err)
			return
		}
		views.ReportResult("Question 02", result, totalInteractions, &buffer)
		buffer.Reset()
	}
}

func quest03() {
	const (
		a1, b1  = 0.1, 0.15
		epsilon = 1e-3
		N       = 10

		// Function constants
		g = 9.81
		m = 17
		v = 36
		t = 6
	)

	var (
		buffer strings.Builder
		method = methods.SecantMethod{}
		fX     = func(x float64) float64 { return (((g * m) / x) * (1 - math.Pow(math.E, (-x*t)/m))) - v }
	)

	result, totalInteractions, err := method.Run(
		a1, b1, fX, epsilon, N,
	)

	if err = errors.Join(err, utils.WriteInteractionAsCSV(&buffer, method.InteractionData())); err != nil {
		views.ReportError(err)
		return
	}
	views.ReportResult("Question 03", result, totalInteractions, &buffer)
	buffer.Reset()
}

func quest04() {
	const (
		x0      = 6
		epsilon = 1e-3
		N       = 10

		// Function constants
		y = 10
	)

	var (
		buffer strings.Builder
		method = methods.NewtonRaphsonMethod{}
		fX     = func(x float64) float64 { return (7 * (2 - math.Pow(0.9, x))) - y }
		f1X    = func(x float64) float64 { return 0.737524 * math.Pow(0.9, x) }
	)

	result, totalInteractions, err := method.Run(
		x0, fX, f1X, epsilon, N,
	)

	if err = errors.Join(err, utils.WriteInteractionAsCSV(&buffer, method.InteractionData())); err != nil {
		views.ReportError(err)
		return
	}
	views.ReportResult("Question 04", result, totalInteractions, &buffer)
	buffer.Reset()
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
	quest01()
	quest02()
	quest03()
	quest04()
}
