package views

import (
	"encoding/csv"
	"io"
	"strconv"

	"github.com/jictyvoo/tec217_computing-methods/internal/models"
)

func MatrixTableCSV[T models.Numeric](matrix [][]T, writer io.Writer) (err error) {
	csvWriter := csv.NewWriter(writer)
	defer csvWriter.Flush()

	for _, row := range matrix {
		stringRow := make([]string, len(row))
		for i, val := range row {
			stringRow[i] = strconv.FormatFloat(float64(val), 'f', 4, 64)
		}
		if err = csvWriter.Write(stringRow); err != nil {
			return
		}
	}
	return
}
