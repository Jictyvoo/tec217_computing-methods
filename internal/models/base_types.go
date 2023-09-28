package models

type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~uintptr | ~float32 | ~float64
}

type AlgebraicOperation rune

const (
	OpSum    AlgebraicOperation = '+'
	OpSub    AlgebraicOperation = '-'
	OpMult   AlgebraicOperation = '*'
	OpDiv    AlgebraicOperation = '/'
	OpPermut AlgebraicOperation = '`'
)
