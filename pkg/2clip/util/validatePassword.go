package util

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

func ValidatePassword(password string) error {
	if password == "" {
		fmt.Println("Password cannot be empty")
	} else {
		return nil
	}
	return nil
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
