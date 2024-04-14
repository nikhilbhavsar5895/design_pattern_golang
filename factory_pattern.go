
// Factory Design Pattern in Go:
// Why
// What
- Creational design pattern
- To hide the creation logic of the instances
- Client interacts with factory and tells what kind  of instances need to be created 
- then the factory interact with concrete object and returns corresponding instance back
// How
package main

import "fmt"

// App
type Car interface {
	getCar() string
}

type SedanType struct {
	Name string
}

func getNewSedan() *SedanType {
	return &SedanType{}
}

func (s *SedanType) getCar() string {
	return "Honda City"
}

type HatchBackType struct {
	Name string
}

func getNewHatchBack() *HatchBackType {
	return &HatchBackType{}
}

func (h *HatchBackType) getCar() string {
	return "Polo"
}

func main() { // Client
	getCarFactory(1)
	getCarFactory(2)
}

func getCarFactory(cartype int) { // Factory Class/Object
	var car Car
	if cartype == 1 {
		car = getNewHatchBack()
	} else {
		car = getNewSedan()
	}
	carInfo := car.getCar()
	fmt.Println(carInfo)
}
