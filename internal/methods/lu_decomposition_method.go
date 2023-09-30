package methods

import (
	"errors"

	"github.com/jictyvoo/tec217_computing-methods/internal/models"
	"github.com/jictyvoo/tec217_computing-methods/internal/utils"
)

type LUDecompositionMethod[T models.Numeric] struct {
	Determinant T
	lMatrix     [][]T
	uMatrix     [][]T

	matrixHelper[T]
	lTracker, uTracker trackedMatrixOperations[T]
}

func (mtd *LUDecompositionMethod[T]) L() [][]T {
	return mtd.lMatrix
}

func (mtd *LUDecompositionMethod[T]) U() [][]T {
	return mtd.uMatrix
}

func (mtd *LUDecompositionMethod[T]) LInteractionData() ([]models.MatrixTransformationStep[T], []models.RootCalculationStep[T]) {
	return mtd.lTracker.InteractionData()
}

func (mtd *LUDecompositionMethod[T]) UInteractionData() ([]models.MatrixTransformationStep[T], []models.RootCalculationStep[T]) {
	return mtd.uTracker.InteractionData()
}

func (mtd *LUDecompositionMethod[T]) Run(inputMatrix [][]T) (det T, foundRoots []T, err error) {
	//  Checking if the given matrix is a square
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

	det = 1
	for eqIndex := 0; eqIndex < matrixSize; eqIndex++ {
		pivot := mtd.uMatrix[eqIndex][eqIndex]
		equation := mtd.uMatrix[eqIndex]

		for index := eqIndex + 1; index < matrixSize; index++ {
			nextEquation := mtd.uMatrix[index]
			factor := nextEquation[eqIndex] / pivot
			mtd.uMatrix[index] = mtd.subtractEquations(factor, nextEquation, equation)
			mtd.uTracker.registerMatrixTransformation(
				mtd.uMatrix, factor,
				uint8(index), uint8(eqIndex),
			)
			mtd.lMatrix[index][eqIndex] = factor
			mtd.lTracker.registerMatrixTransformation(
				mtd.lMatrix, factor, uint8(index), uint8(eqIndex), models.OpAttribution,
			)
		}
		det *= equation[eqIndex]
	}

	mtd.calculateXY(matrixSize, results)
	foundRoots = mtd.uTracker.Roots
	return
}

func (mtd *LUDecompositionMethod[T]) calculateXY(matrixSize int, results []T) (y, x []T) {
	// Solve Inferior triangle
	y = mtd.lTracker.forwardSubstitution(mtd.L(), matrixSize, results)

	// Solve superior triangle
	x = mtd.uTracker.backwardSubstitution(mtd.U(), matrixSize, y)

	return
}
