package main

import "fmt"

type userError struct {
	name string
}

// below method helps above struct to implement Error interface
func (e userError) Error() string {
	return fmt.Sprintf("We can create custom error like this")
}

func main() {

}
