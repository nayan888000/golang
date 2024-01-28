package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// welcome := "welcome"

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for the input,", input)

}
