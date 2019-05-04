package main

import "fmt"

type Speed float64

const (
	MPH Speed = 1
	KPH       = 1.60934
)

type Color string

const (
	BlueColor  Color = "blue"
	GreenColor       = "green"
	RedColor         = "red"
)

type Wheels string

const (
	SportsWheels Wheels = "sports"
	SteelWheels         = "steel"
)

type Sheets int

type Builder interface {
	Color(Color) Builder
	Wheels(Wheels) Builder
	TopSpeed(Speed) Builder
	Sheets(int) Builder
	Build() Interface
}

type Interface interface {
	Drive() error
	Stop() error
}

type carBuilder struct {
	color  Color
	wheels Wheels
	speed  Speed
	sheets int
}

type car struct {
	params carBuilder
}

type sportsWheelCar struct {
	params carBuilder
}

type steelWheelCar struct {
	params carBuilder
}

func NewBuilder() *carBuilder {
	return &carBuilder{
		color:  BlueColor,
		wheels: SportsWheels,
		speed:  MPH,
		sheets: 4,
	}
}

func (b *carBuilder) Color(color Color) Builder {
	b.color = color
	return b
}

func (b *carBuilder) Wheels(wheels Wheels) Builder {
	b.wheels = wheels
	return b
}

func (b *carBuilder) TopSpeed(speed Speed) Builder {
	b.speed = speed
	return b
}

func (b *carBuilder) Sheets(sheets int) Builder {
	b.sheets = sheets
	return b
}

func (b *carBuilder) Build() Interface {
	if b.wheels == SportsWheels {
		return &sportsWheelCar{
			params: *b,
		}
	}
	return &steelWheelCar{
		params: *b,
	}
}

func (c *car) Drive() error {
	fmt.Printf("Driving: %#+v\n", c.params)
	return nil
}

func (c *car) Stop() error {
	fmt.Printf("Stop: %#+v¥\n", c.params)
	return nil
}

func (spc *sportsWheelCar) Drive() error {
	fmt.Printf("Sports Driving: %#+v\n", spc.params)
	return nil
}

func (spc *sportsWheelCar) Stop() error {
	fmt.Printf("Sports Stop: %#+v¥\n", spc.params)
	return nil
}

func (stc *steelWheelCar) Drive() error {
	fmt.Printf("Steel Driving: %#+v\n", stc.params)
	return nil
}

func (stc *steelWheelCar) Stop() error {
	fmt.Printf("Steel Stop: %#+v¥\n", stc.params)
	return nil
}

func main() {
	car := NewBuilder().Color(BlueColor).Wheels(SportsWheels).TopSpeed(KPH).Sheets(2).Build()
	car.Drive()
	car.Stop()
}
