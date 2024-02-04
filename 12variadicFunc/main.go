package main

import "fmt"

func sum(nums ...int) int {
	num := 0
	for i := 0; i < len(nums); i++ {
		num += nums[i]
	}
	return num
}

func main() {
	total := sum(1, 3, 4, 5, 6, 7, 8)
	fmt.Println(total)
}
