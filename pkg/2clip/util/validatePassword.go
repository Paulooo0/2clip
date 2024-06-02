package util

import (
	"github.com/boltdb/bolt"
)

func ValidatePassword(tx *bolt.Tx, password string) bool {
	authBucket, _ := ConnectToBucket(tx, "2clip_password")

	value := authBucket.Get([]byte("2CLIP_PASSWORD"))
	if value == nil {
		return false
	}
	if string(value) != password {
		return false
	}
	return true
}
