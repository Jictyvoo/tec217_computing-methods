package methods

import (
	"errors"

	"github.com/jictyvoo/tec217_computing-methods/internal/models"
)

type GaussAddons uint8

const (
	AddonPivot GaussAddons = 1 << iota
	AddonRecursiveSubstitution
)

type GaussEliminationMethod[T models.Numeric] struct {
	addons GaussAddons
}

func (mtd GaussEliminationMethod[T]) isMatrixSquare(equationsMatrix [][]T) bool {
	totalEquations := len(equationsMatrix)
	for _, equation := range equationsMatrix {
		// Check size by removing the result element
		if len(equation)-1 != totalEquations {
			return false
		}
	}
	return true
}

func (mtd GaussEliminationMethod[T]) subtractEquations(factor T, a, b []T) (result []T) {
	result = make([]T, len(a))
	for index, value := range b {
		result[index] = a[index] - value*factor
	}
	return
}

func (mtd GaussEliminationMethod[T]) determinant(matrix [][]T) (det T) {
	det = 1
	for index := 0; index < len(matrix); index++ {
		det *= matrix[index][index]
	}

	return
}

func (mtd GaussEliminationMethod[T]) Run(inputMatrix [][]T) (det T, foundRoots []T, err error) {
	// Checking if the given matrix is a square one
	if !mtd.isMatrixSquare(inputMatrix) {
		err = errors.New("the given matrix isn't a square one")
		return
	}

	// Create a copy of input matrix
	equationsMatrix := make([][]T, len(inputMatrix))
	foundRoots = make([]T, len(inputMatrix))
	for index, equation := range inputMatrix {
		equationsMatrix[index] = make([]T, len(equation))
		copy(equationsMatrix[index], equation)
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
		}
	}

	// Calculate determinant
	det = mtd.determinant(equationsMatrix)

	// Calculate roots from bottom-up
	divisionIndex := len(equationsMatrix) - 1
	targetEquation := equationsMatrix[divisionIndex]
	// xn ← bn / an,n
	foundRoots[divisionIndex] = targetEquation[len(targetEquation)-1] / targetEquation[divisionIndex]
	for eqIndex := divisionIndex - 1; eqIndex >= 0; eqIndex-- {
		targetEquation = equationsMatrix[eqIndex]
		// Sum ← bi
		sum := targetEquation[len(targetEquation)-1]
		for i := eqIndex + 1; i < len(equationsMatrix); i++ {
			// Sum ← Sum – ai,j * xj
			sum -= targetEquation[i] * foundRoots[i]
		}
		// xi = sum / ai,i
		foundRoots[eqIndex] = sum / targetEquation[eqIndex]
	}

	return
}
