package clip

import (
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
)

var UpdateAuthCmd = &cobra.Command{
	Use:     "update",
	Aliases: []string{"-u"},
	Short:   "Update the authentication",
	Long:    `Update the authentication.`,
	Args:    cobra.ExactArgs(0),
}

func supplyPassword(db *bolt.DB) (string, error) {
	fmt.Print("Enter your actual password: ")

	var oldPassword string
	fmt.Print("\033[8m")
	fmt.Scanln(&oldPassword)
	fmt.Print("\033[28m")

	validateAuth(db, oldPassword)

	fmt.Print("Enter your new password: ")

	var newPassword string
	fmt.Print("\033[8m")
	fmt.Scanln(&newPassword)
	fmt.Print("\033[28m")

	fmt.Print("Enter your new password again: ")

	var newPasswordAgain string
	fmt.Print("\033[8m")
	fmt.Scanln(&newPasswordAgain)
	fmt.Print("\033[28m")

	err := matchPassword(newPassword, newPasswordAgain)
	if err != nil {
		fmt.Println(err)
	}
	return newPassword, nil
}
