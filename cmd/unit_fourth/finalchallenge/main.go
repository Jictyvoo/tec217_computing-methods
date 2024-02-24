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

type interpolationExecutor[T any] struct {
	method T
	buffer strings.Builder
	result float64
	err    error
}

func quest01NewtonInterpolation() {
	var (
		newton   interpolationExecutor[methods.NewtonInterpolationMethod[float64]]
		lagrange interpolationExecutor[methods.LagrangeInterpolationMethod[float64]]
		x        = []float64{0, 1, 2.5, 3, 4.5, 5, 6}
		y        = []float64{2, 5.4375, 7.3516, 7.5625, 8.4453, 9.1875, 12}
	)
	const xx = 3.5

	newton.result, newton.err = newton.method.Run(
		x, y, xx,
	)
	lagrange.result, lagrange.err = lagrange.method.Run(
		x, y, xx,
	)

	if err := errors.Join(
		newton.err, lagrange.err,
		utils.WriteInteractionAsCSV(&newton.buffer, newton.method.InteractionData()),
		utils.WriteInteractionAsCSV(&lagrange.buffer, lagrange.method.InteractionData()),
	); err != nil {
		views.ReportError(err)
		return
	}

	var buffer strings.Builder
	views.ReportResult("NewtonInterpolation", &buffer, uint32(1), newton.result)
	views.ReportResult("LagrangeInterpolation", &buffer, uint32(1), lagrange.result)
}

func quest02NewtonLagrangeInterpolation() {
	var (
		newton   interpolationExecutor[methods.NewtonInterpolationMethod[float64]]
		lagrange interpolationExecutor[methods.LagrangeInterpolationMethod[float64]]
		x        = []float64{1, 2, 3, 5, 6}
		y        = []float64{7, 4, 5.5, 40, 82}
	)
	const xx = 4

	newton.result, newton.err = newton.method.Run(x, y, xx)
	lagrange.result, lagrange.err = lagrange.method.Run(x, y, xx)

	if err := errors.Join(
		newton.err, lagrange.err,
		utils.WriteInteractionAsCSV(&newton.buffer, newton.method.InteractionData()),
		utils.WriteInteractionAsCSV(&lagrange.buffer, lagrange.method.InteractionData()),
	); err != nil {
		views.ReportError(err)
		return
	}

	var buffer strings.Builder
	views.ReportResult("NewtonInterpolation", &buffer, uint32(1), newton.result)
	views.ReportResult("LagrangeInterpolation", &buffer, uint32(1), lagrange.result)
}

func quest03Bisection() {
	var (
		a, b   float64 = 2, 3
		buffer strings.Builder
		method = methods.BisectionMethod{}
	)

	result, totalInteractions, err := method.Run(
		a, b,
		func(x float64) float64 {
			return 4.62 - (2.13 * x) + (0.42 * math.Pow(x, 2)) - (0.03 * math.Pow(x, 3)) - 1.7
		},
		10e-5, 1000,
	)

	if err = errors.Join(err, utils.WriteInteractionAsCSV(&buffer, method.InteractionData())); err != nil {
		views.ReportError(err)
		return
	}
	views.ReportResult("Bisection", &buffer, totalInteractions, result)
}

func quest04LinearInterpolationLetterB() {
	var (
		buffer strings.Builder
		method = methods.NewtonInterpolationMethod[float64]{}
	)

	result, err := method.Run(
		[]float64{10, 20},
		[]float64{9.672, 8.608},
		15,
	)

	if err = errors.Join(err, utils.WriteInteractionAsCSV(&buffer, method.InteractionData())); err != nil {
		views.ReportError(err)
		return
	}
	views.ReportResult("Linear Interpolation", &buffer, uint32(1), result)
}

func quest05NewtonInterpolation() {
	var (
		buffer strings.Builder
		method = methods.NewtonInterpolationMethod[float64]{}
	)

	result, err := method.Run(
		[]float64{0.25, 0.75, 1.25, 1.5, 2},
		[]float64{-0.45, -0.6, 0.7, 1.88, 6},
		1.15,
	)

	if err = errors.Join(err, utils.WriteInteractionAsCSV(&buffer, method.InteractionData())); err != nil {
		views.ReportError(err)
		return
	}
	views.ReportResult("Newton Interpolation", &buffer, uint32(1), result)
}

func quest06BisectionLetterC() {
	var (
		a, b   float64 = 2, 3
		buffer strings.Builder
		method = methods.BisectionMethod{}
	)

	result, totalInteractions, err := method.Run(
		a, b,
		func(x float64) float64 {
			return (0.0063 * math.Pow(x, 3)) - (0.08643 * math.Pow(x, 2)) + (0.4118 * x) + 0.2715 - 0.85
		},
		10e-5, 1000,
	)

	if err = errors.Join(err, utils.WriteInteractionAsCSV(&buffer, method.InteractionData())); err != nil {
		views.ReportError(err)
		return
	}
	views.ReportResult("Bisection", &buffer, totalInteractions, result)
}

func quest07NewtonInterpolation() {
	var (
		buffer strings.Builder
		method = methods.NewtonInterpolationMethod[float64]{}
	)

	result, err := method.Run(
		[]float64{-1, -0.5, -0.25, 0.25, 0.5, 1},
		[]float64{-637, -96.5, -20.5, 20.5, 96.5, 637},
		0.1,
	)

	if err = errors.Join(err, utils.WriteInteractionAsCSV(&buffer, method.InteractionData())); err != nil {
		views.ReportError(err)
		return
	}
	views.ReportResult("Newton Interpolation", &buffer, uint32(1), result)
}

func quest08NewtonInterpolation() {
	var (
		buffer strings.Builder
		method = methods.NewtonInterpolationMethod[float64]{}
	)

	result, err := method.Run(
		[]float64{0, 30000, 60000, 90000, 120000},
		[]float64{9.8100, 9.7487, 9.6879, 9.6278, 9.5682},
		55000,
	)

	if err = errors.Join(err, utils.WriteInteractionAsCSV(&buffer, method.InteractionData())); err != nil {
		views.ReportError(err)
		return
	}
	views.ReportResult("Newton Interpolation", &buffer, uint32(1), result)
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

	quest01NewtonInterpolation()
	quest02NewtonLagrangeInterpolation()
	quest03Bisection()
	quest04LinearInterpolationLetterB()
	quest05NewtonInterpolation()
	quest06BisectionLetterC()
	quest07NewtonInterpolation()
	quest08NewtonInterpolation()
}
