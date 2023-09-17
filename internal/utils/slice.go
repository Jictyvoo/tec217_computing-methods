package utils

func DeepCopyBidimensionalSlice[T any, M ~[][]T](input M) (destination M) {
	destination = make(M, len(input))
	for index, equation := range input {
		destination[index] = make([]T, len(equation))
		copy(destination[index], equation)
	}

	return destination
}

func CloneSlice[T any, S ~[]T](input S) (destination S) {
	destination = make(S, len(input))
	copy(destination, input)
	return destination
}
