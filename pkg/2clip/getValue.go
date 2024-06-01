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

		db := database.OpenDatabase()

		defer db.Close()

		readValue(db, key)
	},
}

func readValue(db *bolt.DB, searchKey string) {
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("2clip"))
		if b == nil {
			return fmt.Errorf("bucket 2clip not found")
		}
		v := b.Get([]byte(searchKey))
		if v == nil {
			return fmt.Errorf("key %s not found", searchKey)
		}
		fmt.Println(string(v))
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
