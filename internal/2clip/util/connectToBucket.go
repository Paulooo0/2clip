package util

import (
	"fmt"

	"github.com/boltdb/bolt"
)

func ConnectToBucket(tx *bolt.Tx, name string) (*bolt.Bucket, error) {
	bucket := tx.Bucket([]byte(name))
	if bucket == nil {
		return nil, fmt.Errorf("bucket %s not found", name)
	}
	return bucket, nil
}
