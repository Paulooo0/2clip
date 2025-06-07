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

var AuthCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate to the database",
	Long:  `Authenticate to the database.`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flags().NFlag() == 0 {
			CommandAuth()
		} else if cmd.Flags().Changed("update") {
			CommandAuthUpdate()
		}
	},
}

func AuthCmdFlags() {
	AuthCmd.Flags().BoolP("update", "u", true, "Update your authentication")

	AuthCmd.AddCommand(UpdateAuthCmd)
}

func CommandAuth() {
	db, _ := database.OpenDatabase("2clip.db", "2clip_password")
	defer db.Close()

	checkAlreadyHaveAuth(db)

	password := getPassword()

	util.SaveAuthentication(db, password)
}

func checkAlreadyHaveAuth(db *bolt.DB) error {
	err := db.Update(func(tx *bolt.Tx) error {
		if util.CheckKeyAlreadyExists("2CLIP_PASSWORD", db, "2clip_password") {
			fmt.Println("You already have an authentication")
			fmt.Println("If want to update your password, run '2clip auth -u'")
			os.Exit(0)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func getPassword() string {
	var password string

	condition := true
	for condition {
		password = enterPassword()
		passwordAgain := enterPasswordAgain()

		err := matchPassword(password, passwordAgain)
		condition = matchLoop(err)
	}

	return password
}

func enterPassword() string {
	var password string

	condition := true
	for condition {
		fmt.Print("Enter your password: ")

		fmt.Print("\033[8m")
		fmt.Scanln(&password)
		fmt.Print("\033[28m")

		condition = util.ValidatePassword(password)
	}

	return password
}

func enterPasswordAgain() string {
	fmt.Print("Enter your password again: ")

	var passwordAgain string
	fmt.Print("\033[8m")
	fmt.Scanln(&passwordAgain)
	fmt.Print("\033[28m")

	return passwordAgain
}

func matchPassword(password1 string, password2 string) error {
	if password1 != password2 {
		return fmt.Errorf("%s Passwords don't match", util.Err)
	}
	return nil
}

func matchLoop(err error) bool {
	if err != nil {
		fmt.Println(err)
		util.AskTryAgain(true)
		return true
	}
	return false
}
