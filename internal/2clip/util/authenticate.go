package util

import (
	"fmt"
	"log"
	"syscall"

	"github.com/boltdb/bolt"
	"golang.org/x/term"
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
		fmt.Println("\033[1m\033[0m\033[32m" + "Authentication saved" + "\033[0m")
		return nil
	})
	if err != nil {
		log.Fatalf("%s %v", Err, err)
	}
}

func Authenticate(db *bolt.DB) error {
	condition := true
	for condition {
		fmt.Print("Enter your password: ")

		bytePassword, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			return err
		}

		password := string(bytePassword)

		err = CheckPassword(db, password)
		if err != nil {
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
		fmt.Println("\n\033[1m\033[0m\033[31m" + "Authentication failed!" + "\033[0m")
		fmt.Println("TIP: if you don't have a password, run '2clip auth' to create one")
		AskTryAgain(true)
	} else {
		fmt.Println("\n\033[1m\033[0m\033[32m" + "Successfully authenticated!" + "\033[0m")
	}
}
