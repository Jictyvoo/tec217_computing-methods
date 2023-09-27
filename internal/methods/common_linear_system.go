package methods

import (
	"math"

	"github.com/jictyvoo/tec217_computing-methods/internal/models"
	"github.com/jictyvoo/tec217_computing-methods/internal/utils"
)

type matrixHelper[T models.Numeric] struct{}

func (mtd matrixHelper[T]) pivot(matrix [][]T, currentIndex, lookupColumn int) (pivotElement T) {
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

	// Swap the current row and the maxValue row if necessary
	if swapIndex != currentIndex {
		matrix[swapIndex], matrix[currentIndex] = matrix[currentIndex], matrix[swapIndex]
	}

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

type commonLinearSystemState[T models.Numeric] struct {
	transformationSteps  []models.MatrixTransformationStep[T]
	rootCalculationSteps []models.RootCalculationStep[T]
	Determinant          T
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
