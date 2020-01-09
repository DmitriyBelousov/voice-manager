package main

import "fmt"

func main() {
	list := new(PassengersList)
	list.Add(NewEconomPassenger("Vasya"))
	list.Add(NewBusinessPassenger("Voldemar"))
	list.Add(NewEconomPassenger("Fedya"))

	foodStuart := new(MealVisitor)
	list.Accept(foodStuart)

	barmen := new(DrinkVisitor)
	list.Accept(barmen)

	fmt.Println(foodStuart.Meal)
	fmt.Println(barmen.Drink)
}

type Passenger interface {
	Accept(PassengerVisitor)
}

type BusinessPassenger struct {
	Name string
}

func (bp *BusinessPassenger) Accept(visitor PassengerVisitor) {
	visitor.visitBusinessPassenger(bp)
}

type EconomPassenger struct {
	Name string
}

func (ep *EconomPassenger) Accept(visitor PassengerVisitor) {
	visitor.visitEconomPassenger(ep)
}

type PassengersList struct {
	passengers []Passenger
}

func (pl *PassengersList) Add(pass Passenger) {
	pl.passengers = append(pl.passengers, pass)
}

func (pl *PassengersList) Accept(visitor PassengerVisitor) {
	for _, pass := range pl.passengers {
		pass.Accept(visitor)
	}
}

type PassengerVisitor interface {
	visitBusinessPassenger(bp *BusinessPassenger)
	visitEconomPassenger(ep *EconomPassenger)
}

func NewBusinessPassenger(name string) *BusinessPassenger {
	return &BusinessPassenger{Name: name}
}

func NewEconomPassenger(name string) *EconomPassenger {
	return &EconomPassenger{Name: name}
}

type MealVisitor struct {
	Meal []string
}

func (this *MealVisitor) visitBusinessPassenger(bp *BusinessPassenger) {
	this.Meal = append(this.Meal, fmt.Sprintf("visit pass %s and give meat \n", bp.Name))
}

func (this *MealVisitor) visitEconomPassenger(ep *EconomPassenger) {
	this.Meal = append(this.Meal, fmt.Sprintf("visit pass %s and give bread \n", ep.Name))
}

type DrinkVisitor struct {
	Drink []string
}

func (this *DrinkVisitor) visitBusinessPassenger(bp *BusinessPassenger) {
	this.Drink = append(this.Drink, fmt.Sprintf("visit pass %s and give wine \n", bp.Name))
}

func (this *DrinkVisitor) visitEconomPassenger(ep *EconomPassenger) {
	this.Drink = append(this.Drink, fmt.Sprintf("visit pass %s and give water \n", ep.Name))
}
