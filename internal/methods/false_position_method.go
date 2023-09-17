package methods

import (
	"errors"
	"math"

	"github.com/jictyvoo/tec217_computing-methods/internal/models"
	"github.com/jictyvoo/tec217_computing-methods/internal/models/mtderrs"
)

type FalsePositionMethod struct {
	commonFuncZeroState[float64]
}

func (mtd *FalsePositionMethod) Run(
	a, b float64,
	fX models.RootFunctionCallback[float64], epsilon float64, maxIterations uint32,
) (result float64, totalIteration uint32, err error) {
	if fX(a)*fX(b) >= 0 {
		err = errors.New("there's no root on given interval\n")
		return
	}

	var absoluteError float64 = 1
	for math.Abs(absoluteError) > epsilon && totalIteration < maxIterations {
		totalIteration += 1
		var (
			fResultA, fResultB = fX(a), fX(b)
			oldResult          = result
		)

		if fResultA == fResultB {
			err = errors.New("division by zero error")
			return
		}
		if result = b - (fResultB*(a-b))/(fResultA-fResultB); totalIteration > 1 {
			absoluteError = math.Abs(oldResult - result)
		}
		rootResult := fX(result)

		mtd.registerInteraction(
			[]float64{a, b}, totalIteration,
			absoluteError, rootResult, result,
		)

		if rootResult == 0 {
			return
		}

		if fResultA*rootResult < 0 {
			b = result
		} else {
			a = result
		}
	}

	mtd.finalResult = result
	if totalIteration >= maxIterations {
		err = mtderrs.ErrMaxIterations(totalIteration)
	}
	return
}
