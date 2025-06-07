package add

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"

	"github.com/Paulooo0/2clip/internal/2clip/util"
	"github.com/Paulooo0/2clip/internal/database"
)

var AddCmd = &cobra.Command{
	Use:        "add",
	Aliases:    []string{"a"},
	ValidArgs:  []string{"-p", "-e"},
	ArgAliases: []string{"-p", "-e"},
	Short:      "Add a key-value to database",
	Long:       "Add a value by input to database, linked to the provided key.",
	Example: `
	2clip add <key>
	2clip add [ARG] <key>
	`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]
		var input string

		db, _ := database.OpenDatabase("2clip.db", "2clip")
		defer db.Close()

		if (cmd.Flags().Changed("protected")) || (cmd.Flags().Changed("p")) {
			key = key + " (protected)"
		}

		key, err := overwrite(db, key)
		if err != nil {
			log.Printf("overwrite failed: %v", err)
			os.Exit(0)
		}

		if (cmd.Flags().NFlag() == 0) || ((cmd.Flags().NFlag() == 1 && cmd.Flags().Changed("protected")) || (cmd.Flags().Changed("p"))) {
			input = GetInput()
		}
		if cmd.Flags().Changed("extended") || cmd.Flags().Changed("e") {
			input = GetExtendedInput()
		}

		addToDatabase(key, input, db, "2clip")
	},
}

func AddCmdFlags() {
	AddCmd.Flags().BoolP("protected", "p", false, "Add a protected value to the database")
	AddCmd.Flags().BoolP("extended", "e", false, "Add a multiline value to the database using the END word as delimiter")
}

func addToDatabase(key string, value string, db *bolt.DB, bucketName string) {
	err := db.Update(func(tx *bolt.Tx) error {
		bucket, err := util.ConnectToBucket(tx, bucketName)
		if err != nil {
			return err
		}

		err = bucket.Put([]byte(key), []byte(value))
		if err != nil {
			return err
		}

		if strings.HasSuffix(key, " (protected)") {
			fmt.Printf("Added \033[94m"+"%s"+"\033[0m 🔒 with protected value\n", key[:len(key)-12])
		} else {
			fmt.Printf("Added \033[33m"+"%s"+"\033[0m successfully\n", key)
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
	fmt.Printf("key \033[33m"+"%s"+"\033[0m already exists, you want to overwrite it? [Y/N]: ", key)

	util.AnswerCondition()
}

func GetInput() string {
	fmt.Println("Input value:")
	reader := bufio.NewReader(os.Stdin)
	value, _ := reader.ReadString('\n')
	value = strings.TrimSpace(value)

	return value
}
