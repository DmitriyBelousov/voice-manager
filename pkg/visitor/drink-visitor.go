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

func NewDrinkVisitor() *drinkVisitor {
	return &drinkVisitor{}
}
