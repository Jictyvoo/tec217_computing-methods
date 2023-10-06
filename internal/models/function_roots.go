package models

type (
	RootFunctionCallback[T Numeric] func(x T) T
	InteractionData[T Numeric]      struct {
		Interaction    uint64
		InputValues    []T
		RelativeError  T
		ResultValues   []T
		FunctionResult T
	}
)
