package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	var index int8 = 15

	var bigIndex int32 = int32(index)

	var bigNumber int64 = 64
	var littleNumber int8
	littleNumber = int8(bigNumber)

	var x int64 = 57
	var y float64 = float64(x)

	fmt.Printf("index data type -> %T\n", index)
	fmt.Printf("bigIndex data type -> %T\n", bigIndex)
	fmt.Println("little number -> ", littleNumber)
	fmt.Printf("Int to float conversion %.2f\n", y)

	intToString := strconv.Itoa(10)
	fmt.Printf("%T\n", intToString)
	fmt.Printf("%q\n", intToString)

	user := "Sammy"
	lines := 50

	fmt.Println("Congratulations, " + user + "! You just wrote " + strconv.Itoa(lines) + " lines of code.")

	//floats to string
	fmt.Println(fmt.Sprint(421.034))

	f := 5524.53
	fmt.Println(fmt.Sprint(f))
	fmt.Println("Sammy has " + fmt.Sprint(f) + " points.")

	// strings to integer
	lines_yesterday := "50"
	lines_today := "108"

	yesterday, err := strconv.Atoi(lines_yesterday)
	if err != nil {
		log.Fatal(err)
	}

	today, err := strconv.Atoi(lines_today)
	if err != nil {
		log.Fatal(err)
	}
	lines_more := today - yesterday

	fmt.Println("lines more", lines_more)

	//strings and bytes
	stringA := "my string"
	sliceB := []byte(stringA)
	stringC := string(sliceB)

	fmt.Println("String A -> ", stringA)
	fmt.Println("Slice B -> ", sliceB)
	fmt.Println("String C -> ", stringC)
}
