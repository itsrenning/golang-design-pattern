package main

type Computer interface {
	Cost() int
}

type Basic struct{}

func (b *Basic) Cost() int {
	return 10
}

type CPUWithTwoCores struct {
	Computer
}

func (c *CPUWithTwoCores) Cost() int {
	return c.Computer.Cost() + 10_000
}

type CPUWithFourCores struct {
	Computer
}

func (c *CPUWithFourCores) Cost() int {
	return c.Computer.Cost() + 20_000
}

type RAM16G struct {
	Computer
}

func (r *RAM16G) Cost() int {
	return r.Computer.Cost() + 3_000
}

type RAM32G struct {
	Computer
}

func (r *RAM32G) Cost() int {
	return r.Computer.Cost() + 6_000
}

type SSD128G struct {
	Computer
}

func (s *SSD128G) Cost() int {
	return s.Computer.Cost() + 1_500
}

type SSD500G struct {
	Computer
}

func (s *SSD500G) Cost() int {
	return s.Computer.Cost() + 2_500
}

type Power700W struct {
	Computer
}

func (p *Power700W) Cost() int {
	return p.Computer.Cost() + 1_000
}

type Power1200W struct {
	Computer
}

func (p *Power1200W) Cost() int {
	return p.Computer.Cost() + 3_000
}

type VideoCardGTX4050 struct {
	Computer
}

func (v *VideoCardGTX4050) Cost() int {
	return v.Computer.Cost() + 10_000
}

type VideoCardGTX4090 struct {
	Computer
}

func (v *VideoCardGTX4090) Cost() int {
	return v.Computer.Cost() + 20_000
}
