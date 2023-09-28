package methods

import (
	"errors"

	"github.com/jictyvoo/tec217_computing-methods/internal/models"
	"github.com/jictyvoo/tec217_computing-methods/internal/utils"
)

type LUDecompositionMethod[T models.Numeric] struct {
	matrixHelper[T]
	commonLinearSystemState[T]
	lMatrix [][]T
	uMatrix [][]T
}

func (mtd *LUDecompositionMethod[T]) L() [][]T {
	return mtd.lMatrix
}

func (mtd *LUDecompositionMethod[T]) U() [][]T {
	return mtd.uMatrix
}

func (mtd *LUDecompositionMethod[T]) Run(inputMatrix [][]T) (x, y []T, err error) {
	// Checking if the given matrix is a square
	if !mtd.isMatrixSquare(inputMatrix) {
		err = errors.New("the given matrix isn't a square one")
		return
	}

	// Create the L and U matrices
	matrixSize := len(inputMatrix)
	results := make([]T, matrixSize)
	mtd.lMatrix = mtd.makeIdentity(uint8(matrixSize))
	mtd.uMatrix = utils.DeepCopyBidimensionalSlice(inputMatrix)
	for index, equation := range mtd.uMatrix {
		mtd.uMatrix[index] = equation[:matrixSize]
		results[index] = equation[matrixSize]
	}

	for eqIndex := 0; eqIndex < matrixSize; eqIndex++ {
		pivot := mtd.uMatrix[eqIndex][eqIndex]
		equation := mtd.uMatrix[eqIndex]

		for index := eqIndex + 1; index < matrixSize; index++ {
			nextEquation := mtd.uMatrix[index]
			factor := nextEquation[eqIndex] / pivot
			mtd.uMatrix[index] = mtd.subtractEquations(factor, nextEquation, equation)
			mtd.lMatrix[index][eqIndex] = factor
		}
	}

	y, x = mtd.calculateXY(matrixSize, results)
	return
}

func (mtd *LUDecompositionMethod[T]) calculateXY(matrixSize int, results []T) (y, x []T) {
	// Solve Inferior triangle
	y = make([]T, matrixSize)
	for eqIndex := 0; eqIndex < matrixSize; eqIndex++ {
		targetEquation := mtd.L()[eqIndex]
		sum := results[eqIndex]
		for index := 0; index < matrixSize; index++ {
			if index != eqIndex {
				sum -= targetEquation[index] * y[index]
			}
		}
		y[eqIndex] = sum / targetEquation[eqIndex]
		mtd.registerRootCalculation(y, targetEquation[eqIndex], sum)
	}

	// Solve superior triangle
	x = make([]T, matrixSize)
	for eqIndex := matrixSize - 1; eqIndex >= 0; eqIndex-- {
		targetEquation := mtd.U()[eqIndex]
		sum := y[eqIndex]
		for index := eqIndex + 1; index < matrixSize; index++ {
			sum -= targetEquation[index] * x[index]
		}
		x[eqIndex] = sum / targetEquation[eqIndex]
		mtd.registerRootCalculation(x, targetEquation[eqIndex], sum)
	}

	return
}
