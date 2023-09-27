package models

type MatrixTransformationStep[T Numeric] struct {
	Matrix     [][]T              // The state of the matrix after this step
	Multiplier T                  // The multiplier used in this step
	RightRow   uint8              // The row that has been subtracted
	LeftRow    uint8              // The row that has been used for subtraction
	Operation  AlgebraicOperation // The row that defines which operation was made with
}

type RootCalculationStep[T Numeric] struct {
	Roots       []T // The roots after this step
	DividendSum T   // The sum of all elements to later be used to divide
	Divisor     T   // The value that will be divided
}
