package utils

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/jictyvoo/tec217_computing-methods/internal/models"
)

func writeTriangulationSteps(
	writer *csv.Writer, triangulationSteps []models.TriangulationStep[float64],
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
		matrixStr := fmt.Sprintf("%v", step.Matrix)
		record := []string{
			strconv.Itoa(i + 1),
			matrixStr,
			strconv.FormatFloat(step.Multiplier, 'f', 6, 64),
			fmt.Sprintf(
				"L%d = L%d - (%.4f * L%d)",
				step.SubtractedRow+1, step.SubtractedRow+1,
				step.Multiplier, step.SubtractionRow+1,
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
		rootsStr := fmt.Sprintf("%v", step.Roots)
		record := []string{
			strconv.Itoa(index + 1),
			rootsStr,
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
	triangulationSteps []models.TriangulationStep[float64],
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
