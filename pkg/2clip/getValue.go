package clip

import (
	"fmt"
	"log"
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

		commandGet(key)
	},
}

func commandGet(key string) {
	db, _ := database.OpenDatabase("2clip.db", "2clip")

	defer db.Close()

	readValue(db, key)
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
				util.Authenticate(db)
			}
		} else {
			value = bucket.Get([]byte(key))
			keyString = key
			if value == nil {
				return fmt.Errorf(`key '%s' not found`, key)
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
