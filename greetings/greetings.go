package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Please enter your name.")
	var name string
	fmt.Scanln(&name)
	name = strings.TrimSpace(name)
	fmt.Printf("Hi, %s! I'm Go!\n", name)

	fmt.Println("What is your favorite color?")
	var color string
	fmt.Scanln(&color)
	fmt.Printf("Humm, your favorite color is %s\n", color)
}
