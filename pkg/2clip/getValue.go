package clip

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"

	"github.com/Paulooo0/2clip/pkg/2clip/util"
	database "github.com/Paulooo0/2clip/pkg/database"
)

var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a value from the database",
	Long:  `Get a value from the database based on the provided key.`,
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]

		db, _ := database.OpenDatabase("2clip.db", "2clip")

		defer db.Close()

		readValue(db, key)
	},
}

func readValue(db *bolt.DB, key string) {
	err := db.View(func(tx *bolt.Tx) error {
		bucket, err := util.ConnectToBucket(tx, "2clip")
		if err != nil {
			return err
		}

		value := bucket.Get([]byte(key + " (protected)"))
		keyString := key + " (protected)"
		if value != nil {
			condition := true
			for condition {
				var password string
				fmt.Print("\nEnter your password: ")
				fmt.Scanln(&password)

				authenticated := util.ValidatePassword(tx, password)

				condition = isAutheticationFailed(authenticated)
			}
		} else {
			value = bucket.Get([]byte(key))
			keyString = key
			if value == nil {
				return fmt.Errorf(`key "%s" not found`, key)
			}
		}

		if strings.HasSuffix(keyString, " (protected)") {
			err := clipboard.WriteAll(string(value))

			fmt.Println("Protected value copied to clipboard")
			if err != nil {
				log.Fatal(err)
			}
		} else {
			fmt.Println(string(value))

			err := clipboard.WriteAll(string(value))
			fmt.Println("Value copied to clipboard")
			if err != nil {
				log.Fatal(err)
			}
		}

		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
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
