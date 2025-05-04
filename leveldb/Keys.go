package leveldb

import "github.com/ethereum/go-ethereum/log"

type Keys struct {
	db *LevelStore
}

func NewKeyStore(path string) (*Keys, error) {
	db, err := NewLevelStore(path)
	if err != nil {
		log.Error("Could not open leveldb", "path", path, "err", err)
		return nil, err
	}
	return &Keys{
		db: db,
	}, nil
}

// GetPrivateKey get privateKey from levelDB
func (keys *Keys) GetPrivateKey(publicKey string) (string, bool) {
	key := []byte(publicKey)
	data, err := keys.db.Get(key)
	if err != nil {
		return "0x00", false
	}
	byteToString := toString(data)
	return byteToString, true
}

// StoreKeys store privateKeys
func (keys *Keys) StoreKeys(keyList []Key) bool {
	for _, item := range keyList {
		key := []byte(item.PublicKey)
		value := toBytes(item.PrivateKey)
		err := keys.db.Put(key, value)
		if err != nil {
			log.Error("Could not store key", "publicKey", item.PublicKey, "err", err)
			return false
		}
	}
}
