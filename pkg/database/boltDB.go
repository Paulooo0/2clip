package database

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/boltdb/bolt"
)

func OpenDatabase() (*bolt.DB, error) {
	// Get the home directory path
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get home directory: %w", err)
	}

	// Construct the data folder path
	dataDir := filepath.Join(homeDir, ".2clip", "data")

	// Create the data folder if it doesn't exist
	err = os.MkdirAll(dataDir, 0700)
	if err != nil {
		return nil, fmt.Errorf("failed to create data folder: %w", err)
	}

	// Construct the database path
	dbPath := filepath.Join(dataDir, "2clip.db")

	db, err := bolt.Open(dbPath, 0644, nil)
	if err != nil {
		log.Fatalf("Failed to open database: %s", err)
	}

	createBucketIfNotExists(db)

	return db, nil
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
