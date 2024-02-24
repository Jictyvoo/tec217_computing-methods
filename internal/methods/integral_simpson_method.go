package methods

import (
	"errors"

	"github.com/jictyvoo/tec217_computing-methods/internal/models"
)

type integrationSimpsonKind uint8

const (
	IntegrationSimpson13 integrationSimpsonKind = iota
	IntegrationSimpson38
	IntegrationSimpsonCompound13
)

type (
	IntegralSimpsonMethod[T models.Numeric] struct {
		Kind              integrationSimpsonKind
		NumberSubdivision uint64
		commonFuncZeroState[T]
	}
)

func (mtd *IntegralSimpsonMethod[T]) simple13(a, b T, f funcCallback[T]) (result T, err error) {
	var (
		h  = (b - a) / 2
		f1 = f(a)
		f2 = f((a + b) / 2)
		f3 = f(b)
	)

	result = h * (f1 + 4*f2 + f3) / 3
	return
}

func (mtd *IntegralSimpsonMethod[T]) simple38(a, b T, f funcCallback[T]) (result T, err error) {
	var (
		h  = (b - a) / 3
		f1 = f(a)
		f2 = f((2*a + b) / 3)
		f3 = f((a + 2*b) / 3)
		f4 = f(b)
	)
	result = 3 * h * (f1 + 3*(f2+f3) + f4) / 8
	return
}

func (mtd *IntegralSimpsonMethod[T]) compound13(a, b T, f funcCallback[T]) (result T, err error) {
	if mtd.NumberSubdivision%2 != 0 || mtd.NumberSubdivision == 0 {
		err = errors.New("the given n should be a even number")
		return
	}
	h := (b - a) / T(mtd.NumberSubdivision)
	var sum T
	for i := uint64(1); i < mtd.NumberSubdivision; i++ {
		xi := a + T(i)*h
		switch {
		case i == 1 || i == mtd.NumberSubdivision+1:
			sum += f(xi)
		case i%2 == 0:
			sum += 4 * f(xi)
		default:
			sum += 2 * f(xi)
		}

	}
	result = h * sum / 3
	return
}

func (mtd *IntegralSimpsonMethod[T]) Run(a, b T, f func(T) T) (result T, err error) {
	switch mtd.Kind {
	case IntegrationSimpson13:
		return mtd.simple13(a, b, f)
	case IntegrationSimpson38:
		return mtd.simple38(a, b, f)
	case IntegrationSimpsonCompound13:
		return mtd.compound13(a, b, f)
	}

	return
}
