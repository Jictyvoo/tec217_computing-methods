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
	"github.com/jictyvoo/tec217_computing-methods/pkg/views"
)

func quest01NewtonInterpolation() {
	var (
		buffer strings.Builder
		method = methods.NewtonInterpolationMethod[float64]{}
	)

	result, err := method.Run(
		[]float64{0, 0.5, 1},
		[]float64{1, 2.12, 3.55},
		0.7,
	)

	if err = errors.Join(err, utils.WriteInteractionAsCSV(&buffer, method.InteractionData())); err != nil {
		views.ReportError(err)
		return
	}
	views.ReportResult("NewtonInterpolation", &buffer, uint32(1), result)
}

func quest02LagrangeInterpolation() {
	var (
		buffer strings.Builder
		method = methods.LagrangeInterpolationMethod[float64]{}
	)

	result, err := method.Run(
		[]float64{0, 0.5, 1},
		[]float64{1.3, 2.5, 0.9},
		0.8,
	)

	if err = errors.Join(err, utils.WriteInteractionAsCSV(&buffer, method.InteractionData())); err != nil {
		views.ReportError(err)
		return
	}
	views.ReportResult("LagrangeInterpolation", &buffer, uint32(1), result)
}

func quest03VandermondeInterpolation() {
	method := methods.GaussJordanMethod[float64]{}

	var questions struct {
		one, two struct {
			det        float64
			foundRoots []float64
			err        error
		}
	}

	questions.one.det, questions.one.foundRoots, questions.one.err = method.Run(
		[][]float64{
			{0, 0, 1, 1},
			{0.25, 0.5, 1, 2.12},
			{1, 1, 1, 3.55},
		},
	)

	questions.two.det, questions.two.foundRoots, questions.two.err = method.Run(
		[][]float64{
			{0, 0, 1, 1.3},
			{0.25, 0.5, 1, 2.5},
			{1, 1, 1, 0.9},
		},
	)

	err := errors.Join(
		questions.one.err,
		questions.two.err,
	)
	if err != nil {
		views.ReportError(err)
		return
	}
	views.ReportLinearSystemResult(
		"Vandermonde - GaussJordan [Question 01]",
		questions.one.det,
		questions.one.foundRoots,
		nil,
	)
	views.ReportLinearSystemResult(
		"Vandermonde - GaussJordan [Question 02]",
		questions.two.det,
		questions.two.foundRoots,
		nil,
	)
}

func quest04MinimumSquareDirectly() {
	x := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	y := []float64{1.3, 3.5, 4.2, 5.0, 7.0, 8.8, 10.1, 12.5, 13.0, 15.6}
	method := methods.MinimumSquareRegressionMethod[float64]{}

	a, r2, err := method.Run(x, y)
	if err != nil {
		views.ReportError(err)
		return
	}
	views.ReportResult("MinimumSquareDirectly", nil, uint32(1), r2, a[0], a[1])
	performAnalyticsCalculations(a, y)
}

func quest05MinimumSquareWithAdjusts() {
	xi := []float64{10, 20, 30, 40, 50, 60, 70, 80}
	yi := []float64{125, 70, 380, 550, 610, 1220, 830, 1450}

	method := methods.MinimumSquareRegressionMethod[float64]{}

	// Fitting a simple power equation (y = alpha * x^beta)
	logXi := make([]float64, len(xi))
	logYi := make([]float64, len(yi))
	for i := range xi {
		logXi[i] = math.Log(xi[i])
		logYi[i] = math.Log(yi[i])
	}

	a, r2, err := method.Run(logXi, logYi)
	if err != nil {
		views.ReportError(err)
		return
	}

	coefficients := struct{ alpha, beta float64 }{alpha: a[0], beta: a[1]}

	// Creating a function for the fitted power equation
	fittedFunction := func(x float64) float64 {
		return coefficients.alpha * math.Pow(x, coefficients.beta)
	}

	// New adjusted values
	newY := make([]float64, len(xi))
	for index, x := range xi {
		newY[index] = fittedFunction(x)
	}

	a, r2, err = method.Run(xi, newY)
	if err != nil {
		views.ReportError(err)
		return
	}
	views.ReportResult("MinimumSquareWithAdjust", nil, uint32(1), r2, a[0], a[1])
	performAnalyticsCalculations(a, newY)
}

func quest06PolynomialRegression() {
	x := []float64{0, 1, 2, 3, 4, 5}
	y := []float64{0, 20, 60, 68, 77, 100}
	method := methods.MinimumSquarePolynomialRegressionMethod[float64]{}

	for degree := 1; degree <= 5; degree++ {
		result, err := method.Run(x, y, float64(degree))
		if err != nil {
			views.ReportError(err)
			return
		}
		views.ReportResult(
			"PolynomialRegression - "+strconv.FormatInt(int64(degree), 10)+" degree",
			nil, uint32(1), result...,
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

	quest01NewtonInterpolation()
	quest02LagrangeInterpolation()
	quest03VandermondeInterpolation()
	quest04MinimumSquareDirectly()
	quest05MinimumSquareWithAdjusts()
	quest06PolynomialRegression()
}
