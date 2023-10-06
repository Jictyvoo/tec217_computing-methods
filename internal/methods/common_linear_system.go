package methods

import (
	"math"

	"github.com/jictyvoo/tec217_computing-methods/internal/models"
	"github.com/jictyvoo/tec217_computing-methods/internal/utils"
)

type matrixHelper[T models.Numeric] struct{}

func (mtd matrixHelper[T]) swapMatrixLines(matrix [][]T, indexes [2]uint32) {
	// Swap the current row and the maxValue row if necessary
	if indexes[0] != indexes[1] {
		matrix[indexes[0]], matrix[indexes[1]] = matrix[indexes[1]], matrix[indexes[0]]
	}

	return
}

func (mtd matrixHelper[T]) pivot(
	matrix [][]T, currentIndex, lookupColumn int,
) (pivotElement T, swapIndexes [2]uint32) {
	// Find the index of the maxValue value in column
	pivotElement = matrix[currentIndex][currentIndex]
	maxValue := math.Abs(float64(pivotElement))
	swapIndex := currentIndex
	for index := currentIndex + 1; index < len(matrix); index++ {
		element := matrix[index][lookupColumn]
		if abs := math.Abs(float64(element)); abs > maxValue {
			swapIndex = index
			maxValue = abs
			pivotElement = element
		}
	}

	swapIndexes = [2]uint32{uint32(swapIndex), uint32(currentIndex)}
	return
}

func (mtd matrixHelper[T]) determinant(matrix [][]T) (det T) {
	det = 1
	for index := 0; index < len(matrix); index++ {
		det *= matrix[index][index]
	}

	return
}

func (mtd matrixHelper[T]) makeIdentity(size uint8) (identity [][]T) {
	identity = make([][]T, size)
	for index := uint8(0); index < size; index++ {
		identity[index] = make([]T, size)
		identity[index][index] = 1
	}

	return
}

func (mtd matrixHelper[T]) newEmptyMatrix(size uint8) (identity [][]T) {
	identity = make([][]T, size)
	for index := uint8(0); index < size; index++ {
		identity[index] = make([]T, size)
	}

	return
}

func (mtd matrixHelper[T]) isMatrixSquare(equationsMatrix [][]T) bool {
	totalEquations := len(equationsMatrix)
	for _, equation := range equationsMatrix {
		// Check size by removing the result element
		if len(equation)-1 != totalEquations {
			return false
		}
	}
	return true
}

func (mtd matrixHelper[T]) execOnEquation(a []T, callback func(T) T) (result []T) {
	result = make([]T, len(a))
	for index := 0; index < len(a); index += 1 {
		result[index] = callback(a[index])
	}
	return
}

func (mtd matrixHelper[T]) subtractEquations(factor T, a, b []T) (result []T) {
	result = make([]T, len(a))
	for index, value := range b {
		result[index] = a[index] - value*factor
	}
	return
}

func (mtd matrixHelper[T]) isDiagonallyDominant(A [][]T) bool {
	// Get the dimensions of the matrix
	matrixSize := len(A)

	for index := 0; index < matrixSize; index++ {
		sum := T(0)
		for subIndex := 0; subIndex < matrixSize; subIndex++ {
			if index != subIndex {
				// Add the absolute value to the accumulator
				absoluteAij := A[index][subIndex]
				if A[index][subIndex] < 0 {
					absoluteAij = -A[index][subIndex]
				}
				sum += absoluteAij
			}
		}
		// Check if the value on the main diagonal is less than sum
		if A[index][index] < sum {
			return false
		}
	}

	// The matrix is diagonally dominant
	return true
}

type commonLinearSystemState[T models.Numeric] struct {
	transformationSteps  []models.MatrixTransformationStep[T]
	rootCalculationSteps []models.RootCalculationStep[T]
	Roots                []T
}

//goland:noinspection GoMixedReceiverTypes

func (mtd *commonLinearSystemState[T]) registerMatrixTransformation(
	inputMatrix [][]T, multiplier T,
	subtractionRow, destinationRow uint8,
	operationOpt ...models.AlgebraicOperation,
) {
	operation := models.OpSub
	if len(operationOpt) > 0 {
		operation = operationOpt[0]
	}
	mtd.transformationSteps = append(mtd.transformationSteps, models.MatrixTransformationStep[T]{
		Matrix:     utils.DeepCopyBidimensionalSlice(inputMatrix),
		Multiplier: multiplier,
		Operation:  operation,
		RightRow:   subtractionRow,
		LeftRow:    destinationRow,
	})
}

//goland:noinspection GoMixedReceiverTypes

func (mtd *commonLinearSystemState[T]) registerRootCalculation(
	calculatedRoots []T, dividend, divisorSum T,
) {
	mtd.rootCalculationSteps = append(mtd.rootCalculationSteps, models.RootCalculationStep[T]{
		Roots:       utils.CloneSlice(calculatedRoots),
		DividendSum: divisorSum,
		Divisor:     dividend,
	})
}

//goland:noinspection GoMixedReceiverTypes

func (mtd commonLinearSystemState[T]) InteractionData() ([]models.MatrixTransformationStep[T], []models.RootCalculationStep[T]) {
	return mtd.transformationSteps, mtd.rootCalculationSteps
}

//goland:noinspection GoMixedReceiverTypes

func (mtd *commonLinearSystemState[T]) Reset() {
	mtd.transformationSteps = mtd.transformationSteps[:0]
	mtd.rootCalculationSteps = mtd.rootCalculationSteps[:0]
}

type trackedMatrixOperations[T models.Numeric] struct {
	commonLinearSystemState[T]
}

// forwardSubstitution Solve Inferior triangle
func (mtd *trackedMatrixOperations[T]) forwardSubstitution(
	inputMatrix [][]T, matrixSize int, results []T,
) (y []T) {
	y = make([]T, matrixSize)
	for eqIndex := 0; eqIndex < matrixSize; eqIndex++ {
		targetEquation := inputMatrix[eqIndex]
		sum := results[eqIndex]
		for index := 0; index < matrixSize; index++ {
			if index != eqIndex {
				sum -= targetEquation[index] * y[index]
			}
		}
		y[eqIndex] = sum / targetEquation[eqIndex]
		mtd.registerRootCalculation(y, targetEquation[eqIndex], sum)
	}

	mtd.Roots = y
	return
}

// backwardSubstitution Solve superior triangle
func (mtd *trackedMatrixOperations[T]) backwardSubstitution(
	inputMatrix [][]T, matrixSize int, y []T,
) (x []T) {
	x = make([]T, matrixSize)
	for eqIndex := matrixSize - 1; eqIndex >= 0; eqIndex-- {
		targetEquation := inputMatrix[eqIndex]
		sum := y[eqIndex]
		for index := eqIndex + 1; index < matrixSize; index++ {
			sum -= targetEquation[index] * x[index]
		}
		x[eqIndex] = sum / targetEquation[eqIndex]
		mtd.registerRootCalculation(x, targetEquation[eqIndex], sum)
	}

	mtd.Roots = x
	return
}

func (mtd *trackedMatrixOperations[T]) mountCAndD(
	results []T, workingMatrix, C [][]T, d []T,
) {
	matrixSize := len(C)

	for eqIndex := 0; eqIndex < matrixSize; eqIndex++ {
		for index := 0; index < matrixSize; index++ {
			var setValue T = 0
			if eqIndex == index {
				divisor, dividend := results[eqIndex], workingMatrix[eqIndex][eqIndex]
				d[eqIndex] = divisor / dividend
				mtd.registerRootCalculation(d, dividend, divisor)
			} else {
				setValue = -workingMatrix[eqIndex][index] / workingMatrix[eqIndex][eqIndex]
			}

			C[eqIndex][index] = setValue
			mtd.registerMatrixTransformation(
				C, setValue, uint8(eqIndex), uint8(index), models.OpAttribution,
			)
		}
	}

	return
}
