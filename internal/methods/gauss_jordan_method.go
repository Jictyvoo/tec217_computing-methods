package methods

import (
	"errors"
	"slices"

	"github.com/jictyvoo/tec217_computing-methods/internal/models"
	"github.com/jictyvoo/tec217_computing-methods/internal/utils"
)

type GaussJordanMethod[T models.Numeric] struct {
	Determinant   T
	inverseMatrix [][]T

	matrixHelper[T]
	commonLinearSystemState[T]
}

func (mtd *GaussJordanMethod[T]) Inverse() [][]T {
	return mtd.inverseMatrix
}

func (mtd *GaussJordanMethod[T]) Run(inputMatrix [][]T) (det T, foundRoots []T, err error) {
	// Checking if the given matrix is a square
	if !mtd.isMatrixSquare(inputMatrix) {
		err = errors.New("the given matrix isn't a square one")
		return
	}

	// Create a copy of input matrix
	equationsMatrix := utils.DeepCopyBidimensionalSlice(inputMatrix)
	matrixSize := len(equationsMatrix)

	// Insert identity matrix between the result and original matrix
	identity := mtd.makeIdentity(uint8(matrixSize))
	for index, idEq := range identity {
		equationsMatrix[index] = slices.Insert(equationsMatrix[index], matrixSize, idEq...)
	}

	det = 1
	for eqIndex := 0; eqIndex < matrixSize; eqIndex += 1 {
		pivotElement, swapIndexes := mtd.pivot(equationsMatrix, eqIndex, eqIndex)
		{
			// Register step before swapping the lines
			mtd.registerMatrixTransformation(
				equationsMatrix, 0, uint8(swapIndexes[0]), uint8(swapIndexes[1]), models.OpPermut,
			)
			mtd.swapMatrixLines(equationsMatrix, swapIndexes)
		}
		// keep updating determinant
		det *= pivotElement
		{
			equation := equationsMatrix[eqIndex]
			equationsMatrix[eqIndex] = mtd.execOnEquation(
				equation, func(value T) T {
					return value / pivotElement
				},
			)
			mtd.registerMatrixTransformation(
				equationsMatrix, pivotElement, uint8(eqIndex), uint8(eqIndex), models.OpDiv,
			)

			// Eliminate the next equations
			if eqIndex+1 >= len(equation) {
				continue
			}
		}

		for index := 0; index < matrixSize; index++ {
			if index != eqIndex {
				equation := equationsMatrix[index]
				factor := equation[eqIndex]
				equationsMatrix[index] = mtd.subtractEquations(
					factor, equation,
					equationsMatrix[eqIndex],
				)
				mtd.registerMatrixTransformation(
					equationsMatrix, factor,
					uint8(eqIndex), uint8(index),
				)
			}
		}
	}

	foundRoots = make([]T, matrixSize)
	for index, equation := range equationsMatrix {
		foundRoots[index] = equation[len(equation)-1]
	}

	mtd.inverseMatrix = make([][]T, matrixSize)
	for index, equation := range equationsMatrix {
		mtd.inverseMatrix[index] = equation[matrixSize : matrixSize*2]
	}

	mtd.Determinant = det
	mtd.Roots = foundRoots
	return
}
