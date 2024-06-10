package util

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

func ValidatePassword(password string) bool {
	condition := true
	for condition {
		if len(password) < 1 {
			fmt.Println("Password cannot be empty")
			return condition == TryAgain()
		} else {
			return condition == false
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
			return fmt.Errorf("Wrong password!")
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
