package main

import (
	"advent_of_code_2023/jour1"
	"advent_of_code_2023/jour2"
	"advent_of_code_2023/jour3"
	"advent_of_code_2023/jour4"
	"advent_of_code_2023/utils"
	"fmt"
	"os"
	"strconv"
	"time"
)

func printAnswer(mess string, startTime time.Time) {
	fmt.Printf("The answer is :\n%s\n", mess)
	fmt.Printf("It took %s to run today's solution\n", time.Since(startTime))
}

func mesureAnswers(iterations int, function func() string) {
	var smallestTime time.Duration

	for i := 0; i < iterations-1; i++ {
		startTime := time.Now()
		for j := 0; j < 1000; j++ {
			function()
		}
		newDuration := time.Since(startTime)
		if smallestTime > newDuration || smallestTime == 0 {
			smallestTime = newDuration
		}
	}

	fmt.Printf("The smallest time to run this day's solution is : %s", smallestTime)
}

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Fprintln(os.Stderr, "Please enter one day selection")
		return
	}

	day, err := strconv.Atoi(args[1])
	utils.CheckErr(err, "Please enter a number for the day selection")

	startTime := time.Now()
	switch day {
	case 1:
		{
			answer, err := jour1.Main()
			utils.CheckErr(err, "")

			printAnswer(answer, startTime)
		}
	case 2:
		{
			printAnswer(jour2.Main(), startTime)
		}
	case 3:
		{
			printAnswer(jour3.Main(), startTime)
			mesureAnswers(3, jour3.Main)
		}
	case 4:
		{
			printAnswer(jour4.Main(), startTime)
			mesureAnswers(3, jour4.Main)
		}
	}
}
