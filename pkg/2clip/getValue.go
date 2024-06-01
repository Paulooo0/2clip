package clip

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"

	database "github.com/Paulooo0/2clip/pkg/database"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a value from the database",
	Long:  `Get a value from the database based on the provided key.`,
	Run: func(cmd *cobra.Command, args []string) {
		searchKey := args[0]

		db := database.OpenDatabase()

		populateDatabase(db) // Add test data to the database
		defer db.Close()

		readValue(db, searchKey)
	},
}

var rootCmd = &cobra.Command{
	Use:   "2clip",
	Short: "2clip is a simple clipboard manager",
	Long:  `2clip is a simple CLI tool for managing your clipboard`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("2clip is a simple CLI tool for managing your clipboard")
	},
}

func GetValue() {
	getCmd.Flags().String("get", "", "The key to search for")

	rootCmd.AddCommand(getCmd)
	rootCmd.Execute()
}

func readValue(db *bolt.DB, searchKey string) {
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("2clip"))
		if b == nil {
			return fmt.Errorf("bucket 2clip not found")
		}
		v := b.Get([]byte(searchKey))
		if v == nil {
			return fmt.Errorf("key %s not found", searchKey)
		}
		fmt.Println(string(v))
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

func populateDatabase(db *bolt.DB) {
	err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("2clip"))

		testData := map[string]string{
			"key1": "value1",
			"key2": "value2",
			"key3": "value3",
		}

		for key, value := range testData {
			err := bucket.Put([]byte(key), []byte(value))
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}
