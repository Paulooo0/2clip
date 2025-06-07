package add

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Paulooo0/2clip/internal/2clip/util"
	"github.com/Paulooo0/2clip/internal/database"
	"github.com/boltdb/bolt"
)

func CommandAddExtended(key string) {
	db, _ := database.OpenDatabase("2clip.db", "2clip")
	defer db.Close()

	addExtendedToDatabase(key, db, "2clip")
}

func addExtendedToDatabase(key string, db *bolt.DB, bucketName string) {
	value := getExtendedInput()

	err := db.Update(func(tx *bolt.Tx) error {
		bucket, err := util.ConnectToBucket(tx, bucketName)
		if err != nil {
			return err
		}

		err = bucket.Put([]byte(key), []byte(value))
		if err != nil {
			return err
		}

		fmt.Printf(`Added "%s" successfully`+"\n", key)
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}

func getExtendedInput() string {
	fmt.Println("Enter your text (type 'END' on a new line to finish):")
	reader := bufio.NewReader(os.Stdin)
	var lines []string

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			os.Exit(1)
		}

		line = strings.TrimSpace(line)
		if line == "END" {
			break
		}

		lines = append(lines, line)
	}

	input := strings.Join(lines, "\n")

	return input
}
