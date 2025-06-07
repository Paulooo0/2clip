package add

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Paulooo0/2clip/internal/2clip/util"
	"github.com/Paulooo0/2clip/internal/database"
	"github.com/boltdb/bolt"
)

func CommandAddProtected(key string) {
	db, _ := database.OpenDatabase("2clip.db", "2clip")
	defer db.Close()

	key, err := overwrite(db, key)
	if err != nil {
		log.Printf("overwrite failed: %v", err)
		os.Exit(0)
	}

	input := GetInput('\n')
	addProtectedToDatabase(db, key, input)
}

func addProtectedToDatabase(db *bolt.DB, key string, value string) {
	err := db.Update(func(tx *bolt.Tx) error {
		bucket, err := util.ConnectToBucket(tx, "2clip")
		if err != nil {
			log.Print("bucket 2clip not found")
		}

		err = addProtectedValue(key, value, bucket)
		if err != nil {
			return err
		}
		fmt.Printf(`Added '%s' with protected value`+"\n", key)
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

	// TODO: Change sufix to a boolean value
}
