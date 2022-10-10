package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "Hello, Go!"
	for i, c := range a {
		fmt.Printf("%d: %s\n", i, string(c))
	}
	fmt.Println("length of 'Hello, Go!': ", len(a))

	s := "Sammy likes declaring strings."
	fmt.Println(s)

	fmt.Println("This string is in UPPERCASE ->", strings.ToUpper(s))
	fmt.Println("This other string is in lowercase ->", strings.ToLower(s))

	fmt.Println("This string starts with \"Sammy\"? ->", strings.HasPrefix(s, "Sammy"))
	fmt.Println("This string ends with \"strings\"? ->", strings.HasSuffix(s, "strings"))
	fmt.Println("This string contains the word \"likes\"? ->", strings.Contains(s, "likes"))
	fmt.Println("How many times the letter \"s\" appears in the sentece? ->", strings.Count(s, "s"))
	fmt.Println("The size of the this sentence is ->", len(s))

	//Join
	fmt.Println("Using Join ->", strings.Join([]string{"sharks", "crustaceans", "plankton"}, ", "))

	//Split
	text := "Sammy has a balloon."
	x := strings.Split(text, " ")
	fmt.Printf("Using Split -> %q\n", x)

	//Fields - ignore whitespaces
	data := "username password     email date"
	fields := strings.Fields(data)
	fmt.Printf("Using Fields -> %q\n", fields)

	//ReplaceAll
	fmt.Println("Replace \"has\" to \"had\" using \"ReplaceAll\"->", strings.ReplaceAll(text, "has", "had"))

}
