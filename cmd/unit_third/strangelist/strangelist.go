package main

import (
	"fmt"
)

func main() {
	// Dados fornecidos
	data := []float64{
		0.90, 1.42, 1.30, 1.55, 1.63,
		1.32, 1.35, 1.47, 1.95, 1.66,
		1.96, 1.47, 1.92, 1.35, 1.05,
		1.85, 1.74, 1.65, 1.78, 1.71,
		2.29, 1.82, 2.06, 2.14, 1.27,
	}

	// Calcular estatísticas descritivas
	mean := calculateMean(data)
	median := calculateMedian(data)
	mode := calculateMode(data)
	rangeVal := calculateRange(data)
	stdDev := calculateStandardDeviation(data)
	variance := calculateVariance(data)
	coefficientOfVariation := calculateCoefficientOfVariation(stdDev, mean)

	// Imprimir resultados
	fmt.Println("Mean:", mean)
	fmt.Println("Median:", median)
	fmt.Println("Mode:", mode)
	fmt.Println("Range:", rangeVal)
	fmt.Println("Standard deviation:", stdDev)
	fmt.Println("Variance:", variance)
	fmt.Printf("Coefficient of variation: %.6f%%\n", coefficientOfVariation*100)
}
