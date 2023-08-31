package models

type (
	Numeric interface {
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
			~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
			~uintptr | ~float32 | ~float64
	}
	RootFunctionCallback[T Numeric] func(x T) T
	InteractionData[T Numeric]      struct {
		Interaction    uint64
		InputValues    []T
		RelativeError  T
		Value          T
		FunctionResult T
	}
)
