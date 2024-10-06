package get

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/Paulooo0/2clip/internal/2clip/util"
	database "github.com/Paulooo0/2clip/internal/database"
	"github.com/boltdb/bolt"
)

func commandGetByIndex(index string) {
	db, _ := database.OpenDatabase("2clip.db", "2clip")
	defer db.Close()
	readIndexValue(db, index)
}

func readIndexValue(db *bolt.DB, index string) {
	err := db.View(func(tx *bolt.Tx) error {
		bucket, err := util.ConnectToBucket(tx, "2clip")
		if err != nil {
			return err
		}

		idx, err := parseAndValidateIndex(index)
		if err != nil {
			return err
		}

		keys, err := getSortedKeys(bucket)
		if err != nil {
			return err
		}

		if idx < 0 || idx >= len(keys) {
			return fmt.Errorf("index out of range: %d", idx+1)
		}

		key := keys[idx]
		return processKeyValue(bucket, key, db)
	})

	if err != nil {
		log.Fatal(err)
	}
}

func parseAndValidateIndex(index string) (int, error) {
	idx, err := strconv.Atoi(index)
	if err != nil {
		return 0, fmt.Errorf("invalid index: %s", index)
	}
	return idx - 1, nil
}

func getSortedKeys(bucket *bolt.Bucket) ([]string, error) {
	var keys []string
	err := bucket.ForEach(func(k, v []byte) error {
		keys = append(keys, string(k))
		return nil
	})
	if err != nil {
		return nil, err
	}
	sort.Strings(keys)
	return keys, nil
}

func processKeyValue(bucket *bolt.Bucket, key string, db *bolt.DB) error {
	value := bucket.Get([]byte(key))
	if value == nil {
		return fmt.Errorf("key '%s' not found", key)
	}

	if strings.HasSuffix(key, " (protected)") {
		err := util.Authenticate(db)
		if err != nil {
			return err
		}
	}

	copyToClipboard(key, value)
	return nil
}
