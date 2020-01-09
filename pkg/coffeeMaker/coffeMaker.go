package coffeeMaker

import (
	"fmt"

	models "github.com/DmitriyBelousov/voice-manager/pkg/models/coffeeMaker"
)

type coffeeMaker interface {
	SetWater()
	SetCoffee()
	SetAdditives()
	GetProduct() models.Coffee
}

// CoffeeMachine build director
type CoffeeMachine interface {
	SetBuilder(coffeeMaker)
	Make() models.Coffee
	Taste()
}

type coffeeMachine struct {
	maker coffeeMaker
}

// SetBuilder ...
func (d *coffeeMachine) SetBuilder(cm coffeeMaker) {
	d.maker = cm
}

// Make build composition of concrete product
func (d *coffeeMachine) Make() models.Coffee {
	d.maker.SetWater()
	d.maker.SetCoffee()
	d.maker.SetAdditives()
	return d.maker.GetProduct()
}

// Taste check composition of product
func (d *coffeeMachine) Taste() {
	coffee := d.maker.GetProduct()
	fmt.Println(coffee.Composition.String())
}

// NewCoffeeMachine ...
func NewCoffeeMachine(cm coffeeMaker) CoffeeMachine {
	return &coffeeMachine{
		maker: cm,
	}
}
