package get

import (
	"fmt"
	"log"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"

	"github.com/Paulooo0/2clip/internal/2clip/util"
	database "github.com/Paulooo0/2clip/internal/database"
)

var GetCmd = &cobra.Command{
	Use:        "get",
	Aliases:    []string{"g"},
	ValidArgs:  []string{"-i"},
	ArgAliases: []string{"-i"},
	Short:      "Get a value from the database",
	Long:       `Get a value from the database based on the provided key.`,
	Example: `
	2clip get <key>
	2clip get [ARG] <key>
	`,
	Run: func(cmd *cobra.Command, args []string) {
		arg := args[0]

		if cmd.Flags().NFlag() == 0 {
			commandGet(arg)
		}
		if cmd.Flags().Changed("index") || cmd.Flags().Changed("i") {
			commandGetByIndex(arg)
		}
	},
}

func GetCmdFlags() {
	GetCmd.Flags().BoolP("index", "i", false, "Get a value from the database by index")
}

func commandGet(key string) {
	db, _ := database.OpenDatabase("2clip.db", "2clip")

	defer db.Close()

	readValue(db, key)
}

func readValue(db *bolt.DB, key string) {
	err := db.View(func(tx *bolt.Tx) error {
		bucket, err := util.ConnectToBucket(tx, "2clip")
		if err != nil {
			return err
		}

		value := bucket.Get([]byte(key + " (protected)"))
		keyString := key + " (protected)"
		if value != nil {
			err := util.Authenticate(db)
			if err != nil {
				return err
			}
		} else {
			value = bucket.Get([]byte(key))
			keyString = key
			if value == nil {
				return fmt.Errorf(`key '%s' not found`, key)
			}
		}

		copyToClipboard(keyString, value)

		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

func copyToClipboard(keyString string, value []byte) {
	if strings.HasSuffix(keyString, " (protected)") {
		err := clipboard.WriteAll(string(value))

		fmt.Printf(`"%s" protected value copied to clipboard`, keyString[:len(keyString)-12])
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println(string(value))

		err := clipboard.WriteAll(string(value))
		fmt.Println("Value copied to clipboard")
		if err != nil {
			log.Fatal(err)
		}
	}
}
