package utils

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/jictyvoo/tec217_computing-methods/internal/models"
)

func formatLinearOpLabel(
	operation models.AlgebraicOperation,
	inputElement float64, leftRow, rightRow uint8,
) string {
	switch operation {
	case models.OpSum, models.OpSub:
		return fmt.Sprintf(
			"L%d = L%d %s (%.4f * L%d)",
			leftRow, leftRow, string(operation),
			inputElement, rightRow,
		)
	case models.OpMult, models.OpDiv:
		return fmt.Sprintf(
			"L%d = L%d %s %.4f",
			leftRow, leftRow, string(operation), inputElement,
		)
	case models.OpPermut:
		return fmt.Sprintf("L%d <=> L%d", leftRow, rightRow)
	case models.OpAttribution:
		return fmt.Sprintf("L[%d,%d] = %.4f", leftRow, rightRow, inputElement)
	}

	return "Failed to recognize operation"
}

func writeTriangulationSteps(
	writer *csv.Writer, triangulationSteps []models.MatrixTransformationStep[float64],
) (err error) {
	if len(triangulationSteps) <= 0 {
		err = errors.New("can't write a CSV without any data")
		return
	}

	// Write column headers for triangulation steps
	triangulationHeaders := []string{
		"Step", "Matrix", "Multiplier", "Operation",
	}
	if err = writer.Write(triangulationHeaders); err != nil {
		return
	}

	// Write each triangulation step
	for i, step := range triangulationSteps {
		matrixBytes, _ := json.Marshal(step.Matrix)
		record := []string{
			strconv.Itoa(i + 1),
			string(matrixBytes),
			strconv.FormatFloat(step.Multiplier, 'f', 6, 64),
			formatLinearOpLabel(
				step.Operation, step.Multiplier,
				step.LeftRow+1, step.RightRow+1,
			),
		}
		if err = writer.Write(record); err != nil {
			return
		}
	}

	return
}

func writeRootCalculationStep(
	writer *csv.Writer, rootCalculationSteps []models.RootCalculationStep[float64],
) (err error) {
	// Write column headers for root calculation steps
	rootCalculationHeaders := []string{"Step", "Roots", "DividendSum", "Divisor", "Added Root"}
	if err = writer.Write(rootCalculationHeaders); err != nil {
		return
	}

	// Write each root calculation step
	totalSteps := len(rootCalculationSteps) - 1
	for index, step := range rootCalculationSteps {
		rootsBytes, _ := json.Marshal(step.Roots)
		record := []string{
			strconv.Itoa(index + 1),
			string(rootsBytes),
			strconv.FormatFloat(step.DividendSum, 'f', 6, 64),
			strconv.FormatFloat(step.Divisor, 'f', 6, 64),
			strconv.FormatFloat(step.Roots[totalSteps-index], 'f', 6, 64),
		}
		if err = writer.Write(record); err != nil {
			return
		}
	}

	return
}

func WriteLinearInteractionAsCSV(
	output io.Writer,
	triangulationSteps []models.MatrixTransformationStep[float64],
	rootCalculationSteps []models.RootCalculationStep[float64],
) (err error) {
	writer := csv.NewWriter(output)
	defer writer.Flush()

	if err = writeTriangulationSteps(writer, triangulationSteps); err != nil {
		return
	}

	// Write a blank line before root calculation steps
	if err = writer.Write([]string{}); err != nil {
		return err
	}

	if err = writeRootCalculationStep(writer, rootCalculationSteps); err != nil {
		return
	}

	return writer.Error()
}
