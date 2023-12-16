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

func quest01NewtonInterpolation() {
	var (
		buffer strings.Builder
		method = methods.NewtonInterpolationMethod[float64]{}
	)

	result, err := method.Run(
		[]float64{0, 0.5, 1},
		[]float64{1, 2.12, 3.55},
		0.7,
	)

	if err = errors.Join(err, utils.WriteInteractionAsCSV(&buffer, method.InteractionData())); err != nil {
		views.ReportError(err)
		return
	}
	views.ReportResult("NewtonInterpolation", &buffer, uint32(1), result)
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
	quest01NewtonInterpolation()
}
