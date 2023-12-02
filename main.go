package main

import (
	"advent_of_code_2023/jour1"
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) > 2 {
		println("Do not enter more than one day.")
		return
	}
	day, err := strconv.Atoi(args[1])
	if err != nil {
		println("Please enter a number")
		return
	}

	switch day {
	case 1:
		answer, err := jour1.Main()
		if err != nil {
			println(err.Error())
		}

		fmt.Printf("The answer is : %d \n", answer)
	}
}
