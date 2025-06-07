package list

import (
	"fmt"
	"log"
	"strings"

	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"

	"github.com/Paulooo0/2clip/internal/2clip/util"
	"github.com/Paulooo0/2clip/internal/database"
)

var ListKeysCmd = &cobra.Command{
	Use:     "list",
	Short:   "List all keys in the database",
	Aliases: []string{"ls"},
	Long:    `List all keys in the database.`,
	Example: `
	2clip ls
	`,
	Run: func(cmd *cobra.Command, args []string) {
		db, _ := database.OpenDatabase("2clip.db", "2clip")
		defer db.Close()
		listKeys(db)
	},
}

func listKeys(db *bolt.DB) {
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("2clip"))
		if bucket == nil {
			return fmt.Errorf("bucket not found")
		}

		keys, err := util.GetSortedKeys(bucket)
		if err != nil {
			return err
		}

		printSortedKeys(keys)
		return nil
	})
	if err != nil {
		log.Fatalf("%s %v", util.Err, err)
	}
}

func printSortedKeys(keys []string) {
	if len(keys) == 0 {
		return
	}

	prevLetter := strings.ToUpper(string(keys[0][0]))
	fmt.Printf("\n\033[1m"+"\033[32m"+"➜  "+"\033[0m"+"\033[1m"+"%s"+"\033[0m\n", prevLetter)

	for i, key := range keys {
		letter := strings.ToUpper(string(key[0]))
		if letter != prevLetter {
			fmt.Printf("\n\033[1m"+"\033[32m"+"➜  "+"\033[0m"+"\033[1m"+"%s"+"\033[0m\n", letter)
			prevLetter = letter
		}
		if strings.HasSuffix(key, " (protected)") {
			fmt.Printf("\033[94m"+"%d"+"\033[0m %s 🔒\n", i+1, key[:len(key)-12])
		} else {
			fmt.Printf("\033[33m"+"%d"+"\033[0m %s\n", i+1, key)
		}
	}
}
