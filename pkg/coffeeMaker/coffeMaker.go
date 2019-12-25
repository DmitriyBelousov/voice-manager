package coffeeMaker

import (
	"fmt"
	"strings"
)

type Coffee struct {
	Composition strings.Builder
}

type CoffeeMaker interface {
	SetWater()
	SetCoffee()
	SetAdditives()
	GetProduct() Coffee
}

type CoffeeMachine struct {
	maker CoffeeMaker
}

func (d *CoffeeMachine) SetBuilder(cm CoffeeMaker) {
	d.maker = cm
}

func (d *CoffeeMachine) Make() Coffee {
	d.maker.SetWater()
	d.maker.SetCoffee()
	d.maker.SetAdditives()
	return d.maker.GetProduct()
}

func (d *CoffeeMachine) Taste() {
	coffee := d.maker.GetProduct()
	fmt.Println(coffee.Composition.String())
}
