package util

import (
	"fmt"

	"github.com/boltdb/bolt"
)

func ValidatePassword(tx *bolt.Tx, password string) bool {
	authBucket, _ := ConnectToBucket(tx, "2clip_password")

	value := authBucket.Get([]byte("2CLIP_PASSWORD"))
	if value == nil {
		return false
	} else if string(value) != password {
		return false
	} else if len(password) < 1 {
		fmt.Println("Your password cannot be empty")
		return false
	}
	return true
}
