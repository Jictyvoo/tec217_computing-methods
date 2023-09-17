package methods

import (
	"math"

	"github.com/jictyvoo/tec217_computing-methods/internal/models"
	"github.com/jictyvoo/tec217_computing-methods/internal/models/mtderrs"
)

type LinearIterationMethod struct {
	commonFuncZeroState[float64]
}

func (mtd *LinearIterationMethod) Run(
	x0 float64, fX, gX models.RootFunctionCallback[float64],
	epsilon float64, maxIterations uint32,
) (result float64, totalIteration uint32, err error) {
	var absoluteError float64 = 1
	result = x0
	for math.Abs(absoluteError) >= epsilon && totalIteration < maxIterations {
		totalIteration += 1
		oldResult := result
		result = gX(result)

		if result != 0 && totalIteration > 1 {
			absoluteError = math.Abs(oldResult-result) / result
		}
		rootResult := fX(result)

		mtd.registerInteraction(
			[]float64{x0}, totalIteration,
			absoluteError, rootResult, result,
		)

		if rootResult == 0 {
			return
		}

	}

	mtd.finalResult = result
	if totalIteration >= maxIterations {
		err = mtderrs.ErrMaxIterations(totalIteration)
	}
	return
}
