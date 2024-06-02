package clip

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"

	"github.com/Paulooo0/2clip/pkg/2clip/util"
	"github.com/Paulooo0/2clip/pkg/database"
)

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a value to the database",
	Long:  "Add a value to the database based on the provided key.",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]
		value := args[1]

		db, _ := database.OpenDatabase("2clip.db", "2clip")
		defer db.Close()

		if cmd.Flags().Changed("protected") {
			addProtectedToDatabase(db, key, value)
			return
		}
		addToDatabase(db, key, value)
	},
}

var AddProtectedCmd = &cobra.Command{
	Use:     "protected",
	Aliases: []string{"-p"},
	Short:   "Add a protected value to the database",
	Long:    "Add a protected value to the database based on the provided key.",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]
		value := args[1]

		db, _ := database.OpenDatabase("2clip", "2clip")
		defer db.Close()

		addProtectedToDatabase(db, key, value)
	},
}

func AddCmdFlags() {
	AddCmd.Flags().BoolP("protected", "p", true, "Add a protected value to the database")

	AddCmd.AddCommand(AddProtectedCmd)
}

func addToDatabase(db *bolt.DB, key string, value string) {
	err := db.Update(func(tx *bolt.Tx) error {
		bucket, err := util.ConnectToBucket(tx, "2clip")
		if err != nil {
			return err
		}

		key, err = overwrite(tx, key)
		if err != nil {
			return err
		}

		err = bucket.Put([]byte(key), []byte(value))
		if err != nil {
			return err
		}

		if strings.HasSuffix(key, " (protected)") {
			fmt.Printf(`Added '%s' with protect value`+"\n", key)
		} else {
			fmt.Printf(`Added '%s' with value "%s"`+"\n", key, value)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

func addProtectedToDatabase(db *bolt.DB, key string, value string) {
	err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("2clip"))
		if bucket == nil {
			return fmt.Errorf("bucket 2clip not found")
		}

		key, err := overwrite(tx, key)
		if err != nil {
			return err
		}

		err = bucket.Put([]byte(key+" (protected)"), []byte(value))
		if err != nil {
			return err
		}
		fmt.Printf(`Added "%s" with protect value`+"\n", key)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

func overwrite(tx *bolt.Tx, key string) (string, error) {
	if util.CheckKeyAlreadyExists(key, tx, "2clip") {
		fmt.Printf(`key '%s' already exists, you want to overwrite it? [Y/N]: `, key)

		getOverwriteAnswer()

		return key, nil
	}
	if util.CheckKeyAlreadyExists(key+" (protected)", tx, "2clip") {
		fmt.Printf(`key '%s' already exists, you want to overwrite it? [Y/N]: `, key)

		getOverwriteAnswer()

		err := util.Authenticate(tx)
		if err != nil {
			return "", err
		}
		key = key + " (protected)"
		return key, nil
	}
	return key, nil
}

func getOverwriteAnswer() {
	answerCondition := true
	for answerCondition {
		var answer string
		fmt.Scanln(&answer)
		if answer == "N" || answer == "n" {
			os.Exit(0)
		} else if answer == "Y" || answer == "y" {
			answerCondition = false
		} else {
			fmt.Print("Invalid answer, please type Y or N: ")
			answerCondition = true
		}
	}
}
