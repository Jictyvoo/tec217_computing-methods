package main

import (
	"errors"
	"log/slog"
	"math"
	"os"
	"strings"

	"github.com/jictyvoo/tec217_computing-methods/internal/methods"
	"github.com/jictyvoo/tec217_computing-methods/internal/utils"
	"github.com/jictyvoo/tec217_computing-methods/pkg/views"
)

func quest01SocialControl() {
	const (
		epsilon = 1e-8
		N       = 10000

		Pumax = 80000
		ku    = 0.05
		Pumin = 110000
		Psmax = 320000
		P0    = 10000
		ks    = 0.09
	)

	var (
		buffer strings.Builder
	)

	fX := func(x float64) float64 {
		return (Psmax / (1 + (((Psmax / P0) - 1) * math.Pow(math.E, -ks*x)))) - (1.2 * ((Pumax * math.Pow(math.E, -ku*x)) + Pumin))
	}

	// Bisection
	{
		buffer.Reset()
		const a, b = 0, 100
		method := methods.BisectionMethod{}
		result, totalInteractions, err := method.Run(
			a, b, fX, epsilon, N,
		)

		if err = errors.Join(err, utils.WriteInteractionAsCSV(&buffer, method.InteractionData())); err != nil {
			views.ReportError(err)
			return
		}
		views.ReportResult("SocialControl(BisectionMethod)", result, totalInteractions, &buffer)
	}

	// False Position
	{
		buffer.Reset()
		const a, b = 0, 100
		method := methods.FalsePositionMethod{}
		result, totalInteractions, err := method.Run(
			a, b, fX, epsilon, N,
		)

		if err = errors.Join(err, utils.WriteInteractionAsCSV(&buffer, method.InteractionData())); err != nil {
			views.ReportError(err)
			return
		}
		views.ReportResult("SocialControl(FalsePositionMethod)", result, totalInteractions, &buffer)
	}

	// Secant
	{
		buffer.Reset()
		const a, b = 0, 100
		method := methods.SecantMethod{}
		result, totalInteractions, err := method.Run(
			a, b, fX, epsilon, N,
		)

		if err = errors.Join(err, utils.WriteInteractionAsCSV(&buffer, method.InteractionData())); err != nil {
			views.ReportError(err)
			return
		}
		views.ReportResult("SocialControl(SecantMethod)", result, totalInteractions, &buffer)
	}

	// Linear Iteration
	{
		buffer.Reset()
		const x0 = 0

		gX := func(x float64) float64 {
			return math.Log(((32/((9.6*math.Exp(-0.05*x))+13.2))-1)/31) / -0.09
		}
		method := methods.LinearIterationMethod{}
		result, totalInteractions, err := method.Run(
			x0, fX, gX, epsilon, N,
		)

		if err = errors.Join(err, utils.WriteInteractionAsCSV(&buffer, method.InteractionData())); err != nil {
			views.ReportError(err)
			return
		}
		views.ReportResult("SocialControl(LinearIterationMethod)", result, totalInteractions, &buffer)
	}

	// Newton Raphson
	{
		buffer.Reset()
		const x0 = 0

		fX = func(x float64) float64 {
			return 32/(1+(31*math.Exp(-ks*x))) - 13.2 - (9.6 * math.Exp(-ku*x))
		}

		f1X := func(x float64) float64 {
			return (0.48 * math.Exp(-0.05*x)) + ((89.28 * math.Exp(0.09*x)) / (math.Pow(math.Exp(0.09*x)+31, 2)))
		}
		method := methods.NewtonRaphsonMethod{}
		result, totalInteractions, err := method.Run(
			x0, fX, f1X, epsilon, N,
		)

		if err = errors.Join(err, utils.WriteInteractionAsCSV(&buffer, method.InteractionData())); err != nil {
			views.ReportError(err)
			return
		}
		views.ReportResult("SocialControl(NewtonRaphsonMethod)", result, totalInteractions, &buffer)
	}
}

func quest02ElectricCharge() {
	const (
		epsilon = 1e-8
		N       = 10000

		e0 = 8.9e-12
		F  = 1.25
		a  = 0.85
		q  = 2e-5
		Q  = q
	)

	var buffer strings.Builder

	fX := func(x float64) float64 {
		return ((q * Q * x) / ((4 * math.Pi * e0) * math.Sqrt(math.Pow(math.Pow(x, 2)+math.Pow(a, 2), 3)))) - F
	}

	// Bisection
	{
		buffer.Reset()
		const a, b = 0, 0.5
		method := methods.BisectionMethod{}
		result, totalInteractions, err := method.Run(
			a, b, fX, epsilon, N,
		)

		if err = errors.Join(err, utils.WriteInteractionAsCSV(&buffer, method.InteractionData())); err != nil {
			views.ReportError(err)
			return
		}
		views.ReportResult("ElectricCharge(BisectionMethod)", result, totalInteractions, &buffer)
	}

	// False Position
	{
		buffer.Reset()
		const a, b = 0, 0.5
		method := methods.FalsePositionMethod{}
		result, totalInteractions, err := method.Run(
			a, b, fX, epsilon, N,
		)

		if err = errors.Join(err, utils.WriteInteractionAsCSV(&buffer, method.InteractionData())); err != nil {
			views.ReportError(err)
			return
		}
		views.ReportResult("ElectricCharge(FalsePositionMethod)", result, totalInteractions, &buffer)
	}

	// Secant
	{
		buffer.Reset()
		const a, b = 0, 0.5
		method := methods.SecantMethod{}
		result, totalInteractions, err := method.Run(
			a, b, fX, epsilon, N,
		)

		if err = errors.Join(err, utils.WriteInteractionAsCSV(&buffer, method.InteractionData())); err != nil {
			views.ReportError(err)
			return
		}
		views.ReportResult("ElectricCharge(SecantMethod)", result, totalInteractions, &buffer)
	}

	// Linear Iteration
	{
		buffer.Reset()
		const x0 = 0

		gX := func(x float64) float64 {
			return (F * ((4 * math.Pi * e0) * math.Sqrt(math.Pow(math.Pow(x, 2)+math.Pow(a, 2), 3)))) / (q * Q)
		}
		method := methods.LinearIterationMethod{}
		result, totalInteractions, err := method.Run(
			x0, fX, gX, epsilon, N,
		)

		if err = errors.Join(err, utils.WriteInteractionAsCSV(&buffer, method.InteractionData())); err != nil {
			views.ReportError(err)
			return
		}
		views.ReportResult("ElectricCharge(LinearIterationMethod)", result, totalInteractions, &buffer)
	}

	// Newton Raphson
	{
		buffer.Reset()
		const x0 = 0

		f1X := func(x float64) float64 {
			return 3.57652/math.Sqrt(math.Pow(math.Pow(x, 2)+0.7255, 3)) - ((10.7295 * math.Pow(x, 2)) / math.Sqrt(math.Pow(math.Pow(x, 2)+0.7255, 2)))
		}
		method := methods.NewtonRaphsonMethod{}
		result, totalInteractions, err := method.Run(
			x0, fX, f1X, epsilon, N,
		)

		if err = errors.Join(err, utils.WriteInteractionAsCSV(&buffer, method.InteractionData())); err != nil {
			views.ReportError(err)
			return
		}
		views.ReportResult("ElectricCharge(NewtonRaphsonMethod)", result, totalInteractions, &buffer)
	}
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
	quest01SocialControl()
	quest02ElectricCharge()
}
