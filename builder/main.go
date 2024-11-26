package main

import "fmt"

func main() {
	spb := NewStandardPhoneBuilder()
	ppb := NewPremiumPhoneBuilder()
	d := NewDirector(spb)

	fmt.Println("========================")
	fmt.Println("Building Standard Phone:")
	sp := d.BuildPhone()
	sp.String()
	fmt.Println("========================")

	fmt.Println("")

	fmt.Println("========================")
	fmt.Println("Building Premium Phone:")
	d.SetBuilder(ppb)
	pp := d.BuildPhone()
	pp.String()
	fmt.Println("========================")
}
