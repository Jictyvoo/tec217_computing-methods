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

// ReduceSlice iterates over a slice and applies a callback function to reduce it to a single value.
// The callback function takes as input the current reduced value and the current element of the slice,
// and returns the updated reduced value. The callback function should have the same type as the elements
// of the slice and the return type of the function.
func ReduceSlice[T any, S ~[]T](input S, callback func(old, a T) T) (result T) {
	for _, element := range input {
		result = callback(result, element)
	}
	return
}
