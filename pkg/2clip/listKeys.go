package clip

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"

	"github.com/Paulooo0/2clip/pkg/database"
)

var ListKeysCmd = &cobra.Command{
	Use:   "list",
	Short: "List all keys in the database",
	Long:  `List all keys in the database.`,
	Run: func(cmd *cobra.Command, args []string) {
		db, _ := database.OpenDatabase()
		defer db.Close()
		listKeys(db)
	},
}

func listKeys(db *bolt.DB) {
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("2clip"))
		if bucket == nil {
			return fmt.Errorf("bucket 2clip not found")
		}

		// Create a slice of keys
		keys := make([]string, 0, bucket.Stats().KeyN)
		bucket.ForEach(func(k, _ []byte) error {
			keys = append(keys, string(k))
			return nil
		})

		// Sort the keys
		sort.Strings(keys)

		// Print the keys with sorting by letter
		prevLetter := ""
		for _, key := range keys {
			letter := string(key[0])
			if letter != prevLetter {
				upperLetter := strings.ToUpper(letter)
				fmt.Printf("\n%s\n", upperLetter)
				prevLetter = letter
			}
			fmt.Println(key)
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}
