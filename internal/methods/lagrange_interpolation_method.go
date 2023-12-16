package methods

import (
	"errors"

	"github.com/jictyvoo/tec217_computing-methods/internal/models"
)

type LagrangeInterpolationMethod[T models.Numeric] struct {
	commonFuncZeroState[T]
	matrixHelper[T]
}

func (mtd *LagrangeInterpolationMethod[T]) Run(x, y []T, xx T) (yint T, err error) {
	if len(x) != len(y) {
		err = errors.New("x and y must have the same length")
		return
	}

	//sampleSize := len(x)
	//yint = mtd.zeroState() // Assuming there's Initialized to Zero
	//
	//for i := 0; i < sampleSize; i++ {
	//	term := mtd.identityState() // Assuming there's Initialized to One
	//	for j := 0; j < sampleSize; j++ {
	//		if i != j {
	//			term = term * (xx - x[j]) / (x[i] - x[j])
	//		}
	//	}
	//	yint += term * y[i]
	//	mtd.registerInteraction(x, uint32(i), 0, yint, xx)
	//}

	return
}
