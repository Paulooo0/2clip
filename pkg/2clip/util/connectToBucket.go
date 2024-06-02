package util

import (
	"fmt"

	"github.com/boltdb/bolt"
)

func ConnectToBucket(tx *bolt.Tx) (*bolt.Bucket, error) {
	bucket := tx.Bucket([]byte("2clip"))
	if bucket == nil {
		return nil, fmt.Errorf("bucket 2clip not found")
	}
	return bucket, nil
}
