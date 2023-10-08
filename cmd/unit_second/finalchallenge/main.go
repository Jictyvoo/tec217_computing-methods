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

func quest01GaussElimination() {
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

func quest02ThreeTanks() {
	var (
		buffer, inverseMatrix strings.Builder
		method                = methods.GaussJordanMethod[float64]{}
	)

	const (
		Q33 = 120
		Q13 = 40
		Q12 = 90
		Q23 = 60
		Q21 = 30
	)

	det, foundRoots, err := method.Run(
		[][]float64{
			// 200 + Q21*c2 = Q13*c1 + Q12*c1
			{Q13 + Q12, -Q21, 0, 200},
			// Q12*c1 = Q21*c2 + Q23*c2
			{Q12, -Q21 - Q23, 0, 0},
			// Q13*c1 + Q23*c2 + 500 = Q33*c3
			{Q13, Q23, -Q33, -500},
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
	views.ReportLinearSystemResult("GaussJordan - ThreeTanks", det, foundRoots, &buffer)
	views.ReportMatrixTable("Inverse", &inverseMatrix)
}

func quest03GaussEliminationDumb() {
	var (
		buffer strings.Builder
		method = methods.GaussEliminationMethod[float64]{}
	)

	det, foundRoots, err := method.Run(
		[][]float64{
			{0.52, 0.2, 0.25, 4800},
			{0.3, 0.5, 0.2, 5800},
			{0.18, 0.3, 0.55, 5700},
		},
	)

	triangulationSteps, rootCalculationSteps := method.InteractionData()
	if err = errors.Join(err, utils.WriteLinearInteractionAsCSV(&buffer, triangulationSteps, rootCalculationSteps)); err != nil {
		views.ReportError(err)
		return
	}
	views.ReportLinearSystemResult("GaussElimination", det, foundRoots, &buffer)
}

func quest04LUDecomposition() {
	var (
		buffer, lBuffer, uBuffer strings.Builder
		steps                    struct {
			L, U struct {
				triangulation   []models.MatrixTransformationStep[float64]
				rootCalculation []models.RootCalculationStep[float64]
			}
		}
		method = methods.LUDecompositionMethod[float64]{}
	)

	det, foundRoots, err := method.Run(
		[][]float64{
			{7, 2, -3, -12},
			{2, 5, -3, -20},
			{1, -1, -6, -26},
		},
	)

	// Original b
	{
		steps.L.triangulation, steps.L.rootCalculation = method.LInteractionData()
		steps.U.triangulation, steps.U.rootCalculation = method.UInteractionData()
		err = errors.Join(
			err,
			utils.WriteLinearInteractionAsCSV(
				&buffer,
				steps.L.triangulation,
				steps.L.rootCalculation,
			),
			utils.WriteLinearInteractionAsCSV(
				&buffer,
				steps.U.triangulation,
				steps.U.rootCalculation,
			),
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

	// Fake b2
	{
		method = methods.LUDecompositionMethod[float64]{}
		det, foundRoots, err = method.Run(
			[][]float64{
				{7, 2, -3, 12},
				{2, 5, -3, 18},
				{1, -1, -6, -6},
			},
		)

		steps.L.triangulation, steps.L.rootCalculation = method.LInteractionData()
		steps.U.triangulation, steps.U.rootCalculation = method.UInteractionData()

		buffer, lBuffer, uBuffer = strings.Builder{}, strings.Builder{}, strings.Builder{}
		err = errors.Join(
			err,
			utils.WriteLinearInteractionAsCSV(
				&buffer,
				steps.L.triangulation,
				steps.L.rootCalculation,
			),
			utils.WriteLinearInteractionAsCSV(
				&buffer,
				steps.U.triangulation,
				steps.U.rootCalculation,
			),
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
}

func quest05JordanWithoutPivot() {
	var (
		buffer, inverseMatrix strings.Builder
		method                = methods.GaussJordanMethod[float64]{}
		A                     = [][]float64{
			{10, 2, -1, 27},
			{-3, -6, 2, -61.5},
			{1, 1, 5, -21.5},
		}
	)

	det, foundRoots, err := method.Run(A)

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
	views.ReportLinearSystemResult("GaussJordan - JordanWithoutPivot", det, foundRoots, &buffer)
	views.ReportMatrixTable("Inverse", &inverseMatrix)

	multiplyResult := multiplyMatrix(A, method.Inverse(), 3)
	var identityMatrix strings.Builder
	_ = views.MatrixTableCSV(multiplyResult, &identityMatrix)
	views.ReportMatrixTable("Identity", &identityMatrix)
}

func quest08GaussSeidel() {
	var (
		matrixBuffer, rootBuffer strings.Builder
		method                   = methods.GaussSeidelMethod[float64]{
			Criteria: methods.CriteriaSassenfeld,
		}
	)

	foundRoots, totalIterations, err := method.Run(
		[][]float64{
			{0.8, -0.4, 0, 41},
			{-0.4, 0.8, -0.4, 25},
			{0, -0.4, 0.8, 105},
		}, []float64{0, 0, 0}, 1000, 0.05,
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
	views.ReportResult("GaussSeidel", &rootBuffer, totalIterations, foundRoots...)
	views.ReportLinearSystemResult("", 0, method.D(), &matrixBuffer)
}

func quest09Jacobi() {
	var (
		matrixBuffer, rootBuffer strings.Builder
		method                   = methods.JacobiMethod[float64]{}
	)

	foundRoots, totalIterations, err := method.Run(
		[][]float64{
			{10, 2, -1, 27},
			{-3, -6, 2, -61.5},
			{1, 1, 5, -21.5},
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

func quest10GaussSeidelSOR() {
	var (
		matrixBuffer, rootBuffer strings.Builder
		method                   = methods.GaussSeidelMethod[float64]{RelaxationFactor: 1.2}
	)

	foundRoots, totalIterations, err := method.Run(
		[][]float64{
			{-8, 1, -2, -20},
			{2, -6, -1, -38},
			{-3, -1, 7, -34},
		}, []float64{0, 0, 0}, 1000, 0.05,
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
	views.ReportResult("GaussSeidel-SOR", &rootBuffer, totalIterations, foundRoots...)
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

	quest01GaussElimination()
	quest02ThreeTanks()
	quest03GaussEliminationDumb()
	quest04LUDecomposition()
	quest05JordanWithoutPivot()
	quest08GaussSeidel()
	quest09Jacobi()
	quest10GaussSeidelSOR()
}
