package leveldb

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/log"
	"github.com/syndtr/goleveldb/leveldb"
)

type LevelStore struct {
	*leveldb.DB
}

func NewLevelStore(path string) (*LevelStore, error) {
	handle, err := leveldb.OpenFile(path, nil)
	if err != nil {
		log.Error("open level db file fail", "err", err)
	}
	return &LevelStore{handle}, nil
}

func (db *LevelStore) Put(key []byte, value []byte) error {
	return db.DB.Put(key, value, nil)
}

func (db *LevelStore) Get(key []byte) ([]byte, error) {
	return db.DB.Get(key, nil)
}

func (db *LevelStore) Delete(key []byte) error {
	return db.DB.Delete(key, nil)
}

func toBytes(dataStr string) []byte {
	dataBytes, _ := hex.DecodeString(dataStr)
	return dataBytes
}

func toString(byteArr []byte) string {
	return hex.EncodeToString(byteArr)
}
