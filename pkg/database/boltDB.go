package database

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/boltdb/bolt"
)

func OpenDatabase() *bolt.DB {
	err := os.MkdirAll("data", 0755)
	if err != nil {
		log.Fatal(err)
	}

	dbPath := filepath.Join("data", "2clip_db.db")
	db, err := bolt.Open(dbPath, 0644, nil)
	if err != nil {
		log.Fatalf("Failed to open database: %s", err)
	}

	createBucketIfNotExists(db)

	return db
}

func createBucketIfNotExists(db *bolt.DB) {
	err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("2clip"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
