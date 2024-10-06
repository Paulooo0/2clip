package add

import (
	"fmt"
	"log"
	"strings"

	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"

	"github.com/Paulooo0/2clip/internal/2clip/util"
	"github.com/Paulooo0/2clip/internal/database"
)

var AddCmd = &cobra.Command{
	Use:        "add",
	Aliases:    []string{"a"},
	ValidArgs:  []string{"-p"},
	ArgAliases: []string{"-p"},
	Short:      "Add a key-value to database",
	Long:       "Add a value by input to database, linked to the provided key.",
	Example: `
	2clip add <key>
	2clip add [ARG] <key>
	`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]

		if cmd.Flags().NFlag() == 0 {
			commandAdd(key)
		}
		if (cmd.Flags().Changed("protected")) || (cmd.Flags().Changed("p")) {
			CommandAddProtected(key)
		}
	},
}

func AddCmdFlags() {
	AddCmd.Flags().BoolP("protected", "p", false, "Add a protected value to the database")
}

func commandAdd(key string) {
	db, _ := database.OpenDatabase("2clip.db", "2clip")
	defer db.Close()

	addToDatabase(db, key)
}

func addToDatabase(db *bolt.DB, key string) {
	err := db.Update(func(tx *bolt.Tx) error {
		bucket, err := util.ConnectToBucket(tx, "2clip")
		if err != nil {
			return err
		}

		key, err = overwrite(db, key)
		if err != nil {
			return err
		}

		fmt.Println("Input value:")
		var value string
		fmt.Scanln(&value)

		err = bucket.Put([]byte(key), []byte(value))
		if err != nil {
			return err
		}

		if strings.HasSuffix(key, " (protected)") {
			fmt.Printf(`Added "%s" with protect value`+"\n", key)
		} else {
			fmt.Printf(`Added "%s" with value "%s"`+"\n", key, value)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

func overwrite(db *bolt.DB, key string) (string, error) {
	if util.CheckKeyAlreadyExists(key, db, "2clip") {
		getOverwriteAnswer(key)

		return key, nil
	}
	if util.CheckKeyAlreadyExists(key+" (protected)", db, "2clip") {
		getOverwriteAnswer(key)

		err := util.Authenticate(db)
		if err != nil {
			return "", err
		}
		key = key + " (protected)"
		return key, nil
	}
	return key, nil
}

func getOverwriteAnswer(key string) {
	fmt.Printf(`key '%s' already exists, you want to overwrite it? [Y/N]: `, key)

	util.AnswerCondition()
}
