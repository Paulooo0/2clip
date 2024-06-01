package clip

import (
	"flag"
	"fmt"
	"log"

	"github.com/boltdb/bolt"

	database "github.com/Paulooo0/2clip/pkg/database"
)

func GetValue() {
	// Define command-line flags
	key := flag.String("get", "", "The key to search for")

	// Parse command-line flags
	flag.Parse()

	// Access the flag value
	searchKey := *key

	// Open or create the database
	db := database.OpenDatabase()

	// Populate the database
	populateDatabase(db)

	readValue(db, searchKey)
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

func populateDatabase(db *bolt.DB) {
	err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("2clip"))

		testData := map[string]string{
			"key1": "value1",
			"key2": "value2",
			"key3": "value3",
		}

		for key, value := range testData {
			err := bucket.Put([]byte(key), []byte(value))
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}

// func cliFlags() []string {
// 	getFlag := flag.String("get", "", "The key to search for")
// 	gFlag := flag.String("g", "", "The key to search for")
// 	emptyFlag := flag.String("", "", "The key to search for")

// 	flag.Parse()
// 	return []string{*getFlag, *gFlag, *emptyFlag}
// }
