package methods

import (
	"fmt"
	"math"

	"github.com/jictyvoo/tec217_computing-methods/internal/models"
	"github.com/jictyvoo/tec217_computing-methods/internal/utils"
)

type JacobiMethod[T models.Numeric] struct {
	xError []T

	matrixHelper[T]
	rootCalculationSteps commonFuncZeroState[T]
	matrixSetup          trackedMatrixOperations[T]
}

func (mtd *JacobiMethod[T]) innerProduct(a []T, b []T) T {
	result := T(0)
	for i := range a {
		result += a[i] * b[i]
	}
	return result
}

func (mtd *JacobiMethod[T]) D() []T {
	return mtd.matrixSetup.Roots
}

func (mtd *JacobiMethod[T]) CalculatedErrors() []T {
	return mtd.xError
}

func (mtd *JacobiMethod[T]) InteractionData() (
	[]models.MatrixTransformationStep[T], []models.RootCalculationStep[T], []models.InteractionData[T],
) {
	return mtd.matrixSetup.transformationSteps, mtd.matrixSetup.rootCalculationSteps, mtd.rootCalculationSteps.interactions
}

func (mtd *JacobiMethod[T]) Run(
	inputMatrix [][]T,
	initialGuess T, maxIterations uint32, epsilon T,
) (foundRoots []T, currentIteration uint32, err error) {
	if err = mtd.executeCriteria(inputMatrix, CriteriaDiagonalDominant); err != nil {
		return
	}

	matrixSize := len(inputMatrix)

	results := make([]T, matrixSize)
	workingMatrix := utils.DeepCopyBidimensionalSlice(inputMatrix)
	for index, equation := range workingMatrix {
		workingMatrix[index] = equation[:matrixSize]
		results[index] = equation[matrixSize]
	}

	C := mtd.newEmptyMatrix(uint8(matrixSize))
	d := make([]T, matrixSize)

	// Finished pre-setup, now runs the method
	mtd.matrixSetup.mountCAndD(results, workingMatrix, C, d)

	foundRoots = make([]T, matrixSize)
	mtd.matrixSetup.Roots = d
	xPrev := make([]T, matrixSize)
	mtd.xError = make([]T, matrixSize)

	foundRoots[0] = initialGuess
	for currentIteration < maxIterations {
		currentIteration += 1
		success, maxError := mtd.calculateRoot(C, d, matrixSize, epsilon, foundRoots, xPrev)
		mtd.rootCalculationSteps.registerInteraction(
			xPrev, currentIteration, maxError, 0, foundRoots...,
		)
		if success {
			return
		}
	}

	err = fmt.Errorf("failed to converge in %d iterations", currentIteration)
	return
}

func (mtd *JacobiMethod[T]) calculateRoot(
	C [][]T, d []T,
	matrixSize int, epsilon T,
	foundRoots, xPrev []T,
) (success bool, maximumXError T) {
	// Save previous root results to being able to perform the error calculation correctly
	copy(xPrev, foundRoots)
	for index := 0; index < matrixSize; index++ {
		foundRoots[index] = mtd.innerProduct(C[index], xPrev) + d[index]
	}

	maximumXError = T(0)
	for index := 0; index < matrixSize; index++ {
		mtd.xError[index] = T(
			math.Abs(float64((foundRoots[index]-xPrev[index])/foundRoots[index])) * 100,
		)
		if mtd.xError[index] > maximumXError {
			maximumXError = mtd.xError[index]
		}
	}

	success = maximumXError < epsilon
	return
}
