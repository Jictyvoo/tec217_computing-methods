package main

import (
	"fmt"
	"log/slog"

	"github.com/jictyvoo/tec217_computing-methods/internal/utils/colors"
)

func reportResult(
	methodName string, result float64,
	totalIterations uint32, outTable fmt.Stringer,
) {
	fmt.Println(
		colors.Purple(methodName) +
			" - " +
			colors.Cyan(fmt.Sprintf("%d Iterations", totalIterations)),
	)
	slog.Info(
		fmt.Sprintf(
			"The value of root is : %.5f. After %d iteractions",
			result, totalIterations,
		),
		slog.String("method-name", methodName),
		slog.String("output-table", outTable.String()),
	)
}

func reportError(err error) {
	slog.Error("Failed to find root", slog.String("error", err.Error()))
}
