package main

import (
	"fmt"
	"math"

	"github.com/jictyvoo/tec217_computing-methods/internal/utils"
)

func performAnalyticsCalculations(a [2]float64, newY []float64) {
	// Calculate the fit line
	xFit := make([]float64, 100)
	yFit := make([]float64, 100)
	for index := range xFit {
		xFit[index] = float64(index + 1)
		yFit[index] = a[0]*xFit[index] + a[1]
	}

	// Calculate total standard deviation
	totalY := utils.ReduceSlice(newY, func(old, a float64) float64 { return old + a })
	meanY := totalY / float64(len(newY))

	tempStdTotal := utils.ReduceSlice(
		newY, func(old, a float64) float64 { diff := a - meanY; return old + (diff * diff) },
	)
	stdTotal := math.Sqrt(tempStdTotal / float64(len(newY)-1))

	// Calculate standard error of the estimate
	var sumSquaredError float64
	for index, y := range newY {
		diff := y - yFit[index]
		sumSquaredError += diff * diff
	}
	stdError := math.Sqrt(sumSquaredError / float64(len(newY)-2))

	// Statistical analysis
	fmt.Printf("Total Standard Deviation: %.4f\n", stdTotal)
	fmt.Printf("Standard error of the Estimate: %.4f\n", stdError)
}
