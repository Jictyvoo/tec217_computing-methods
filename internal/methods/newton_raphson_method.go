package methods

import (
	"errors"
	"math"
	"strconv"

	"github.com/jictyvoo/tec217_computing-methods/internal/models"
)

type NewtonRaphsonMethod struct {
	commonState[float64]
}

func (mtd *NewtonRaphsonMethod) Run(
	x0 float64, fX, f1X models.RootFunctionCallback[float64],
	epsilon float64, maxIterations uint32,
) (result float64, totalIteration uint32, err error) {
	var relativeError float64 = 1
	result = x0
	for math.Abs(relativeError) >= epsilon && totalIteration < maxIterations {
		totalIteration += 1
		oldResult := result
		result -= fX(oldResult) / f1X(oldResult)

		if result != 0 && totalIteration > 1 {
			relativeError = math.Abs(oldResult-result) / result
		}
		rootResult := fX(result)

		mtd.interactions = append(mtd.interactions, models.InteractionData[float64]{
			Interaction:    uint64(totalIteration),
			InputValues:    []float64{x0},
			RelativeError:  relativeError,
			FunctionResult: rootResult,
			Value:          result,
		})
	}

	mtd.finalResult = result
	if totalIteration >= maxIterations {
		err = errors.New(
			"failed to find root after " + strconv.Itoa(int(totalIteration)) + " iterations",
		)
	}
	return
}
