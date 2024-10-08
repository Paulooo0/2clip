package remove

import (
	"fmt"
	"log"
	"strings"

	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"

	"github.com/Paulooo0/2clip/internal/2clip/util"
	"github.com/Paulooo0/2clip/internal/database"
)

var RemoveCmd = &cobra.Command{
	Use:        "remove",
	Aliases:    []string{"rm"},
	ValidArgs:  []string{"-i"},
	ArgAliases: []string{"-i"},
	Short:      "Remove a key-value from the database",
	Long:       `Remove a key-value from the database based on the provided key.`,
	Example: `
	2clip remove <key>
	2clip remove [ARG] <key>
	`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]

		if cmd.Flags().NFlag() == 0 {
			commandRemove(key)
		}
		if cmd.Flags().Changed("index") || cmd.Flags().Changed("i") {
			commandRemoveByIndex(key)
		}
	},
}

func RemoveCmdFlags() {
	RemoveCmd.Flags().BoolP("index", "i", false, "Remove a value from the database by index")
}

func commandRemove(key string) {
	db, _ := database.OpenDatabase("2clip.db", "2clip")
	defer db.Close()

	removeValue(db, key)
}

func removeValue(db *bolt.DB, key string) {
	err := db.Update(func(tx *bolt.Tx) error {
		bucket, err := util.ConnectToBucket(tx, "2clip")
		if err != nil {
			return err
		}

		key, err := processKey(bucket, key)
		if err != nil {
			return err
		}

		err = deleteByKey(db, bucket, key)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

func processKey(bucket *bolt.Bucket, key string) (string, error) {
	value := bucket.Get([]byte(key + " (protected)"))
	originalKey := key
	key = key + " (protected)"
	if value == nil {
		key = originalKey
		value = bucket.Get([]byte(key))

		if value == nil {
			return "", fmt.Errorf(`key "%s" not found`, key)
		}
	}
	return key, nil
}

func deleteByKey(db *bolt.DB, bucket *bolt.Bucket, key string) error {
	if strings.HasSuffix(key, " (protected)") {
		err := util.Authenticate(db)
		if err != nil {
			return err
		}

		err = bucket.Delete([]byte(key))
		if err != nil {
			return err
		}

		fmt.Printf(`Removed "%s"`+"\n", key[:len(key)-12])
		return nil
	} else {
		err := bucket.Delete([]byte(key))
		if err != nil {
			return err
		}

		fmt.Printf(`Removed "%s"`+"\n", key)
		return nil
	}
}
