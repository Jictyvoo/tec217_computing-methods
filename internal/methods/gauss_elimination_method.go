package methods

import (
	"errors"

	"github.com/jictyvoo/tec217_computing-methods/internal/models"
	"github.com/jictyvoo/tec217_computing-methods/internal/utils"
)

type GaussAddons uint8

const (
	AddonPivot GaussAddons = 1 << iota
)

type GaussEliminationMethod[T models.Numeric] struct {
	Addons      GaussAddons
	Determinant T

	matrixHelper[T]
	commonLinearSystemState[T]
}

func (mtd *GaussEliminationMethod[T]) Run(inputMatrix [][]T) (det T, foundRoots []T, err error) {
	// Checking if the given matrix is a square
	if !mtd.isMatrixSquare(inputMatrix) {
		err = errors.New("the given matrix isn't a square one")
		return
	}

	// Create a copy of input matrix
	equationsMatrix := utils.DeepCopyBidimensionalSlice(inputMatrix)

	if mtd.Addons&AddonPivot == AddonPivot {
		_, swapIndexes := mtd.pivot(equationsMatrix, 0, 0)
		mtd.swapMatrixLines(equationsMatrix, swapIndexes)
	}

	for eqIndex, equation := range equationsMatrix {
		if eqIndex == len(equationsMatrix)-1 {
			break
		}

		// Eliminate the next equations
		if eqIndex+1 >= len(equation) {
			continue
		}
		nextEquation := equationsMatrix[eqIndex+1]
		if nextEquation[eqIndex] == 0 {
			continue
		}

		for index := eqIndex + 1; index < len(equationsMatrix); index++ {
			nextEquation = equationsMatrix[index]
			divisionFactor := nextEquation[eqIndex] / equation[eqIndex]
			equationsMatrix[index] = mtd.subtractEquations(divisionFactor, nextEquation, equation)
			mtd.registerMatrixTransformation(
				equationsMatrix, divisionFactor,
				uint8(eqIndex), uint8(index),
			)
		}
	}

	// Calculate determinant
	det = mtd.determinant(equationsMatrix)

	// Calculate roots from bottom-up
	foundRoots = make([]T, len(inputMatrix))
	for eqIndex := len(equationsMatrix) - 1; eqIndex >= 0; eqIndex-- {
		targetEquation := equationsMatrix[eqIndex]
		// Sum ← b_{i}
		sum := targetEquation[len(targetEquation)-1]
		for i := eqIndex + 1; i < len(equationsMatrix); i++ {
			// Sum ← Sum – a_{i,j} * x_{j}
			sum -= targetEquation[i] * foundRoots[i]
		}
		// x_{i} = sum / a_{i,i}
		foundRoots[eqIndex] = sum / targetEquation[eqIndex]
		mtd.registerRootCalculation(foundRoots, targetEquation[eqIndex], sum)
	}

	mtd.Determinant = det
	mtd.Roots = foundRoots
	return
}
