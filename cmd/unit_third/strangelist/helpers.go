package main

import (
	"math"
	"slices"

	"github.com/jictyvoo/tec217_computing-methods/internal/models"
)

// Função para calcular a média
func calculateMean[T models.Numeric](data []T) T {
	var sum T = 0
	for _, value := range data {
		sum += value
	}
	return sum / T(len(data))
}

// Função para calcular a mediana
func calculateMedian[T models.Numeric](data []T) T {
	slices.Sort(data)
	n := len(data)
	if n%2 == 0 {
		return (data[n/2-1] + data[n/2]) / 2
	}
	return data[n/2]
}

// Função para calcular a moda
func calculateMode[T models.Numeric](data []T) T {
	frequencies := make(map[T]int)
	for _, value := range data {
		frequencies[value]++
	}
	maxFrequency := 0
	var mode T
	for value, frequency := range frequencies {
		if frequency > maxFrequency {
			mode = value
			maxFrequency = frequency
		}
	}
	return mode
}

// Função para calcular a amplitude
func calculateRange[T models.Numeric](data []T) T {
	var results struct {
		min, max T
	}
	results.min = data[0]
	results.max = data[0]
	for _, value := range data {
		if value < results.min {
			results.min = value
		}
		if value > results.max {
			results.max = value
		}
	}
	return results.max - results.min
}

// Função para calcular o desvio-padrão
func calculateStandardDeviation[T models.Numeric](data []T) T {
	mean := calculateMean(data)
	var sumOfSquaredDifferences T
	for _, value := range data {
		diff := value - mean
		sumOfSquaredDifferences += diff * diff
	}
	variance := sumOfSquaredDifferences / T(len(data)-1)
	return T(math.Sqrt(float64(variance)))
}

// Função para calcular a variância
func calculateVariance[T models.Numeric](data []T) T {
	mean := calculateMean(data)
	var sumOfSquaredDifferences T
	for _, value := range data {
		diff := value - mean
		sumOfSquaredDifferences += diff * diff
	}
	return sumOfSquaredDifferences / T(len(data)-1)
}

// Função para calcular o coeficiente de variação
func calculateCoefficientOfVariation[T models.Numeric](stdDev, mean T) T {
	return stdDev / mean
}
