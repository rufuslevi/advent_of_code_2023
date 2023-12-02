package utils

import (
	"fmt"
	"os"
)

func CheckErr(err error, mess string) {
	if err != nil {
		if mess == "" {
			fmt.Fprintln(os.Stderr, err.Error())
		} else {
			fmt.Fprintln(os.Stderr, mess)
		}
	}
}
