package main

import "fmt"

func main() {
	a := "Hello, Go!"
	for i, c := range a {
		fmt.Printf("%d: %s\n", i, string(c))
	}
	fmt.Println("length of 'Hello, Go!': ", len(a))

	s := "Sammy likes declaring strings."
	fmt.Println(s)
}
