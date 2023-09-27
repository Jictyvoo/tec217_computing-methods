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

func quest01Gauss() {
	var (
		buffer strings.Builder
		method = methods.GaussEliminationMethod[float64]{}
	)

	det, foundRoots, err := method.Run(
		[][]float64{
			{10, 2, -1, 27},
			{-3, -5, 2, -61.5},
			{1, 1, 6, -21.5},
		},
	)

	triangulationSteps, rootCalculationSteps := method.InteractionData()
	if err = errors.Join(err, utils.WriteLinearInteractionAsCSV(&buffer, triangulationSteps, rootCalculationSteps)); err != nil {
		views.ReportError(err)
		return
	}
	views.ReportLinearSystemResult("GaussElimination", det, foundRoots, &buffer)
}

func quest02GaussPivot() {
	var (
		buffer strings.Builder
		method = methods.GaussEliminationMethod[float64]{Addons: methods.AddonPivot}
	)

	det, foundRoots, err := method.Run(
		[][]float64{
			{2, -6, -1, -38},
			{-3, -1, 7, -34},
			{-8, 1, -2, -20},
		},
	)

	triangulationSteps, rootCalculationSteps := method.InteractionData()
	if err = errors.Join(err, utils.WriteLinearInteractionAsCSV(&buffer, triangulationSteps, rootCalculationSteps)); err != nil {
		views.ReportError(err)
		return
	}
	views.ReportLinearSystemResult("GaussElimination", det, foundRoots, &buffer)
}

func quest03GaussJordan() {
	var (
		buffer, inverseMatrix strings.Builder
		method                = methods.GaussJordanMethod[float64]{}
	)

	det, foundRoots, err := method.Run(
		[][]float64{
			{-0.04, 0.04, 0.12, 3},
			{0.56, -1.56, 0.32, 1},
			{-0.24, 1.24, -0.28, 0},
		},
	)

	triangulationSteps, rootCalculationSteps := method.InteractionData()
	err = errors.Join(
		err,
		utils.WriteLinearInteractionAsCSV(&buffer, triangulationSteps, rootCalculationSteps),
		views.MatrixTableCSV(method.Inverse(), &inverseMatrix),
	)

	if err != nil {
		views.ReportError(err)
		return
	}
	views.ReportLinearSystemResult("GaussJordan", det, foundRoots, &buffer)
	views.ReportMatrixTable("Inverse", &inverseMatrix)
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
	quest01Gauss()
	quest02GaussPivot()
	quest03GaussJordan()
}
