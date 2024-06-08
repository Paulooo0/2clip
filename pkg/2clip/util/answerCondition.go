package util

import (
	"fmt"
	"os"
)

func AnswerCondition() bool {
	answerCondition := true
	for answerCondition {
		fmt.Print("\nDo you want to try again? [Y/N]: ")
		var answer string
		fmt.Scanln(&answer)
		if answer == "N" || answer == "n" {
			os.Exit(0)
		} else if answer == "Y" || answer == "y" {
			answerCondition = false
		} else {
			fmt.Println("\nInvalid answer, please type Y or N")
		}
	}
	return true
}
