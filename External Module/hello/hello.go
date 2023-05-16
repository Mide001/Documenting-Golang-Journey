package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Israel")
	fmt.Println(message)
}
