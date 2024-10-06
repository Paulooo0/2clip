package util

import (
	"sort"
	"strings"

	"github.com/boltdb/bolt"
)

func GetSortedKeys(bucket *bolt.Bucket) ([]string, error) {
	var keys []string
	err := bucket.ForEach(func(k, v []byte) error {
		keys = append(keys, string(k))
		return nil
	})
	if err != nil {
		return nil, err
	}

	sort.Slice(keys, func(i, j int) bool {
		return strings.ToLower(keys[i]) < strings.ToLower(keys[j])
	})

	return keys, nil
}
