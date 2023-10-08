package main

import "github.com/jictyvoo/tec217_computing-methods/internal/models"

// multiplyMatrix to multiply matrices
func multiplyMatrix[T models.Numeric](matrixA, matrixB [][]T, size int) [][]T {
	multipliedMatrix := make([][]T, size)
	for i := 0; i < size; i++ {
		multipliedMatrix[i] = make([]T, size)
		for j := 0; j < size; j++ {
			for k := 0; k < size; k++ {
				multipliedMatrix[i][j] += matrixA[i][k] * matrixB[k][j]
			}
		}
	}
	return multipliedMatrix
}
