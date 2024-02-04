package main

import "fmt"

func main() {

	// Array declaration
	var array1 [5]int

	// array with values
	// prime := [4]int{1, 2, 3, 4}

	for i := 0; i < len(array1); i++ {
		fmt.Println(array1[i])
	}

	// Slices declaration
	slice := []int{}
	// using make function
	// eg. it iis creating a int slice with 5 length and 10 capacity
	makeSlice := make([]int, 5, 10)
	fmt.Println(makeSlice)
	fmt.Println(slice)
}
