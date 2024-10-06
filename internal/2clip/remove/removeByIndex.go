package remove

import (
	"fmt"
	"log"
	"strconv"

	"github.com/Paulooo0/2clip/internal/2clip/util"
	database "github.com/Paulooo0/2clip/internal/database"
	"github.com/boltdb/bolt"
)

func commandRemoveByIndex(index string) {
	db, _ := database.OpenDatabase("2clip.db", "2clip")
	defer db.Close()

	removeValueByIndex(db, index)
}

func removeValueByIndex(db *bolt.DB, index string) {
	err := db.Update(func(tx *bolt.Tx) error {
		bucket, err := util.ConnectToBucket(tx, "2clip")
		if err != nil {
			return err
		}

		idx, err := parseAndValidateIndex(index)
		if err != nil {
			return err
		}

		idx--

		keys, err := util.GetSortedKeys(bucket)
		if err != nil {
			return err
		}

		if idx >= len(keys) {
			return fmt.Errorf("index %d is out of range", idx)
		}

		return deleteByKey(db, bucket, keys[idx])
	})
	if err != nil {
		log.Fatal(err)
	}
}

func parseAndValidateIndex(index string) (int, error) {
	idx, err := strconv.Atoi(index)
	if err != nil {
		return -1, fmt.Errorf("invalid index: %v", err)
	}
	return idx, nil
}
