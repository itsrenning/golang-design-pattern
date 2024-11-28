package main

import (
	"fmt"
	"strings"
)

type IMobile interface {
	getName() string
	getCPU() int
	getRAM() int
}

type Mobile struct {
	name string
	cpu  int
	ram  int
}

func (m *Mobile) getName() string {
	return m.name
}

func (m *Mobile) getCPU() int {
	return m.cpu
}

func (m *Mobile) getRAM() int {
	return m.ram
}

type SamsungS12 struct {
	Mobile
}

func Print(m IMobile) {
	fmt.Println(strings.Repeat("=", 50))
	fmt.Printf("name: %s, cpu: %d, ram: %d\n", m.getName(), m.getCPU(), m.getRAM())
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println()
}

type IPhone16Pro struct {
	Mobile
}

func (m *IPhone16Pro) Print() {
	fmt.Println(strings.Repeat("=", 50))
	fmt.Printf("name: %s, cpu: %d, ram: %d\n", m.name, m.cpu, m.ram)
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println()
}

type MobileType string

const (
	IPhone  MobileType = "IPhone16Pro"
	Samsung MobileType = "SamsungS12"
)

func GetMobile(mt MobileType) IMobile {
	if mt == IPhone {
		return &IPhone16Pro{
			Mobile{
				name: "IPhone 16 Pro",
				cpu:  4,
				ram:  16,
			},
		}
	}
	if mt == Samsung {
		return &SamsungS12{
			Mobile{
				name: "Samsung S12",
				cpu:  8,
				ram:  32,
			},
		}
	}
	return nil
}
