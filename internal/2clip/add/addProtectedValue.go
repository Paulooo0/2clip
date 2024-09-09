package add

import (
	"fmt"
	"log"
	"strings"

	"github.com/Paulooo0/2clip/internal/database"
	"github.com/boltdb/bolt"
)

func CommandAddProtected(key string) {
	db, _ := database.OpenDatabase("2clip.db", "2clip")
	defer db.Close()

	addProtectedToDatabase(db, key)
}

func addProtectedToDatabase(db *bolt.DB, key string) {
	err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("2clip"))
		if bucket == nil {
			return fmt.Errorf("bucket 2clip not found")
		}

		key, err := overwrite(db, key)
		if err != nil {
			return err
		}

		fmt.Println("Input value:")
		var value string
		fmt.Scanln(&value)

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
