package utils

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"io"
	"strconv"

	"github.com/jictyvoo/tec217_computing-methods/internal/models"
)

func WriteInteractionAsCSV(output io.Writer, data []models.InteractionData[float64]) (err error) {
	if len(data) <= 0 {
		err = errors.New("can't write a CSV without any data")
		return
	}
	writer := csv.NewWriter(output)
	defer writer.Flush()

	totalInputs := len(data[0].InputValues)

	headers := make([]string, 0, 4+totalInputs)
	valueInput := make([]string, 4+totalInputs)
	headers = append(headers, "Interaction")
	for i := 0; i < totalInputs; i++ {
		headers = append(headers, string([]rune{'a' + rune(i)}))
	}
	headers = append(headers, []string{"Found-X", "f(x)", "Relative-Error"}...)

	if err = writer.Write(headers); err != nil {
		return
	}

	var tempBytes []byte
	for _, value := range data {
		valueInput[0] = strconv.FormatUint(value.Interaction, 10)
		index := 1
		for count, inputValue := range value.InputValues {
			index += count
			valueInput[index] = strconv.FormatFloat(inputValue, 'f', 6, 64)
		}
		if tempBytes, err = json.Marshal(value.ResultValues); err != nil {
			return
		}
		valueInput[index+1] = string(tempBytes)
		valueInput[index+2] = strconv.FormatFloat(value.FunctionResult, 'f', 6, 64)
		valueInput[index+3] = strconv.FormatFloat(value.RelativeError, 'f', 6, 64)
		if err = writer.Write(valueInput); err != nil {
			return
		}
	}

	err = writer.Error()
	return
}
