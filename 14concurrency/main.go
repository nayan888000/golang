package main

import (
	"fmt"
	"time"
)

func sendEmail() {
	for i := 0; i < 6; i++ {
		go func() {
			fmt.Printf("Email received -%v\n", i)
		}()

		fmt.Printf("Email sent -%v\n", i)
	}
	time.Sleep(time.Second)
}

func main() {
	sendEmail()
}
