package add

import (
	"fmt"
	"log"
	"strings"

	"github.com/Paulooo0/2clip/pkg/database"
	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
)

var AddProtectedCmd = &cobra.Command{
	Use:     "protected",
	Aliases: []string{"-p"},
	Short:   "Add a protected value to the database",
	Long:    "Add a protected value to the database based on the provided key.",
	Args:    cobra.ExactArgs(2),
}

func CommandAddProtected(key string, value string) {
	db, _ := database.OpenDatabase("2clip.db", "2clip")
	defer db.Close()

	addProtectedToDatabase(db, key, value)
}

func addProtectedToDatabase(db *bolt.DB, key string, value string) {
	err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("2clip"))
		if bucket == nil {
			return fmt.Errorf("bucket 2clip not found")
		}

		key, err := overwrite(db, key)
		if err != nil {
			return err
		}
		err = addProtectedValue(key, value, bucket)
		if err != nil {
			return err
		}
		fmt.Printf(`Added '%s' with protect value`+"\n", key)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

func addProtectedValue(key string, value string, bucket *bolt.Bucket) error {
	if strings.HasSuffix(key, " (protected)") {
		err := bucket.Put([]byte(key), []byte(value))
		return err
	} else {
		err := bucket.Put([]byte(key+" (protected)"), []byte(value))
		return err
	}
}
