package clip

import (
	"fmt"

	"github.com/spf13/cobra"
)

var AuthCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate to the database",
	Long:  `Authenticate to the database.`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		authenticate()
	},
}

func authenticate() error {
	enterPassword()

	return nil
}

func enterPassword() string {
	var password string

	condition := true
	for condition {
		fmt.Println("Enter your password:")

		fmt.Scanln(&password)

		fmt.Println("Enter your password again:")

		var passwordAgain string
		fmt.Scanln(&passwordAgain)

		err := checkPassword(password, passwordAgain)

		if err != nil {
			condition = checkAnswer()
		} else {
			condition = false
		}
	}
	return password
}

func checkPassword(password1 string, password2 string) error {
	if password1 != password2 {
		return fmt.Errorf("passwords do not match")
	}
	return nil
}

func checkAnswer() bool {
	answerCondition := true
	for answerCondition {
		fmt.Println("You want to try again? [Y/N]")

		var answer string
		fmt.Scanln(&answer)
		if answer == "N" || answer == "n" {
			fmt.Println("Exiting...")
			answerCondition = false
			return false
		} else if answer == "Y" || answer == "y" {
			answerCondition = true
			return true
		} else {
			fmt.Println("Invalid answer, please type Y or N")
		}
		answerCondition = false
	}
	return false
}

