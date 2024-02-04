package main

import (
	"fmt"
	"math"
	"reflect"
)

// Shape is an interface that defines the Area method.
type Shape interface {
	Area() float64
}

// Circle is a type representing a circle.
type Circle struct {
	Radius float64
}

// Rectangle is a type representing a rectangle.
type Rectangle struct {
	Width  float64
	Height float64
}

// Implementing the Area method for Circle.
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func (c Circle) Perimeter() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Implementing the Area method for Rectangle.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Function that takes any type implementing the Shape interface and calculates its area.
func calculateArea(s Shape) {
	area := s.Area()
	fmt.Printf("Area: %.2f\n", area)
	// this is an example of type assertion
	if circle, isCircle := s.(Circle); isCircle {
		fmt.Printf("type is%T\n", circle)
	} else {
		fmt.Println("type is", reflect.TypeOf(circle))
	}
}
func main() {
	// Creating instances of Circle and Rectangle.
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 3, Height: 4}
	// Calculating and printing the area for each shape.
	calculateArea(circle)    // Outputs: Area: 78.54
	calculateArea(rectangle) // Outputs: Area: 12.00
}
