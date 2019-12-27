package main

import (
	"github.com/DmitriyBelousov/voice-manager/pkg/coffee"
	"github.com/DmitriyBelousov/voice-manager/pkg/coffeeMaker"
	"github.com/DmitriyBelousov/voice-manager/pkg/service/coffeeService"
)

func main() {

	cm := coffeeService.NewCoffeeMaker()
	machine := coffeeMaker.NewCoffeeMachine(cm)

	machine.SetBuilder(new(coffee.Americano))
	machine.Make()
	machine.Taste()

	machine.SetBuilder(new(coffee.Capuchino))
	machine.Make()
	machine.Taste()
}
