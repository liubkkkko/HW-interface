package main

import (
	"fmt"
)

type dog struct {
	name            string  //ім'я
	feedConsumption float64 //розхід корму
	weight          float64 //вага
}

func (d dog) consumption() float64 { //розрахунок кількості корму на місяць
	return d.feedConsumption * d.weight
}

func (d dog) String() string {
	return fmt.Sprintf("A %s weighs %.2f kilograms", d.name, d.weight)
}

type cat struct {
	name            string  //ім'я
	feedConsumption float64 //розхід корму
	weight          float64 //вага
}

func (c cat) consumption() float64 { //розрахунок кількості корму на місяць
	return c.feedConsumption * c.weight
}

func (c cat) String() string {
	return fmt.Sprintf("A %s weighs %.2f kilograms", c.name, c.weight)
}

type cov struct {
	name            string  //ім'я
	feedConsumption float64 //розхід корму
	weight          float64 //вага
}

func (v cov) consumption() float64 { //розрахунок кількості корму на місяць
	return v.feedConsumption * v.weight
}
func (v cov) String() string {
	return fmt.Sprintf("A %s weighs %.2f kilograms", v.name, v.weight)
}

type Constructer interface {
	consumption() float64
	fmt.Stringer
}

func PrintAll(c Constructer) {
	fmt.Printf("%s and consumes %.2f kg of feed per month\n\n", c.String(), c.consumption())
}

func main() {

	dogConsumption := 10 / 2 //розрахунок кількості корму на один кілограм для собак
	c := cat{name: "murchuk", feedConsumption: 7, weight: 5}
	PrintAll(c)
	d := dog{name: "barsik", feedConsumption: float64(dogConsumption), weight: 20}
	PrintAll(d)
	v := cov{name: "munia", feedConsumption: 25, weight: 100}
	PrintAll(v)
	fmt.Println(c.consumption()+v.consumption()+d.consumption(), "kilograms of fodder are required per month")
}
