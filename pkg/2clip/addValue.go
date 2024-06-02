package clip

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"

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

		db, _ := database.OpenDatabase()
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

		db, _ := database.OpenDatabase()
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
		bucket := tx.Bucket([]byte("2clip"))
		if bucket == nil {
			return fmt.Errorf("bucket 2clip not found")
		}
		err := bucket.Put([]byte(key), []byte(value))
		if err != nil {
			return err
		}

		fmt.Printf(`Added "%s" with value "%s"`+"\n", key, value)
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

		err := bucket.Put([]byte(key+" (protected)"), []byte(value))
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
