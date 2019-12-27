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

type CoffeeMachine interface {
	SetBuilder(coffeeMaker)
	Make() models.Coffee
	Taste()
}

type coffeeMachine struct {
	maker coffeeMaker
}

func (d *coffeeMachine) SetBuilder(cm coffeeMaker) {
	d.maker = cm
}

func (d *coffeeMachine) Make() models.Coffee {
	d.maker.SetWater()
	d.maker.SetCoffee()
	d.maker.SetAdditives()
	return d.maker.GetProduct()
}

func (d *coffeeMachine) Taste() {
	coffee := d.maker.GetProduct()
	fmt.Println(coffee.Composition.String())
}

func NewCoffeeMachine(cm coffeeMaker) CoffeeMachine {
	return &coffeeMachine{
		maker: cm,
	}
}
