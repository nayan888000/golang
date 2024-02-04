package main

import "fmt"

func main() {
	squareFunc := selfMath(multiply)
	doubleFunc := selfMath(add)

	fmt.Println(squareFunc(5))
	fmt.Println(doubleFunc(5))

	addition := closures()

	addition(5)
	addition(6)
	addition(15)
	addition(50)
	fmt.Println(addition(32))
}

func closures() func(int) int {
	add := 0
	return func(num int) int {
		add += num
		return add
	}
}
func multiply(x, y int) int {
	return x * y
}

func add(x, y int) int {
	return x + y
}

func selfMath(mathfunc func(int, int) int) func(int) int {
	return func(x int) int {
		return mathfunc(x, x)
	}
}
