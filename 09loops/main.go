package main

import "fmt"

func fizzball() {
	// normal loop
	for i := 0; i < 20; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizz bizz")
		} else if i%5 == 0 {
			fmt.Println("bizz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else {
			fmt.Println(i)
		}
	}

	// infinite loop
	// for i:= 0; ; i++{
	// }

	//while loop
	// for statement{
	// }
}

func main() {
	fizzball()
}
