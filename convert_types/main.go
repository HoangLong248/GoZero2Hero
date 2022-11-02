package main

import (
	"fmt"
	"os"
	"strconv"
)

func add(first int, second int) int {
	return first + second
}

func main() {
	// no1, _ := strconv.Atoi(os.Args[1])
	// no2, _ := strconv.Atoi(os.Args[2])
	// var sum = add(no1, no2)
	// fmt.Println(sum)

	no, err := strconv.Atoi(os.Args[1])
	fmt.Println(no)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Couldn't covert: " + os.Args[1])
	} else {
		fmt.Println(no)
	}
}
