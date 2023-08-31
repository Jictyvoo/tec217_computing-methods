package methods

import (
	"errors"
	"math"
	"strconv"

	"github.com/jictyvoo/tec217_computing-methods/internal/models"
)

type LinearIterationMethod struct {
	commonState[float64]
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

		mtd.interactions = append(mtd.interactions, models.InteractionData[float64]{
			Interaction:    uint64(totalIteration),
			InputValues:    []float64{x0},
			RelativeError:  absoluteError,
			FunctionResult: rootResult,
			Value:          result,
		})

		if rootResult == 0 {
			return
		}

	}

	mtd.finalResult = result
	if totalIteration >= maxIterations {
		err = errors.New(
			"failed to find root after " + strconv.Itoa(int(totalIteration)) + " iterations",
		)
	}
	return
}
