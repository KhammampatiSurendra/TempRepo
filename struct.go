package main

import (
	"fmt"
)

// Define the Engine struct
type Engine struct {
	Horsepower int
	Type       string
}

// Define the Car struct
type Car struct {
	Name           string
	Brand          string
	Type           string
	YearsInService int
	Engine         Engine
}

func main() {
	car := Car{
		Name:           "Model S",
		Brand:          "Tesla",
		Type:           "Electric",
		YearsInService: 3,
		Engine: Engine{
			Horsepower: 670,
			Type:       "Dual Motor",
		},
	}

	fmt.Println(car.Engine.Horsepower) // Output: 670
}
