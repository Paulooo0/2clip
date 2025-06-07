package util

import (
	"fmt"

	"github.com/boltdb/bolt"
)

func ConnectToBucket(tx *bolt.Tx, name string) (*bolt.Bucket, error) {
	bucket := tx.Bucket([]byte(name))
	if bucket == nil {
		return nil, fmt.Errorf("%s bucket %s not found", Err, name)
	}
	return bucket, nil
}
