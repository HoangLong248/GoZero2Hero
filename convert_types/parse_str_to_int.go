package main

import (
	"fmt"
	"strconv"
)

func main() {
	// var no int = 100
	// fmt.Println(reflect.TypeOf(no))

	// var intStr string = "100"
	// fourBaseEightBitInt, _ := strconv.ParseInt(intStr, 4, 8)    // becomes no 16 and int64
	// tenBaseSixTeenBitInt, _ := strconv.ParseInt(intStr, 10, 16) // no 100, and int64
	// fmt.Println(reflect.TypeOf(fourBaseEightBitInt))
	// fmt.Println(reflect.TypeOf(tenBaseSixTeenBitInt))

	var noOfPlayers = 8
	fmt.Printf("%T\n", noOfPlayers)
	str := strconv.Itoa(noOfPlayers)
	fmt.Printf("%T", str)
}
