package util

import (
	"fmt"

	"github.com/boltdb/bolt"
)

func ValidatePassword(password string) bool {
	condition := true
	for condition {
		if len(password) < 1 {
			fmt.Println("Password can't be empty")
			return condition == AskTryAgain(true)
		} else {
			return !condition
		}
	}
	return condition
}

func CheckPassword(db *bolt.DB, password string) error {
	err := db.View(func(tx *bolt.Tx) error {
		authBucket, err := ConnectToBucket(tx, "2clip_password")
		if err != nil {
			return err
		}

		value := authBucket.Get([]byte("2CLIP_PASSWORD"))

		if password != string(value) {
			return fmt.Errorf("%s %v", Err, err)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("%s %v", Err, err)
	}
	return nil
}
