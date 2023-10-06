package main

import (
	"errors"
	"log/slog"
	"os"
	"strings"

	"github.com/jictyvoo/tec217_computing-methods/internal/methods"
	"github.com/jictyvoo/tec217_computing-methods/internal/models"
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

func quest04LUDecomposition() {
	var (
		buffer, lBuffer, uBuffer strings.Builder
		method                   = methods.LUDecompositionMethod[float64]{}
	)

	det, foundRoots, err := method.Run(
		[][]float64{
			{2, 1, -1, 3},
			{-1, 3, 2, 1},
			{3, 1, -3, 2},
			/*{2, 3, -4, 4, -1},
			{-4, -7, 11, -6, 5},
			{6, 11, -20, 10, -13},
			{-2, -7, 22, -6, 25},*/
		},
	)

	var steps struct {
		L, U struct {
			triangulation   []models.MatrixTransformationStep[float64]
			rootCalculation []models.RootCalculationStep[float64]
		}
	}
	steps.L.triangulation, steps.L.rootCalculation = method.LInteractionData()
	steps.U.triangulation, steps.U.rootCalculation = method.UInteractionData()
	err = errors.Join(
		err,
		utils.WriteLinearInteractionAsCSV(&buffer, steps.L.triangulation, steps.L.rootCalculation),
		utils.WriteLinearInteractionAsCSV(&buffer, steps.U.triangulation, steps.U.rootCalculation),
		views.MatrixTableCSV(method.L(), &lBuffer),
		views.MatrixTableCSV(method.U(), &uBuffer),
	)

	if err != nil {
		views.ReportError(err)
		return
	}

	views.ReportLinearSystemResult("LU-Decomposition", det, foundRoots, &buffer)
	views.ReportMatrixTable("L", &lBuffer)
	views.ReportMatrixTable("U", &uBuffer)
}

func quest05Jacobi() {
	var (
		matrixBuffer, rootBuffer strings.Builder
		method                   = methods.JacobiMethod[float64]{}
	)

	foundRoots, totalIterations, err := method.Run(
		[][]float64{
			{5, 1, 1, 5},
			{3, 4, 1, 6},
			{3, 3, 6, 0},
		}, 0, 1000, 0.05,
	)

	triangulationSteps, dCalculationSteps, rootCalculationSteps := method.InteractionData()
	err = errors.Join(
		err,
		utils.WriteInteractionAsCSV(&rootBuffer, rootCalculationSteps),
		utils.WriteLinearInteractionAsCSV(&matrixBuffer, triangulationSteps, dCalculationSteps),
	)
	if err != nil {
		views.ReportError(err)
		return
	}
	views.ReportResult("Jacobi", &rootBuffer, totalIterations, foundRoots...)
	views.ReportLinearSystemResult("", 0, method.D(), &matrixBuffer)
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
	quest04LUDecomposition()
	quest05Jacobi()
}
