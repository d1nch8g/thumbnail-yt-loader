package cache

import (
	"os"

	"github.com/syndtr/goleveldb/leveldb"
)

var db *leveldb.DB

func init() {
	if _, err := os.Stat("cache"); err != nil {
		os.Mkdir("cache", os.ModePerm)
	}
	database, err := leveldb.OpenFile("cache/db.ldb", nil)
	if err != nil {
		panic(err)
	}
	db = database
}

func CheckThumbnail(uuid string) ([]byte, error) {
	return db.Get([]byte(uuid), nil)
}

func StoreThumbnail(uuid string, content []byte) error {
	return db.Put([]byte(uuid), content, nil)
}
