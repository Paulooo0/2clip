package util

import (
	"fmt"
	"os"
)

func AnswerCondition() bool {
	answerCondition := true
	for answerCondition {
		answerCondition = getAnswer()
	}
	return true
}

func TryAgain() bool {
	fmt.Print("\nDo you want to try again? [Y/N]: ")
	answerCondition := true
	for answerCondition {
		answerCondition = getAnswer()
	}
	return true
}

func getAnswer() bool {
	var answer string
	fmt.Scanln(&answer)
	if answer == "N" || answer == "n" {
		os.Exit(0)
	} else if answer == "Y" || answer == "y" {
		return false
	} else {
		fmt.Print("\nInvalid answer, please type Y or N: ")
		return true
	}
	return true
}
