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

func calculated(animal []Constructer) float64 {
	var foodNeedle float64
	for _, animal:= range animal{
		foodNeedle += animal.consumption()	
		fmt.Printf("%s and consumes %.2f kg of feed per month\n\n", animal.String(), animal.consumption())	
	}	
	return foodNeedle
}

func main() {
	dogConsumption := 10 / 2 //розрахунок кількості корму на один кілограм для собак
	
	sAnimal := []Constructer{
		cat{name: "murchuk", weight: 2, feedConsumption: 7},
		cat{name: "kitsa", weight: 3, feedConsumption: 7},
		dog{name: "barsik", weight: 7, feedConsumption: float64(dogConsumption)},
		dog{name: "vulkan", weight: 10, feedConsumption: float64(dogConsumption)},
		cov{name: "munia", weight: 100, feedConsumption: 100},
	}
	fmt.Println(calculated(sAnimal), "kilograms of fodder are required per month")

}


