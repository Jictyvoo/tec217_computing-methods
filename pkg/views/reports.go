package views

import (
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/jictyvoo/tec217_computing-methods/internal/utils/colors"
)

func ReportError(err error) {
	slog.Error("Failed to find root", slog.String("error", err.Error()))
}

func ReportResult(
	methodName string, outTable fmt.Stringer,
	totalIterations uint32, result ...float64,
) {
	if methodName != "" {
		fmt.Println(
			colors.Purple(methodName) +
				" - " +
				colors.Cyan(fmt.Sprintf("%d Iterations", totalIterations)),
		)
	}

	var logMessage string
	if len(result) == 1 {
		logMessage = fmt.Sprintf(
			"The value of root is : %.5f. After %d iteractions",
			result, totalIterations,
		)
	} else {
		msgBytes, _ := json.Marshal(result)
		logMessage = string(msgBytes)
	}

	var tableStr string
	if outTable != nil {
		tableStr = outTable.String()
	}
	slog.Info(
		logMessage,
		slog.String("method-name", methodName),
		slog.String("output-table", tableStr),
	)
}

func ReportLinearSystemResult(
	methodName string,
	det float64,
	foundRoots []float64,
	outTable fmt.Stringer,
) {
	if methodName != "" {
		fmt.Println(colors.Purple(methodName))
	}

	var tableStr string
	if outTable != nil {
		tableStr = outTable.String()
	}
	slog.Info(
		fmt.Sprintf(
			"Det: %.2f. Found roots: %v", det, foundRoots,
		),
		slog.String("method-name", methodName),
		slog.String("output-table", tableStr),
	)
}

func ReportMatrixTable(tableName string, outputTable fmt.Stringer) {
	fmt.Println(colors.Cyan(tableName))
	slog.Info(outputTable.String())
}
