package methods

import (
	"errors"

	"github.com/jictyvoo/tec217_computing-methods/internal/models"
)

type MinimumSquareRegressionMethod[T models.Numeric] struct {
	commonFuncZeroState[T]
}

func (mtd *MinimumSquareRegressionMethod[T]) Run(x, y []T) (a [2]T, r2 T, err error) {
	// Check for consistent input
	if len(x) != len(y) {
		err = errors.New("x and y must have the same length")
		return
	}

	sampleSize := len(x)
	sums := struct{ x, y, xPow2, yPow2, xy T }{}
	for index := 0; index < sampleSize; index++ {
		values := struct{ x, y T }{x: x[index], y: y[index]}
		sums.x += values.x
		sums.y += values.y
		sums.xy += values.x * values.y
		sums.xPow2 += values.x * values.x
		sums.yPow2 += values.y * values.y
	}

	// Calculate coefficients
	a[0] = ((T(sampleSize) * sums.xy) - (sums.x * sums.y)) / ((T(sampleSize) * sums.xPow2) - (sums.x * sums.x))
	a[1] = (sums.y / T(sampleSize)) - (a[0] * (sums.x / T(sampleSize)))

	var (
		averageResiduum      T
		linearAdjustResiduum T
	)

	for index := 0; index < sampleSize; index++ {
		averageDiff := y[index] - (sums.y / T(sampleSize))
		averageResiduum += averageDiff * averageDiff

		linearTempVal := y[index] - a[1] - (a[0] * x[index])
		linearAdjustResiduum += linearTempVal * linearTempVal
	}

	// stdDeviationAverage := math.Sqrt(float64(averageResiduum / T(sampleSize-1)))
	// stdDeviationLinear := math.Sqrt(float64(linearAdjustResiduum / T(sampleSize-2)))

	r2 = (averageResiduum - linearAdjustResiduum) / averageResiduum

	return
}
