package methods

import (
	"errors"
	"math"

	"github.com/jictyvoo/tec217_computing-methods/internal/models"
	"github.com/jictyvoo/tec217_computing-methods/internal/models/mtderrs"
)

type BisectionMethod struct {
	commonState[float64]
}

func (mtd *BisectionMethod) Run(
	a, b float64,
	fX models.RootFunctionCallback[float64], epsilon float64, maxIterations uint32,
) (result float64, totalIteration uint32, err error) {
	// Check for signal change
	if fX(a)*fX(b) >= 0 {
		err = errors.New("you have not assumed right a and b")
		return
	}

	var absoluteError float64 = 1
	for absoluteError > epsilon && totalIteration < maxIterations {
		totalIteration += 1
		oldResult := result
		if result = (a + b) / 2; totalIteration > 1 {
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

		if fX(a)*rootResult < 0 {
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
