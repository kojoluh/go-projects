package main

import (
	"io/ioutil"
	"encoding/json"
	"log"
	"fmt"
)

type contactInfo struct {
	Name string
	Email string
}

type purchaseInfo struct {
	Name string
	Price float32
	Amount int
}

type gasEngine struct {
	gallons float32
	mpg float32
}

type electricEngine struct {
	kwh float32
	mpkwh float32
}

type car [T gasEngine | electricEngine] struct {
	carMake string
	carModel string
	engine T
}

func makeCars() {
	var gasCar = car[gasEngine] {
		carMake: "Honda",
		carModel: "Civic",
		engine: gasEngine {
			gallons: 12.4,
			mpg: 40,
		},
	}

	var electricCar = car[electricEngine] {
		carMake: "Tesla",
		carModel: "Model 3",
		engine: electricEngine {
			kwh: 57.5,
			mpkwh: 4.17,
		},
	}
	fmt.Printf("\n\nCars:\n\tType: %T\n\tDetails: %v\n\n\tType: %T\n\tDetails: %v\n", gasCar, gasCar, electricCar, electricCar)
}

func loadJSON[T contactInfo | purchaseInfo](filePath string) []T {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
	 log.Print(err)
	}

	var loaded = []T{}
	json.Unmarshal(data, &loaded)

	return loaded
}


func main() {
	var contacts []contactInfo = loadJSON[contactInfo]("./contactInfo.json")
	fmt.Println("contacts: ", contacts)
	fmt.Println("purchaseInfo: ", loadJSON[purchaseInfo]("./purchaseInfo.json"))

	makeCars()
}

