package methods

import (
	"errors"

	"github.com/jictyvoo/tec217_computing-methods/internal/models"
)

type NewtonInterpolationMethod[T models.Numeric] struct {
	commonFuncZeroState[T]
	matrixHelper[T]
}

func (mtd *NewtonInterpolationMethod[T]) Run(x, y []T, xx T) (yint T, err error) {
	sampleSize := len(x)
	// Check for consistent input
	if sampleSize != len(y) {
		err = errors.New("x and y must have the same length")
		return
	}

	// Initialize b as a zero 2D slice
	b := mtd.newEmptyMatrix(uint8(sampleSize))

	// Set the first column of b as y
	for i := 0; i < sampleSize; i++ {
		b[i][0] = y[i]
	}

	// Calculate finite differences
	for j := 1; j < sampleSize; j++ {
		for i := 0; i < sampleSize-j; i++ {
			b[i][j] = (b[i+1][j-1] - b[i][j-1]) / (x[i+j] - x[i])
			mtd.registerInteraction(x, uint32(j), 0, b[i][j], 0)
		}
	}

	// Apply finite differences to interpolate
	var xt T = 1
	yint = b[0][0]
	for index := 0; index < sampleSize-1; index++ {
		xt *= xx - x[index]
		yint += b[0][index+1] * xt
		mtd.registerInteraction(x, uint32(index), 0, yint, xx)
	}

	return yint, nil
}
