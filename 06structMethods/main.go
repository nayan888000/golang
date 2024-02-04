package main

import "fmt"

type rect struct {
	width  int
	length int
}

func (r rect) area() int {
	return r.length * r.width
}

var Ro rect = rect{
	width:  5,
	length: 25,
}

func main() {
	fmt.Println(Ro.area())
}
