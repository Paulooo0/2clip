package util

import (
	"log"

	"github.com/boltdb/bolt"
)

func CheckKeyAlreadyExists(key string, tx *bolt.Tx, bucketName string) bool {
	var exists bool

	bucket, err := ConnectToBucket(tx, bucketName)
	if err != nil {
		log.Fatal(err)
	}
	exists = bucket.Get([]byte(key)) != nil

	return exists
}
