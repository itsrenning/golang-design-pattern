package main

import "fmt"

func main() {
	fmt.Println("=================================")
	fmt.Println("Create c1 with singleton instance")
	c1 := GetCarSingletonInstance()
	c1.name = "c1 car"
	c1.windows = 4
	c1.wheels = 4
	c1.String()
	fmt.Println("=================================")

	fmt.Println("")

	fmt.Println("=================================")
	fmt.Println("Create c2 with singleton instance")
	c2 := GetCarSingletonInstance()
	c2.String()
	fmt.Println("=================================")
}
