package main

import (
	"errors"
	"fmt"
	"log/slog"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/jictyvoo/tec217_computing-methods/internal/methods"
	"github.com/jictyvoo/tec217_computing-methods/internal/models"
	"github.com/jictyvoo/tec217_computing-methods/pkg/views"
)

type integrationQuestion[T models.Numeric] func(T) T

func (quest integrationQuestion[T]) trapezoidal(prefix string, a, b T) {
	var (
		buffer strings.Builder
		method = struct {
			simple, compound2, compound4 methods.IntegralTrapezoidalMethod[T]
		}{
			compound2: methods.IntegralTrapezoidalMethod[T]{
				Kind:              methods.IntegrationTrapezoidalCompound,
				NumberSubdivision: 2,
			},
			compound4: methods.IntegralTrapezoidalMethod[T]{
				Kind:              methods.IntegrationTrapezoidalCompound,
				NumberSubdivision: 4,
			},
		}
		fX      = quest
		results [3]struct {
			value T
			err   error
		}
	)

	results[0].value, results[0].err = method.simple.Run(a, b, fX)
	results[1].value, results[1].err = method.compound2.Run(a, b, fX)
	results[2].value, results[2].err = method.compound4.Run(a, b, fX)

	if err := errors.Join(results[0].err, results[1].err, results[2].err); err != nil {
		views.ReportError(err)
		return
	}

	views.ReportResult(
		prefix+" - Trapezoidal Simple Integration",
		&buffer, uint32(0), results[0].value,
	)
	views.ReportResult(
		prefix+" - Trapezoidal Compound Integration n=2",
		&buffer, uint32(2), results[1].value,
	)
	views.ReportResult(
		prefix+" - Trapezoidal Compound Integration n=4",
		&buffer, uint32(4), results[2].value,
	)
}

func (quest integrationQuestion[T]) simpson(prefix string, a, b T) {
	var (
		buffer strings.Builder
		method = struct {
			simple13, simple38, compound13 methods.IntegralSimpsonMethod[T]
		}{
			simple13: methods.IntegralSimpsonMethod[T]{
				Kind: methods.IntegrationSimpson13,
			},
			compound13: methods.IntegralSimpsonMethod[T]{
				Kind:              methods.IntegrationSimpsonCompound13,
				NumberSubdivision: 4,
			},
			simple38: methods.IntegralSimpsonMethod[T]{
				Kind: methods.IntegrationSimpson38,
			},
		}
		fX      = quest
		results [3]struct {
			value T
			err   error
		}
	)

	results[0].value, results[0].err = method.simple13.Run(a, b, fX)
	results[1].value, results[1].err = method.compound13.Run(a, b, fX)
	results[2].value, results[2].err = method.simple38.Run(a, b, fX)

	if err := errors.Join(results[0].err, results[1].err, results[2].err); err != nil {
		views.ReportError(err)
		return
	}

	views.ReportResult(prefix+"- Simpson Integration 1/3", &buffer, uint32(13), results[0].value)
	views.ReportResult(
		prefix+"- Simpson Integration 1/3 n=4", &buffer,
		uint32(113), results[1].value,
	)
	views.ReportResult(prefix+"- Simpson Integration 3/8", &buffer, uint32(38), results[2].value)
}

func quest01IntegralTrapezoidalSimpson() {
	var quest01 integrationQuestion[float64] = func(x float64) float64 {
		return 1 - math.Pow(math.E, -x)
	}
	const prefix = "quest01"
	quest01.trapezoidal(prefix, 0, 4)
	quest01.simpson(prefix, 0, 4)
}

func quest02IntegralTrapezoidalSimpson() {
	var quest02 integrationQuestion[float64] = func(x float64) float64 {
		return 8 + (4 * math.Cos(x))
	}
	const prefix = "quest02"
	quest02.trapezoidal(prefix, 0, math.Pi/2)
	quest02.simpson(prefix, 0, math.Pi/2)
}

func quest03IntegralTrapezoidalSimpson() {
	var quest03 integrationQuestion[float64] = func(x float64) float64 {
		return 1 - x - (4 * math.Pow(x, 3)) + (2 * math.Pow(x, 5))
	}
	const prefix = "quest03"
	quest03.trapezoidal(prefix, -2, 4)
	quest03.simpson(prefix, -2, 4)
}

func quest04IntegralTrapezoidalSimpson() {
	var quest04 integrationQuestion[float64] = func(x float64) float64 {
		return math.Pow(math.E, -x)
	}

	var method methods.IntegralTrapezoidalMethod[float64]
	result, err := method.Run(0, 1.2, quest04)
	if err != nil {
		views.ReportError(err)
	}

	views.ReportSimple("Question 04 - Simple Trapeizodal 0 ~ 1.2", result)

	limits := [...]float64{0, 0.1, 0.3, 0.5, 0.7, 0.95, 1.2}
	for index := 0; index < len(limits)-1; index++ {
		a := limits[index]
		b := limits[index+1]
		prefix := fmt.Sprintf("%s %.2f ~ %.2f", "quest04", a, b)
		quest04.trapezoidal(prefix, a, b)
		quest04.simpson(prefix, a, b)
	}
}

func quest05IntegralTrapezoidalAndPolynomials() {
	t := []float64{1, 2, 3.25, 4.5, 6, 7, 8, 8.5, 9, 10}
	v := []float64{5, 6, 5.5, 7, 8.5, 8, 6, 7, 7, 5}

	// Convert time to seconds instead of minutes
	for index, inMinutes := range t {
		t[index] = inMinutes * 60
	}

	// Calculating distance traveled using the trapezoidal rule
	n := len(t) - 1 // Number of subintervals
	distance := 0.0
	for index := 0; index < n; index++ {
		h := t[index+1] - t[index]                    // Difference between time values
		distance += (h / 2) * (v[index] + v[index+1]) // Applying the trapezoidal rule
	}

	// Average speed calculation
	sum := 0.0
	for _, velocity := range v {
		sum += velocity
	}
	averageVelocity := sum / float64(len(v))

	views.ReportSimple("Distance", distance)
	views.ReportSimple("Average Velocity", averageVelocity)

	var method struct {
		polynomial          methods.MinimumSquarePolynomialRegressionMethod[float64]
		integralTrapezoidal methods.IntegralTrapezoidalMethod[float64]
	}

	result, err := method.polynomial.Run(t, v, float64(3))
	if err != nil {
		views.ReportError(err)
		return
	}
	views.ReportResult(
		"PolynomialRegression - "+strconv.FormatInt(int64(3), 10)+" degree",
		nil, uint32(1), result...,
	)

	fX := func(x float64) (funcResult float64) {
		for index, value := range result {
			var xValue float64 = 1
			if index > 0 {
				xValue = math.Pow(x, float64(index))
			}
			funcResult += value * xValue
		}

		return
	}

	/*fX = func(x float64) (funcResult float64) {
		return 4.850657832943775 + (0.0010048084709803518 * x) +
			(0.00004869867922122433 * math.Pow(x, 2)) + (-8.333763969120286e-8 * math.Pow(x, 3))
	}*/

	var integrationResult float64
	integrationResult, err = method.integralTrapezoidal.Run(60, 600, fX)
	if err != nil {
		views.ReportError(err)
	}

	views.ReportSimple("Question 05 - Simple Trapeizodal 1 ~ 10", integrationResult)
}

func main() {
	slog.SetDefault(
		slog.New(
			slog.NewTextHandler(
				os.Stdout,
				&slog.HandlerOptions{
					Level: slog.LevelInfo,
				},
			),
		),
	)

	quest01IntegralTrapezoidalSimpson()
	quest02IntegralTrapezoidalSimpson()
	quest03IntegralTrapezoidalSimpson()
	quest04IntegralTrapezoidalSimpson()
	quest05IntegralTrapezoidalAndPolynomials()
}
