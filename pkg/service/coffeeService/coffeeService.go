package coffeeService

import (
	coffeeMaker2 "github.com/DmitriyBelousov/voice-manager/pkg/models/coffeeMaker"
)

type coffeeMaker interface {
	SetWater()
	SetCoffee()
	SetAdditives()
	GetProduct() coffeeMaker2.Coffee
}

type CoffeeMaker interface {
	SetWater()
	SetCoffee()
	SetAdditives()
	GetProduct() coffeeMaker2.Coffee
}

type coffeeService struct {
	maker coffeeMaker
}

func (cs *coffeeService) SetWater() {

}

func (cs *coffeeService) SetCoffee() {

}

func (cs *coffeeService) SetAdditives() {

}

func (cs *coffeeService) GetProduct() coffeeMaker2.Coffee {
	return cs.maker.GetProduct()
}

func NewCoffeeMaker() CoffeeMaker {
	return &coffeeService{}
}
