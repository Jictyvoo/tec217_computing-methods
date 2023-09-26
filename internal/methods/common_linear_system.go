package methods

import (
	"github.com/jictyvoo/tec217_computing-methods/internal/models"
	"github.com/jictyvoo/tec217_computing-methods/internal/utils"
)

type (
	matrixHelper[T models.Numeric]            struct{}
	commonLinearSystemState[T models.Numeric] struct {
		triangulationSteps   []models.TriangulationStep[T]
		rootCalculationSteps []models.RootCalculationStep[T]
		Determinant          T
		Roots                []T
		matrixHelper[T]
	}
)

func (mtd matrixHelper[T]) determinant(matrix [][]T) (det T) {
	det = 1
	for index := 0; index < len(matrix); index++ {
		det *= matrix[index][index]
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

func (mtd matrixHelper[T]) subtractEquations(factor T, a, b []T) (result []T) {
	result = make([]T, len(a))
	for index, value := range b {
		result[index] = a[index] - value*factor
	}
	return
}

//goland:noinspection GoMixedReceiverTypes

func (mtd *commonLinearSystemState[T]) registerTriangulation(
	inputMatrix [][]T, multiplier T, subtractionRow, destinationRow uint8,
) {
	mtd.triangulationSteps = append(mtd.triangulationSteps, models.TriangulationStep[T]{
		Matrix:         utils.DeepCopyBidimensionalSlice(inputMatrix),
		Multiplier:     multiplier,
		SubtractionRow: subtractionRow,
		SubtractedRow:  destinationRow,
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

func (mtd commonLinearSystemState[T]) InteractionData() ([]models.TriangulationStep[T], []models.RootCalculationStep[T]) {
	return mtd.triangulationSteps, mtd.rootCalculationSteps
}

//goland:noinspection GoMixedReceiverTypes

func (mtd *commonLinearSystemState[T]) Reset() {
	mtd.triangulationSteps = mtd.triangulationSteps[:0]
	mtd.rootCalculationSteps = mtd.rootCalculationSteps[:0]
}
