package clip

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"

	"github.com/Paulooo0/2clip/pkg/database"
)

var RemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a value from the database",
	Long:  `Remove a value from the database based on the provided key.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]
		db, _ := database.OpenDatabase()
		defer db.Close()
		removeValue(db, key)
	},
}

func removeValue(db *bolt.DB, key string) {
	err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("2clip"))
		if bucket == nil {
			return fmt.Errorf("bucket 2clip not found")
		}
		value := bucket.Get([]byte(key))
		if value == nil {
			return fmt.Errorf(`key "%s" not found`, key)
		}
		err := bucket.Delete([]byte(key))
		if err != nil {
			return err
		}
		fmt.Printf(`Removed "%s"`+"\n", key)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
