package methods

import (
	"errors"

	"github.com/jictyvoo/tec217_computing-methods/internal/models"
	"github.com/jictyvoo/tec217_computing-methods/internal/utils"
)

type MinimumSquarePolynomialRegressionMethod[T models.Numeric] struct {
	commonFuncZeroState[T]
}

func (mtd *MinimumSquarePolynomialRegressionMethod[T]) power(x, y T) T {
	var result T = 1
	for i := 0; i < int(y); i++ {
		result *= x
	}
	return result
}

func (mtd *MinimumSquarePolynomialRegressionMethod[T]) solveSuperiorTriangle(
	n int, a [][]T, b []T,
) {
	for k := 0; k < n-1; k++ {
		for i := k + 1; i < n; i++ {
			m := a[i][k] / a[k][k]

			// Change a matrix in-memory
			a[i][k] = 0
			for j := k + 1; j < n; j++ {
				a[i][j] -= m * a[k][j]
			}

			// Change b slice in-memory
			b[i] -= m * b[k]
		}
	}
}

func (mtd *MinimumSquarePolynomialRegressionMethod[T]) Run(
	x, y []T, polyDegree T,
) (result []T, err error) {
	sampleSize := len(x)
	// Check for consistent input
	if sampleSize != len(y) {
		err = errors.New("x and y must have the same length")
		return
	}

	n := int(polyDegree + 1)
	a := make([][]T, n)
	for i := range a {
		a[i] = make([]T, n)
	}
	a[0][0] = T(sampleSize)

	b := make([]T, n)

	// Generating a linear system
	for _, value := range y {
		b[0] += value
	}
	for index := 1; index <= n; index++ {
		k1 := T(index - 1)
		k2 := T(index)
		for subIndex := 2; subIndex <= n; subIndex++ {
			sumAuxiliary := utils.ReduceSlice(
				x, func(old, value T) T { return old + mtd.power(value, k2) },
			)
			a[index-1][subIndex-1] = sumAuxiliary
			a[subIndex-1][index-1] = a[index-1][subIndex-1]
			k2++
		}

		var sumAuxiliary T
		if index != 1 {
			for indexIterator := 0; indexIterator < sampleSize; indexIterator++ {
				sumAuxiliary += y[indexIterator] * mtd.power(x[indexIterator], k1)
			}
			b[index-1] = sumAuxiliary
		}
	}

	// Start the linear system solution

	// Superior Triangular Matrix
	mtd.solveSuperiorTriangle(n, a, b)

	// Start the system solution calculation
	result = make([]T, n)
	result[n-1] = b[n-1] / a[n-1][n-1]
	for k := n - 2; k >= 0; k-- {
		var s T
		for j := k + 1; j < n; j++ {
			s += a[k][j] * result[j]
		}
		result[k] = (b[k] - s) / a[k][k]
	}

	return
}
