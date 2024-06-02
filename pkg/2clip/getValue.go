package clip

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"

	database "github.com/Paulooo0/2clip/pkg/database"
)

var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a value from the database",
	Long:  `Get a value from the database based on the provided key.`,
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]

		db, _ := database.OpenDatabase()

		defer db.Close()

		readValue(db, key)
	},
}

func readValue(db *bolt.DB, key string) {
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("2clip"))
		if bucket == nil {
			return fmt.Errorf("bucket 2clip not found")
		}
		value := bucket.Get([]byte(key))
		if value == nil {
			return fmt.Errorf(`key "%s" not found`, key)
		}
		fmt.Println(string(value))
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
