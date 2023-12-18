package main

import (
	"errors"
	"log/slog"
	"os"
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
}
