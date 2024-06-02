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

		db := database.OpenDatabase()
		defer db.Close()

		addToDatabase(db, key, value)
	},
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
