package main

import "fmt"

const u16bitmax float64 = 65535
const kmhMultiple float64 = 1.60934

type car struct {
	gasPedal uint16 // min 0 max 65535
	brakePedal uint16
	steeringWheel int16 // -32k to +32k
	topSpeedKmh float64
}

/* This is a method of the car struct. (value receiver)
   A value receiver essentially makes a copy of what is passed in. 
   Better for small structs, or read only operations.*/
func (c car) getKmh() float64 {
	return float64(c.gasPedal) * (c.topSpeedKmh/u16bitmax)
}

// Calling another method within this method. (value receiver)
func (c car) getMph() float64 {
	return c.getKmh() / kmhMultiple
}

/* Method for modifying instance of a struct. (pointer receiver)
   Typically pass by reference if we modify the receiver. */
func (c *car) setTopSpeed(newSpeed float64) {
	c.topSpeedKmh = newSpeed
}

// Not a car method, but a func.
func newCarSpeed(c car, speed float64) car {
	c.topSpeedKmh = speed
	return c
}

func main() {
	aCar := car{gasPedal: 22341,
				 brakePedal: 0,
				 steeringWheel: 12561,
				 topSpeedKmh: 225.0}

	fmt.Println(aCar.gasPedal)
	fmt.Println(aCar.getKmh())
	fmt.Println(aCar.getMph())
	aCar.setTopSpeed(4200)
	aCar = newCarSpeed(aCar, 999)
	fmt.Println(aCar.getKmh())
	fmt.Println(aCar.getMph())
}