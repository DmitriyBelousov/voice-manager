package main

import (
	"fmt"

	"github.com/DmitriyBelousov/voice-manager/pkg/visitor"
)

func main() {
	list := new(visitor.PassengersList)
	list.Add(visitor.NewEconomPassenger("Vasya"))
	list.Add(visitor.NewBusinessPassenger("Voldemar"))
	list.Add(visitor.NewEconomPassenger("Fedya"))

	foodStuart := visitor.NewMealVisitor()
	list.Accept(foodStuart)

	barmen := visitor.NewDrinkVisitor()
	list.Accept(barmen)

	fmt.Println(foodStuart.GivenMeal)
	fmt.Println(barmen.GivenDrink)
}
