package methods

import "github.com/jictyvoo/tec217_computing-methods/internal/models"

type commonFuncZeroState[T models.Numeric] struct {
	finalResult  T
	interactions []models.InteractionData[T]
}

//goland:noinspection GoMixedReceiverTypes
func (mtd *commonFuncZeroState[T]) registerInteraction(
	inputs []T, totalIteration uint32, relativeError T,
	rootResult, result T,
) {
	mtd.interactions = append(mtd.interactions, models.InteractionData[T]{
		Interaction:    uint64(totalIteration),
		InputValues:    inputs,
		RelativeError:  relativeError,
		FunctionResult: rootResult,
		Value:          result,
	})
}

//goland:noinspection GoMixedReceiverTypes
func (mtd commonFuncZeroState[T]) InteractionData() []models.InteractionData[T] {
	return mtd.interactions
}

//goland:noinspection GoMixedReceiverTypes
func (mtd *commonFuncZeroState[T]) Reset() {
	mtd.interactions = mtd.interactions[:0]
}