package auth

import (
	"fmt"
	"log"
	"os"

	"github.com/Paulooo0/2clip/internal/2clip/util"
	"github.com/Paulooo0/2clip/internal/database"
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

func CommandAuthUpdate() {
	db, _ := database.OpenDatabase("2clip.db", "2clip_password")
	defer db.Close()

	password, err := supplyPassword(db)
	if err != nil {
		log.Fatal(err)
	}
	util.SaveAuthentication(db, password)
}

func supplyPassword(db *bolt.DB) (string, error) {
	newPassword, newPasswordAgain := enterPasswords(db)

	err := matchPassword(newPassword, newPasswordAgain)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	return newPassword, nil
}

func enterPasswords(db *bolt.DB) (string, string) {
	oldPassword := getOldPassword()

	util.CheckPassword(db, oldPassword)

	newPassword := getNewPassword()

	newPasswordAgain := getNewPasswordAgain()

	return newPassword, newPasswordAgain
}

func getOldPassword() string {
	fmt.Print("Enter your current password: ")

	var oldPassword string
	fmt.Print("\033[8m")
	fmt.Scanln(&oldPassword)
	fmt.Print("\033[28m")

	return oldPassword
}

func getNewPassword() string {
	var newPassword string
	condition := true
	for condition {
		fmt.Print("Enter your new password: ")

		fmt.Print("\033[8m")
		fmt.Scanln(&newPassword)
		fmt.Print("\033[28m")

		condition = util.ValidatePassword(newPassword)
	}
	return newPassword
}

func getNewPasswordAgain() string {
	fmt.Print("Enter your new password again: ")

	var newPasswordAgain string
	fmt.Print("\033[8m")
	fmt.Scanln(&newPasswordAgain)
	fmt.Print("\033[28m")

	return newPasswordAgain
}
