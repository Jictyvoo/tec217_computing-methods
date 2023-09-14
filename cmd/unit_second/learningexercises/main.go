package main

import (
	"log/slog"
	"os"
	"strings"

	"github.com/jictyvoo/tec217_computing-methods/internal/methods"
	"github.com/jictyvoo/tec217_computing-methods/pkg/views"
)

func quest01Gauss() {
	var (
		buffer strings.Builder
		method = methods.GaussEliminationMethod[float64]{}
	)

	det, _, err := method.Run(
		[][]float64{
			{10, 2, -1, 27},
			{-3, -5, 2, -61.5},
			{1, 1, 6, -21.5},
		},
	)
	if err != nil {
		views.ReportError(err)
		return
	}
	views.ReportResult("GaussElimination", det, 0, &buffer)
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
	quest01Gauss()
}
