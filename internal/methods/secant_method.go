package methods

import (
	"math"

	"github.com/jictyvoo/tec217_computing-methods/internal/models"
	"github.com/jictyvoo/tec217_computing-methods/internal/models/mtderrs"
)

type SecantMethod struct {
	commonFuncZeroState[float64]
}

func (mtd *SecantMethod) Run(
	x0, x1 float64, fX models.RootFunctionCallback[float64],
	epsilon float64, maxIterations uint32,
) (result float64, totalIteration uint32, err error) {
	var relativeError float64 = 1

	for math.Abs(relativeError) >= epsilon && totalIteration < maxIterations {
		totalIteration += 1
		fPrevious, fCurrent := fX(x0), fX(x1)

		// Apply the Secant Method formula
		result = ((x0 * fCurrent) - (x1 * fPrevious)) / (fCurrent - fPrevious)
		if result != 0 && totalIteration > 1 {
			relativeError = math.Abs(x0-result) / result
		}

		rootResult := fX(result)
		mtd.registerInteraction(
			[]float64{x0, x1}, totalIteration,
			relativeError, rootResult, result,
		)

		x0, x1 = result, x0
	}

	mtd.finalResult = result
	if totalIteration >= maxIterations {
		err = mtderrs.ErrMaxIterations(totalIteration)
	}
	return
}
