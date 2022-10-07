package main

import (
	"fmt"
	"strconv"
)

func main() {
	name := "Luís Carvalho"
	age, err := strconv.Atoi("29")
	if err != nil {
		fmt.Println("ups there was a problem converting your age")
	}

	fmt.Println("hello there guy")
	fmt.Println("Porque Maria?, said, ", name, " and it is only this old, ", age)
}
