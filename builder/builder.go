package main

import "fmt"

type IPhoneBuilder interface {
	SetCPU(cpu int)
	SetRAM(ram int)
	Build() *Phone
}

type Director struct {
	PhoneBuilder IPhoneBuilder
}

func NewDirector(pb IPhoneBuilder) *Director {
	return &Director{PhoneBuilder: pb}
}

func (d *Director) SetBuilder(b IPhoneBuilder) {
	d.PhoneBuilder = b
}

func (d *Director) BuildPhone() *Phone {
	return d.PhoneBuilder.Build()
}

type Phone struct {
	CPU int
	RAM int
}

func (p *Phone) String() {
	fmt.Printf("Phone Specs: [%d]Cores CPU & [%d]G RAM\n", p.CPU, p.RAM)
}

type StandardPhoneBuilder struct {
	Phone
}

func NewStandardPhoneBuilder() *StandardPhoneBuilder {
	return &StandardPhoneBuilder{}
}

func (spb *StandardPhoneBuilder) SetCPU(cpu int) {
	spb.CPU = cpu
}

func (spb *StandardPhoneBuilder) SetRAM(ram int) {
	spb.RAM = ram
}

func (spb *StandardPhoneBuilder) Build() *Phone {
	spb.SetCPU(2)
	spb.SetRAM(16)
	return &spb.Phone
}

type PremiumPhoneBuilder struct {
	Phone
}

func NewPremiumPhoneBuilder() *PremiumPhoneBuilder {
	return &PremiumPhoneBuilder{}
}

func (pb *PremiumPhoneBuilder) SetCPU(cpu int) {
	pb.CPU = cpu
}

func (pb *PremiumPhoneBuilder) SetRAM(ram int) {
	pb.RAM = ram
}

func (pb *PremiumPhoneBuilder) Build() *Phone {
	pb.SetCPU(4)
	pb.SetRAM(32)
	return &pb.Phone
}
