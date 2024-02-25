package methods

import (
	"github.com/jictyvoo/tec217_computing-methods/internal/models"
)

type edoAlgorithmKind uint8

const (
	EDOAlgorithmEuler edoAlgorithmKind = iota
	EDOAlgorithmHeun
	EDOAlgorithmMeanPoint
)

type (
	EDOMultiMethod[T models.Numeric] struct {
		Kind edoAlgorithmKind
		commonFuncZeroState[T]
	}
)

func (mtd *EDOMultiMethod[T]) Euler(a, b, h, y0 T, f func(T, T) T) (x, y T, err error) {
	// Calculate number of steps
	n := int((b - a) / h)

	// Initialize slices to store x and y values
	workArr := struct {
		x, y []T
	}{
		x: make([]T, n+1),
		y: make([]T, n+1),
	}

	// Initialize x[0] and y[0]
	workArr.x[0] = a
	workArr.y[0] = y0

	// Iterate over steps
	for index := 0; index < n; index++ {
		// Calculate slope at current point
		F := f(workArr.x[index], workArr.y[index])

		// Update x[index+1]
		workArr.x[index+1] = workArr.x[index] + h
		// Update y[index+1]
		workArr.y[index+1] = workArr.y[index] + F*h

		mtd.registerInteraction([]T{a, b}, uint32(index), 0, workArr.y[index], workArr.x[index])
	}

	// Return x and y as pairs of solution points
	x, y = workArr.x[n], workArr.y[n]
	return
}

func (mtd *EDOMultiMethod[T]) Heun(a, b, h, y0 T, f func(T, T) T) (x, y T, err error) {
	// Calculate number of steps
	n := int((b - a) / h)

	// Initialize slices to store x and y values
	workArr := struct {
		x, y []T
	}{
		x: make([]T, n+1),
		y: make([]T, n+1),
	}

	// Initialize x[0] and y[0]
	workArr.x[0] = a
	workArr.y[0] = y0

	// Iterate over steps
	for index := 0; index < n; index++ {
		// Calculate slope at current point
		F1 := f(workArr.x[index], workArr.y[index])

		// Calculate slope at next point
		F2 := f(workArr.x[index]+h, workArr.y[index]+(F1*h))

		// Update y[index+1]
		workArr.y[index+1] = workArr.y[index] + (F1+F2)*(h/2)

		// Update x[index+1]
		workArr.x[index+1] = workArr.x[index] + h

		mtd.registerInteraction([]T{a, b}, uint32(index), 0, workArr.y[index], workArr.x[index])
	}

	// Return x and y as pairs of solution points
	x, y = workArr.x[n-1], workArr.y[n-1]
	return
}

func (mtd *EDOMultiMethod[T]) MeanPoint(a, b, h, y0 T, f func(T, T) T) (x, y T, err error) {
	// Calculate number of steps
	n := int((b - a) / h)

	// Initialize slices to store x and y values
	workArr := struct {
		x, y []T
	}{
		x: make([]T, n+1),
		y: make([]T, n+1),
	}

	// Initialize x[0] and y[0]
	workArr.x[0] = a
	workArr.y[0] = y0

	// Iterate over steps
	for index := 0; index < n; index++ {
		// Calculate slope at current point
		F1 := f(workArr.x[index], workArr.y[index])

		// Calculate slope at mean point
		F2 := f(workArr.x[index]+(h/2), workArr.y[index]+(F1*(h/2)))

		// Update workArr.y[index+1]
		workArr.y[index+1] = workArr.y[index] + (F1+F2)*(h/2)

		// Update workArr.x[index+1]
		workArr.x[index+1] = workArr.x[index] + h
		mtd.registerInteraction([]T{a, b}, uint32(index), 0, workArr.y[index], workArr.x[index])
	}

	x, y = workArr.x[n], workArr.y[n]
	return
}

func (mtd *EDOMultiMethod[T]) Run(a, b, h, y0 T, f func(T, T) T) (x, y T, err error) {
	switch mtd.Kind {
	case EDOAlgorithmEuler:
		return mtd.Euler(a, b, h, y0, f)
	case EDOAlgorithmHeun:
		return mtd.Heun(a, b, h, y0, f)
	case EDOAlgorithmMeanPoint:
		return mtd.MeanPoint(a, b, h, y0, f)
	}

	return
}
