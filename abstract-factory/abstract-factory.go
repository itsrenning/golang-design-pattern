package main

import (
	"fmt"
	"strings"
)

type IRestaurantFactory interface {
	CookSoup() ISoup
	CookNoodle() INoodle
}

type ISoup interface {
	getName() string
}

type INoodle interface {
	getName() string
}

type Soup struct {
	name string
}

func (s *Soup) getName() string {
	return s.name
}

type Noodle struct {
	name string
}

func (n *Noodle) getName() string {
	return n.name
}

type ChineseSoup struct {
	Soup
}

type ChineseNoodle struct {
	Noodle
}

type JapaneseSoup struct {
	Soup
}

type JapaneseNoodle struct {
	Noodle
}

type Chinese struct{}

func (c *Chinese) CookSoup() ISoup {
	return &ChineseSoup{
		Soup{
			name: "Chinese Soup",
		},
	}
}

func (c *Chinese) CookNoodle() INoodle {
	return &ChineseNoodle{
		Noodle{
			name: "Chinese Noodle",
		},
	}
}

type Japanese struct{}

func (j *Japanese) CookSoup() ISoup {
	return &JapaneseSoup{
		Soup{
			name: "Japanese Soup",
		},
	}
}

func (j *Japanese) CookNoodle() INoodle {
	return &JapaneseNoodle{
		Noodle{
			name: "Japanese Noodle",
		},
	}
}

type Restaurant string

const (
	ChineseRestaurant  Restaurant = "Chinese Restaurant"
	JapaneseRestaurant Restaurant = "Japanese Restaurant"
)

func GetRestaurantFactory(r Restaurant) (IRestaurantFactory, error) {
	if r == ChineseRestaurant {
		return &Chinese{}, nil
	}
	if r == JapaneseRestaurant {
		return &Japanese{}, nil
	}
	return nil, fmt.Errorf("no such type of food")
}

func PrintSoup(s ISoup) {
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println(s.getName())
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println()
}

func PrintNoodle(s INoodle) {
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println(s.getName())
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println()
}
