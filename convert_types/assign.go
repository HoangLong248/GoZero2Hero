package main

import (
	"fmt"
	"os"
	"strconv"
)

func add(num1, num2 int) int {
	return num1 + num2
}

func main() {
	num1, err := strconv.Atoi(os.Args[1])
	num2, err2 := strconv.Atoi(os.Args[2])
	if err != nil || err2 != nil {
		panic("Arg can't be string")
	} else {
		sum := add(num1, num2)
		fmt.Println(sum)
	}

}
