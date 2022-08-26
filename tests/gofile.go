package main

import (
	"fmt"
	"strconv"
)

func main() {
	name := "Luís Carvalho"
	age, errr := strconv.Atoi("29")

	if errr != nil {
		fmt.Println("ups there was a problem converting your age")
	}

	fmt.Println("Porque Maria?, said, ", name, " and it is only this old, ", age)
}
