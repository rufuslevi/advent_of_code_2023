package main

import (
	"advent_of_code_2023/jour1"
	"advent_of_code_2023/utils"
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Fprintln(os.Stderr, "Please enter one day selection")
		return
	}

	day, err := strconv.Atoi(args[1])
	utils.CheckErr(err, "Please enter a number for the day selection")

	switch day {
	case 1:
		answer, err := jour1.Main()
		utils.CheckErr(err, "")

		fmt.Printf("The answer is : %d \n", answer)
	}
}
