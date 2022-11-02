package main

import "fmt"

func main() {
	testScoreGrade5 := 80
	testScoreGrade4 := 60
	testScoreGrade3 := 50
	testScore := 90

	hasGas := true
	hasKeyIgnition := true

	hasBuger := true
	hasSandwiches := false

	printMessage := true

	if printMessage {
		fmt.Println("Message")
	}

	if testScore < 0 {
		fmt.Println("Score can't negative")
	} else if testScore >= testScoreGrade5 {
		fmt.Println("Top mark")
	} else if testScore >= testScoreGrade4 {
		fmt.Println("Pass with distinction")
	} else if testScore >= testScoreGrade3 {
		fmt.Println("Pass with distinction")
	} else {
		fmt.Println("Failed")
	}

	if hasGas && hasKeyIgnition {
		fmt.Println("Can drive car")
	}

	if hasBuger || hasSandwiches {
		fmt.Println("Can eat")
	}

	if !hasSandwiches {
		fmt.Println("No sandwiches, then I will starve, I only eat sandwiches")
	}

}
