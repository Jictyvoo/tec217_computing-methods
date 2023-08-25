package methods

import "github.com/jictyvoo/tec217_computing-methods/internal/models"

type commonState[T models.Numeric] struct {
	finalResult  T
	interactions []models.InteractionData[T]
}

func (mtd commonState[T]) InteractionData() []models.InteractionData[T] {
	return mtd.interactions
}
