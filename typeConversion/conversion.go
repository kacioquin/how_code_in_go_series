package main

import "fmt"

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
}
