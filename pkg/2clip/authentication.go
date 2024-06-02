package clip

import (
	"fmt"
	"log"
	"os"

	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"

	"github.com/Paulooo0/2clip/pkg/2clip/util"
	"github.com/Paulooo0/2clip/pkg/database"
)

var AuthCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate to the database",
	Long:  `Authenticate to the database.`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		db, _ := database.OpenDatabase("2clip.db", "2clip_password")
		defer db.Close()

		checkAlreadyHaveAuth(db)

		password := getPassword()

		validateAuth(db, password)

		saveAuthentication(db, password)
	},
}

func validateAuth(db *bolt.DB, password string) {
	err := db.Update(func(tx *bolt.Tx) error {
		util.ValidatePassword(tx, password)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

func checkAlreadyHaveAuth(db *bolt.DB) error {
	err := db.Update(func(tx *bolt.Tx) error {
		if util.CheckKeyAlreadyExists("2CLIP_PASSWORD", tx, "2clip_password") {
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
		fmt.Print("Enter your password: ")

		fmt.Scanln(&password)

		fmt.Print("Enter your password again: ")

		var passwordAgain string
		fmt.Scanln(&passwordAgain)

		err := matchPassword(password, passwordAgain)

		if err != nil {
			condition = checkAnswer()
		} else {
			condition = false
		}
	}
	return password
}

func matchPassword(password1 string, password2 string) error {
	if password1 != password2 {
		return fmt.Errorf("passwords do not match")
	}
	return nil
}

func checkAnswer() bool {
	answerCondition := true
	for answerCondition {
		fmt.Print("\nYou want to try again? [Y/N]: ")

		var answer string
		fmt.Scanln(&answer)
		if answer == "N" || answer == "n" {
			answerCondition = false
			os.Exit(0)
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

func saveAuthentication(db *bolt.DB, password string) {
	err := db.Update(func(tx *bolt.Tx) error {
		bucket, err := util.ConnectToBucket(tx, "2clip_password")
		if err != nil {
			return err
		}
		err = bucket.Put([]byte("2CLIP_PASSWORD"), []byte(password))
		if err != nil {
			return err
		}
		fmt.Println("Authentication saved")
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
