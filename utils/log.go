package utils

import "fmt"

func LogDayAnswers(inputLines []string, dayParts []func([]string) (string, error)) string {
	var answer string

	for i, part := range dayParts {
		partAnswer, err := part(inputLines)
		CheckErr(err, fmt.Sprintf("Error running part %d", i))

		answer = fmt.Sprintf("%s part %d : %s\n", answer, i, partAnswer)
	}

	return answer
}
