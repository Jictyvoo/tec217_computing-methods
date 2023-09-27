package methods

import (
	"errors"

	"github.com/jictyvoo/tec217_computing-methods/internal/models"
	"github.com/jictyvoo/tec217_computing-methods/internal/utils"
)

type GaussJordanMethod[T models.Numeric] struct {
	matrixHelper[T]
	commonLinearSystemState[T]
}

func (mtd *GaussJordanMethod[T]) Run(inputMatrix [][]T) (det T, foundRoots []T, err error) {
	// Checking if the given matrix is a square
	if !mtd.isMatrixSquare(inputMatrix) {
		err = errors.New("the given matrix isn't a square one")
		return
	}

	// Create a copy of input matrix
	equationsMatrix := utils.DeepCopyBidimensionalSlice(inputMatrix)

	for eqIndex := 0; eqIndex < len(equationsMatrix); eqIndex += 1 {
		mtd.pivot(equationsMatrix, eqIndex, eqIndex)
		{
			equation := equationsMatrix[eqIndex]
			pivotElement := equation[eqIndex]
			equationsMatrix[eqIndex] = mtd.execOnEquation(
				equation, func(value T) T {
					return value / pivotElement
				},
			)

			// Eliminate the next equations
			if eqIndex+1 >= len(equation) {
				continue
			}
		}

		for index := 0; index < len(equationsMatrix); index++ {
			if index != eqIndex {
				equation := equationsMatrix[index]
				factor := equation[eqIndex]
				equationsMatrix[index] = mtd.subtractEquations(
					factor, equation,
					equationsMatrix[eqIndex],
				)
				mtd.registerTriangulation(equationsMatrix, -factor, uint8(eqIndex), uint8(index))
			}
		}
	}

	foundRoots = make([]T, len(equationsMatrix))
	for i, equation := range equationsMatrix {
		foundRoots[i] = equation[len(equation)-1]
	}
	return
}
