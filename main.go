package main

import (
	"advent_of_code_2023/jour1"
	"advent_of_code_2023/jour2"
	"advent_of_code_2023/jour3"
	"advent_of_code_2023/utils"
	"fmt"
	"os"
	"strconv"
)

func printAnswer(mess string) {
	fmt.Printf("The answer is :\n%s\n", mess)
}

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
		{
			answer, err := jour1.Main()
			utils.CheckErr(err, "")

			printAnswer(answer)
		}
	case 2:
		{
			printAnswer(jour2.Main())
		}
	case 3:
		{
			printAnswer(jour3.Main())
		}
	}
}
