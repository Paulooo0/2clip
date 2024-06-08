package util

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

func SaveAuthentication(db *bolt.DB, password string) {
	err := db.Update(func(tx *bolt.Tx) error {
		bucket, err := ConnectToBucket(tx, "2clip_password")
		if err != nil {
			return err
		}
		err = bucket.Put([]byte("2CLIP_PASSWORD"), []byte(password))
		if err != nil {
			return err
		}
		fmt.Println("Authentication saved")
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

func Authenticate(db *bolt.DB) error {
	condition := true
	for condition {
		var password string
		fmt.Print("\nEnter your password: ")
		fmt.Print("\033[8m")
		fmt.Scanln(&password)
		fmt.Print("\033[28m")

		ValidatePassword(password)
		err := CheckPassword(db, password)
		if err != nil {
			condition = true
			isAuthenticationFailed(condition)
		} else {
			condition = false
			isAuthenticationFailed(condition)
		}

	}

	return nil
}

func isAuthenticationFailed(condition bool) {
	if condition {
		fmt.Println("\nAuthentication failed!")
		fmt.Println("TIP: if you don't have a password, run '2clip auth' to create one")

	} else {
		fmt.Println("\nAuthentication successful!")
	}
}
