package database

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/Paulooo0/2clip/internal/2clip/util"
	"github.com/boltdb/bolt"
)

func OpenDatabase(dbName string, bucketName string) (*bolt.DB, error) {
	// Get the home directory path
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("%s Failed to get home directory: %w", util.Err, err)
	}

	// Construct the data folder path
	dataDir := filepath.Join(homeDir, ".2clip", "data")

	// Create the data folder if it doesn't exist
	err = os.MkdirAll(dataDir, 0700)
	if err != nil {
		return nil, fmt.Errorf("%s Failed to create data folder: %w", util.Err, err)
	}

	// Construct the database path
	dbPath := filepath.Join(dataDir, dbName)

	db, err := bolt.Open(dbPath, 0644, nil)
	if err != nil {
		log.Fatalf("Failed to open database: %s", err)
	}

	CreateBucketIfNotExists(db, bucketName)

	return db, nil
}

func CreateBucketIfNotExists(db *bolt.DB, name string) {
	err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(name))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
	if err != nil {
		log.Fatalf("%s %v", util.Err, err)
	}
}
