package visitor

import (
	"fmt"
)

type drinkVisitor struct {
	GivenDrink []string
}

func (this *drinkVisitor) visitBusinessPassenger(bp *businessPassenger) {
	this.GivenDrink = append(this.GivenDrink, fmt.Sprintf("visit pass %s and give wine \n", bp.Name))
}

func (this *drinkVisitor) visitEconomPassenger(ep *economPassenger) {
	this.GivenDrink = append(this.GivenDrink, fmt.Sprintf("visit pass %s and give water \n", ep.Name))
}

// NewDrinkVisitor create new drink visitor
func NewDrinkVisitor() *drinkVisitor {
	return &drinkVisitor{}
}
