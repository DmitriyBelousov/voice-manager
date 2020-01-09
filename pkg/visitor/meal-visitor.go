package visitor

import (
	"fmt"
)

type mealVisitor struct {
	Meal []string
}

func (this *mealVisitor) visitBusinessPassenger(bp *businessPassenger) {
	this.Meal = append(this.Meal, fmt.Sprintf("visit pass %s and give meat \n", bp.Name))
}

func (this *mealVisitor) visitEconomPassenger(ep *economPassenger) {
	this.Meal = append(this.Meal, fmt.Sprintf("visit pass %s and give bread \n", ep.Name))
}

// NewMealVisitor create new meal visitor
func NewMealVisitor() *mealVisitor {
	return &mealVisitor{}
}
