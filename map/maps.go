package main

import "fmt"

func main() {

	//map[type_of_key]type_of_value{value}
	sammy := map[string]string{"name": "Sammy", "animal": "shark", "color": "blue", "location": "ocean"}

	fmt.Println("map ->", sammy)
	fmt.Println("Sammy name ->", sammy["name"])
	fmt.Println("Sammy animal ->", sammy["animal"])
}
