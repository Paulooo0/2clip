package util

import (
	"log"

	"github.com/boltdb/bolt"
)

func CheckKeyAlreadyExists(key string, db *bolt.DB, bucketName string) bool {
	var exists bool
	err := db.View(func(tx *bolt.Tx) error {
		bucket, err := ConnectToBucket(tx, bucketName)
		if err != nil {
			log.Fatal(err)
		}
		exists = bucket.Get([]byte(key)) != nil

		return nil
	})
	if err != nil {
		log.Fatal(err)
		return exists == false
	}
	return exists == true
}
