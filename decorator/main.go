package main

import "fmt"

func main() {
	b := &Basic{}
	cwt := &CPUWithTwoCores{b}
	ram16G := &RAM16G{cwt}
	ssd500 := &SSD500G{ram16G}
	total := VideoCardGTX4090{ssd500}
	fmt.Printf("Total cost: %d\n", total.Cost())
}
