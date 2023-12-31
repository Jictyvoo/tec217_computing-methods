package methods

import (
	"errors"
	"fmt"
	"math"
	"slices"

	"github.com/jictyvoo/tec217_computing-methods/internal/models"
	"github.com/jictyvoo/tec217_computing-methods/internal/utils"
)

type GaussSeidelMethod[T models.Numeric] struct {
	xError           []T
	RelaxationFactor T
	Criteria         MatrixCheckCriteria

	matrixHelper[T]
	matrixSetup          trackedMatrixOperations[T]
	rootCalculationSteps commonFuncZeroState[T]
}

func (mtd *GaussSeidelMethod[T]) InteractionData() (
	[]models.MatrixTransformationStep[T], []models.RootCalculationStep[T], []models.InteractionData[T],
) {
	return mtd.matrixSetup.transformationSteps, mtd.matrixSetup.rootCalculationSteps, mtd.rootCalculationSteps.interactions
}

func (mtd *GaussSeidelMethod[T]) D() []T {
	return mtd.matrixSetup.Roots
}

func (mtd *GaussSeidelMethod[T]) Run(
	inputMatrix [][]T, X0 []T, maxIterations uint32, epsilon T,
) (foundRoots []T, iteration uint32, err error) {
	// Checking if the given matrix is a square
	if !mtd.isMatrixSquare(inputMatrix) {
		err = errors.New("the given matrix isn't a square one")
		return
	}

	matrixSize := len(inputMatrix)
	if err = mtd.executeCriteria(inputMatrix, mtd.Criteria); err != nil {
		return
	}

	results := make([]T, matrixSize)
	workingMatrix := utils.DeepCopyBidimensionalSlice(inputMatrix)
	for index, equation := range workingMatrix {
		workingMatrix[index] = equation[:matrixSize]
		results[index] = equation[matrixSize]
	}

	C := mtd.newEmptyMatrix(uint8(matrixSize))
	d := make([]T, matrixSize)
	mtd.matrixSetup.mountCAndD(results, workingMatrix, C, d)

	foundRoots = X0
	mtd.xError = make([]T, matrixSize)
	for iteration < maxIterations {
		iteration += 1
		oldSolution := slices.Clone(foundRoots)

		foundRoots = mtd.calculateRoots(matrixSize, workingMatrix, foundRoots, oldSolution, results)

		var maximumXError T = 0
		for index := 0; index < matrixSize; index++ {
			mtd.xError[index] = T(math.Abs(
				float64((foundRoots[index]-oldSolution[index])/foundRoots[index]),
			) * 100)

			if mtd.xError[index] > maximumXError {
				maximumXError = mtd.xError[index]
			}
		}
		mtd.rootCalculationSteps.registerInteraction(
			oldSolution, iteration, maximumXError, 0, foundRoots...,
		)
		if maximumXError < epsilon {
			return
		}
	}

	err = fmt.Errorf("failed to converge in %d iterations", iteration)
	return
}

func (mtd *GaussSeidelMethod[T]) calculateRoots(
	matrixSize int, A [][]T, foundRoots []T, oldRoots []T, results []T,
) []T {
	for index := 0; index < matrixSize; index++ {
		var (
			summation      [2]T
			matrixEquation = A[index]
		)

		// First summation
		for j := 0; j < index; j++ {
			summation[0] -= matrixEquation[j] * foundRoots[j]
		}
		// Second summation
		for j := index + 1; j < matrixSize; j++ {
			summation[1] -= matrixEquation[j] * oldRoots[j]
		}

		commonDivisor := matrixEquation[index]
		// Assigning the new value to foundRoots
		tempRoot := ((summation[0] + summation[1]) / commonDivisor) + (results[index] / commonDivisor)
		if mtd.RelaxationFactor > 0 {
			tempRoot = ((1 - mtd.RelaxationFactor) * oldRoots[index]) + (mtd.RelaxationFactor * tempRoot)
		}

		foundRoots[index] = tempRoot
	}
	return foundRoots
}
