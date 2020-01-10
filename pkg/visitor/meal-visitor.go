package visitor

import (
	"fmt"
)

type mealVisitor struct {
	GivenMeal []string
}

func (this *mealVisitor) visitBusinessPassenger(bp *businessPassenger) {
	this.GivenMeal = append(this.GivenMeal, fmt.Sprintf("visit pass %s and give meat \n", bp.Name))
}

func (this *mealVisitor) visitEconomPassenger(ep *economPassenger) {
	this.GivenMeal = append(this.GivenMeal, fmt.Sprintf("visit pass %s and give bread \n", ep.Name))
}

// NewMealVisitor create new meal visitor
func NewMealVisitor() *mealVisitor {
	return &mealVisitor{}
}
