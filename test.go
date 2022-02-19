package main

import (
	"fmt"
)

type Names string
func NewName() Names {
	var x Names = "testing"
	return x
}

type Car struct {
	Name	Names
}
func (c *Car) Load() {
	x := NewName()
	c.Name = x
}
var car = &Car{}

// var car *Car				-> not ok
// var car Car				-> ok
// var car = &Car{}			-> ok				==		var car = new(Car)


func main() {
	car.Load()
	fmt.Println("my car -> ", car)
}