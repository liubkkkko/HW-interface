package main

import (
	"fmt"
)

type Dog struct {
	name            string  //ім'я
	feedConsumption float64 //розхід корму
	weight          float64 //вага
}

func (d Dog) consumption() float64 { //розрахунок кількості корму на місяць
	return d.feedConsumption * d.weight
}

func (d Dog) String() string {
	return fmt.Sprintf("A %s weighs %.2f kilograms", d.name, d.weight)
}

type Cat struct {
	name            string  //ім'я
	feedConsumption float64 //розхід корму
	weight          float64 //вага
}

func (c Cat) consumption() float64 { //розрахунок кількості корму на місяць
	return c.feedConsumption * c.weight
}

func (c Cat) String() string {
	return fmt.Sprintf("A %s weighs %.2f kilograms", c.name, c.weight)
}

type Cov struct {
	name            string  //ім'я
	feedConsumption float64 //розхід корму
	weight          float64 //вага
}

func (v Cov) consumption() float64 { //розрахунок кількості корму на місяць
	return v.feedConsumption * v.weight
}
func (v Cov) String() string {
	return fmt.Sprintf("A %s weighs %.2f kilograms", v.name, v.weight)
}

type Feeder interface {
	consumption() float64
}

type Constructer interface {
	Feeder
	fmt.Stringer
}

func PrintAll(c Constructer) {
	fmt.Printf("%s and consumes %.2f kg of feed per month\n\n", c.String(), c.consumption())
}

func main() {

	dogConsumption := 10 / 2 //розрахунок кількості корму на один кілограм для собак
	c := Cat{name: "murchuk", feedConsumption: 7, weight: 5}
	PrintAll(c)
	d := Dog{name: "barsik", feedConsumption: float64(dogConsumption), weight: 20}
	PrintAll(d)
	v := Cov{name: "munia", feedConsumption: 25, weight: 100}
	PrintAll(v)
	fmt.Println(c.consumption()+v.consumption()+d.consumption(), "kilograms of fodder are required per month")
}
