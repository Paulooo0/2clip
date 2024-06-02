package util

import (
	"fmt"
	"os"

	"github.com/boltdb/bolt"
)

func Authenticate(tx *bolt.Tx) error {
	condition := true
	for condition {
		var password string
		fmt.Print("\nEnter your password: ")
		fmt.Print("\033[8m")
		fmt.Scanln(&password)
		fmt.Print("\033[28m")

		authenticated := ValidatePassword(tx, password)

		condition = isAutheticationFailed(authenticated)
	}

	return nil
}

func isAutheticationFailed(authenticated bool) bool {
	if !authenticated {
		fmt.Println("\nAuthentication failed!")
		fmt.Println("TIP: if you don't have a password, run '2clip auth' to create one")

		answerCondition := true
		for answerCondition {
			fmt.Print("Do you want to try again? [Y/N]: ")
			var answer string
			fmt.Scanln(&answer)
			if answer == "N" || answer == "n" {
				os.Exit(0)
			} else if answer == "Y" || answer == "y" {
				answerCondition = false
			} else {
				fmt.Println("\nInvalid answer, please type Y or N")
				answerCondition = true
			}
		}
		return true
	} else {
		fmt.Println("\nAuthentication successful!")
		return false
	}
}
