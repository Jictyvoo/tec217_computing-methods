package main

import (
	"errors"
	"log/slog"
	"os"
	"strings"

	"github.com/jictyvoo/tec217_computing-methods/internal/methods"
	"github.com/jictyvoo/tec217_computing-methods/internal/utils"
	"github.com/jictyvoo/tec217_computing-methods/pkg/views"
)

func quest01EDOEulerHeunMeanPoint() {
	const (
		L = 50e-3
		R = 20
		E = 10
		//
		a  = 0.0    // Initial timestamp
		b  = 0.02   // Final timestamp
		h  = 0.0001 // Step
		y0 = 0.0    //
	)

	var (
		buffer struct {
			euler, heun, meanPoint strings.Builder
		}
		method struct {
			euler, heun, meanPoint methods.EDOMultiMethod[float64]
		}
		results struct {
			euler, heun, meanPoint struct {
				x, y float64
				err  error
			}
		}

		fX = func(t, i float64) float64 {
			return (E - R*i) / L
		}
	)

	results.euler.x, results.euler.y, results.euler.err = method.euler.Euler(a, b, h, y0, fX)
	results.heun.x, results.heun.y, results.heun.err = method.heun.Heun(a, b, h, y0, fX)
	results.meanPoint.x, results.meanPoint.y, results.meanPoint.err = method.meanPoint.MeanPoint(
		a, b, h, y0, fX,
	)

	if err := errors.Join(
		results.euler.err,
		utils.WriteInteractionAsCSV(&buffer.euler, method.euler.InteractionData()),
		utils.WriteInteractionAsCSV(&buffer.heun, method.heun.InteractionData()),
		utils.WriteInteractionAsCSV(&buffer.meanPoint, method.meanPoint.InteractionData()),
	); err != nil {
		views.ReportError(err)
		return
	}

	views.ReportResult("EDO - Euler", &buffer.euler, uint32(0), results.euler.x, results.euler.y)
	views.ReportResult("EDO - Heun", &buffer.heun, uint32(0), results.heun.x, results.heun.y)
	views.ReportResult(
		"EDO - Mean Points",
		&buffer.meanPoint, uint32(0),
		results.meanPoint.x, results.meanPoint.y,
	)
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

	quest01EDOEulerHeunMeanPoint()
}
