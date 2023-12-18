package methods

import (
	"errors"

	"github.com/jictyvoo/tec217_computing-methods/internal/models"
)

type LagrangeInterpolationMethod[T models.Numeric] struct {
	commonFuncZeroState[T]
}

func (mtd *LagrangeInterpolationMethod[T]) Run(x, y []T, xx T) (result T, err error) {
	if len(x) != len(y) {
		err = errors.New("x and y must have the same length")
		return
	}

	sampleSize := len(x)

	for ctrlIndex := 0; ctrlIndex < sampleSize; ctrlIndex++ {
		term := y[ctrlIndex]

		// Start to perform the prod calculation
		for index := 0; index < sampleSize; index++ {
			if ctrlIndex != index {
				term *= (xx - x[index]) / (x[ctrlIndex] - x[index])
			}
		}
		result += term
		mtd.registerInteraction(x, uint32(ctrlIndex), 0, result, xx)
	}

	return
}
