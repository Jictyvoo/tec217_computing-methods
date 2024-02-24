package methods

import (
	"github.com/jictyvoo/tec217_computing-methods/internal/models"
)

type integrationTrapezoidalKind uint8

const (
	IntegrationTrapezoidalSimple integrationTrapezoidalKind = iota
	IntegrationTrapezoidalCompound
)

type IntegralTrapezoidalMethod[T models.Numeric] struct {
	commonFuncZeroState[T]
	Kind              integrationTrapezoidalKind
	NumberSubdivision uint64
}

func (mtd *IntegralTrapezoidalMethod[T]) simpleCalculation(
	a, b T,
	f funcCallback[T],
) (result T, _ error) {
	h := b - a
	result = h * (f(a) + f(b)) / 2
	return
}

func (mtd *IntegralTrapezoidalMethod[T]) compoundCalculation(
	a, b T,
	f funcCallback[T],
) (result T, _ error) {
	h := (b - a) / T(mtd.NumberSubdivision)
	previousX := a
	sum := f(previousX)
	for i := uint64(1); i < mtd.NumberSubdivision; i++ {
		previousX = previousX + h
		sum += 2 * f(previousX)
	}

	x := previousX + h
	sum += f(x)
	result = h * (sum / 2)
	return
}

func (mtd *IntegralTrapezoidalMethod[T]) Run(a, b T, f func(T) T) (result T, err error) {
	switch mtd.Kind {
	case IntegrationTrapezoidalCompound:
		return mtd.compoundCalculation(a, b, f)
	case IntegrationTrapezoidalSimple:
		return mtd.simpleCalculation(a, b, f)
	}

	return
}
