package main

import (
	"errors"
	"log/slog"
	"math"
	"os"
	"strings"

	"github.com/jictyvoo/tec217_computing-methods/internal/methods"
	"github.com/jictyvoo/tec217_computing-methods/pkg/views"
)

func quest01IntegralTrapezoidal() {
	var (
		buffer strings.Builder
		method = struct {
			simple, compound2, compound4 methods.IntegralTrapezoidalMethod[float64]
		}{
			compound2: methods.IntegralTrapezoidalMethod[float64]{
				Kind:              methods.IntegrationTrapezoidalCompound,
				NumberSubdivision: 2,
			},
			compound4: methods.IntegralTrapezoidalMethod[float64]{
				Kind:              methods.IntegrationTrapezoidalCompound,
				NumberSubdivision: 4,
			},
		}
		fX = func(x float64) float64 {
			return 1 - math.Pow(math.E, -x)
		}
		results [3]struct {
			value float64
			err   error
		}
	)

	results[0].value, results[0].err = method.simple.Run(0, 4, fX)
	results[1].value, results[1].err = method.compound2.Run(0, 4, fX)
	results[2].value, results[2].err = method.compound4.Run(0, 4, fX)

	if err := errors.Join(results[0].err, results[1].err, results[2].err); err != nil {
		views.ReportError(err)
		return
	}

	views.ReportResult("TrapezoidalIntegration", &buffer, uint32(0), results[0].value)
	views.ReportResult("TrapezoidalIntegration", &buffer, uint32(2), results[1].value)
	views.ReportResult("TrapezoidalIntegration", &buffer, uint32(4), results[2].value)
}

func quest01IntegralSimpson() {
	var (
		buffer strings.Builder
		method = struct {
			simple13, simple38, compound13 methods.IntegralSimpsonMethod[float64]
		}{
			simple13: methods.IntegralSimpsonMethod[float64]{
				Kind: methods.IntegrationSimpson13,
			},
			compound13: methods.IntegralSimpsonMethod[float64]{
				Kind:              methods.IntegrationSimpsonCompound13,
				NumberSubdivision: 4,
			},
			simple38: methods.IntegralSimpsonMethod[float64]{
				Kind: methods.IntegrationSimpson38,
			},
		}
		fX = func(x float64) float64 {
			return 1 - math.Pow(math.E, -x)
		}
		results [3]struct {
			value float64
			err   error
		}
	)

	results[0].value, results[0].err = method.simple13.Run(0, 4, fX)
	results[1].value, results[1].err = method.compound13.Run(0, 4, fX)
	results[2].value, results[2].err = method.simple38.Run(0, 4, fX)

	if err := errors.Join(results[0].err, results[1].err, results[2].err); err != nil {
		views.ReportError(err)
		return
	}

	views.ReportResult("SimpsonIntegration", &buffer, uint32(13), results[0].value)
	views.ReportResult("SimpsonIntegration", &buffer, uint32(113), results[1].value)
	views.ReportResult("SimpsonIntegration", &buffer, uint32(38), results[2].value)
}

func quest02IntegralSimpson() {
	const (
		n    float64 = 8
		R0           = 0.5
		vMax         = 1.5
	)

	var (
		buffer strings.Builder
		method = struct {
			simple13, simple38 methods.IntegralSimpsonMethod[float64]
		}{
			simple13: methods.IntegralSimpsonMethod[float64]{
				Kind: methods.IntegrationSimpson13,
			},
			simple38: methods.IntegralSimpsonMethod[float64]{
				Kind: methods.IntegrationSimpson38,
			},
		}
		fX = func(r float64) float64 {
			return r * math.Pow(1-(r/R0), 1/n)
		}
		results [2]struct {
			value float64
			err   error
		}
	)

	results[0].value, results[0].err = method.simple13.Run(0, R0, fX)
	results[1].value, results[1].err = method.simple38.Run(0, R0, fX)

	if err := errors.Join(results[0].err, results[1].err); err != nil {
		views.ReportError(err)
		return
	}

	var vMean [2]float64
	for index, result := range results {
		vMean[index] = (2 * vMax * result.value) / math.Pow(R0, 2)
	}

	views.ReportResult("SimpsonIntegration", &buffer, uint32(13), results[0].value)
	views.ReportSimple("Mean Velocity 1/3", vMean[0])
	views.ReportResult("SimpsonIntegration", &buffer, uint32(38), results[1].value)
	views.ReportSimple("Mean Velocity 3/8", vMean[1])
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

	quest01IntegralTrapezoidal()
	quest01IntegralSimpson()
	quest02IntegralSimpson()
}
