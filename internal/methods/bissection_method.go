package methods

import (
	"errors"
	"math"
	"strconv"

	"github.com/jictyvoo/tec217_computing-methods/internal/models"
)

type BisectionMethod struct {
	finalResult  float64
	interactions []models.InteractionData[float64]
}

func (mtd BisectionMethod) InteractionData() []models.InteractionData[float64] {
	return mtd.interactions
}

func (mtd *BisectionMethod) Run(
	a, b float64,
	fX models.RootFunctionCallback[float64], epsilon float64, maxIterations uint32,
) (result float64, totalIteration uint32, err error) {
	// Check for signal change
	if fX(a)*fX(b) >= 0 {
		err = errors.New("you have not assumed right a and b\n")
		return
	}

	var absoluteError float64 = 1
	for absoluteError > epsilon && totalIteration < maxIterations {
		totalIteration += 1
		oldResult := result
		if result = (a + b) / 2; totalIteration > 1 {
			absoluteError = math.Abs(oldResult - result)
		}

		fR := fX(result)
		mtd.interactions = append(mtd.interactions, models.InteractionData[float64]{
			Interaction:    uint64(totalIteration),
			InputValues:    []float64{a, b},
			AbsoluteError:  absoluteError,
			FunctionResult: fR,
			Value:          result,
		})

		if fR == 0 {
			return
		}

		if fX(a)*fR < 0 {
			b = result
		} else {
			a = result
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
