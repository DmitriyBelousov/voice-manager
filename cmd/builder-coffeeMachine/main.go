package main

import (
	"github.com/DmitriyBelousov/voice-manager/pkg/coffee"
	"github.com/DmitriyBelousov/voice-manager/pkg/coffeeMaker"
)

func main() {

	machine := coffeeMaker.CoffeeMachine{}
	amer := new(coffee.Americano)
	machine.SetBuilder(amer)
	machine.Make()
	machine.Taste()

	capuchino := new(coffee.Capuchino)
	machine.SetBuilder(capuchino)
	machine.Make()
	machine.Taste()
}
