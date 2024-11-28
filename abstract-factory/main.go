package main

import "fmt"

func main() {
	cr, cErr := GetRestaurantFactory(ChineseRestaurant)
	if cErr != nil {
		fmt.Println(cErr)
	}
	cSoup := cr.CookSoup()
	PrintSoup(cSoup)
	cNoodle := cr.CookNoodle()
	PrintNoodle(cNoodle)

	jr, jErr := GetRestaurantFactory(JapaneseRestaurant)
	if jErr != nil {
		fmt.Println(jErr)
	}
	jSoup := jr.CookSoup()
	PrintSoup(jSoup)
	jNoodle := jr.CookNoodle()
	PrintNoodle(jNoodle)
}
