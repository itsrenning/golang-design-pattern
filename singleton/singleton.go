package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type Car struct {
	name    string
	wheels  int
	windows int
}

func (c *Car) String() {
	fmt.Printf("Car name: %s, Car windows: %d, Car wheels: %d\n", c.name, c.windows, c.wheels)
}

var CarSingletonInstance *Car

func GetCarSingletonInstance() *Car {
	if CarSingletonInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if CarSingletonInstance == nil {
			CarSingletonInstance = &Car{}
			return CarSingletonInstance
		}
		return CarSingletonInstance
	}
	return CarSingletonInstance
}
