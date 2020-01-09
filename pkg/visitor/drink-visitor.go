package visitor

import (
	"fmt"
)

type drinkVisitor struct {
	Drink []string
}

func (this *drinkVisitor) visitBusinessPassenger(bp *businessPassenger) {
	this.Drink = append(this.Drink, fmt.Sprintf("visit pass %s and give wine \n", bp.Name))
}

func (this *drinkVisitor) visitEconomPassenger(ep *economPassenger) {
	this.Drink = append(this.Drink, fmt.Sprintf("visit pass %s and give water \n", ep.Name))
}

// NewDrinkVisitor create new drink visitor
func NewDrinkVisitor() *drinkVisitor {
	return &drinkVisitor{}
}
