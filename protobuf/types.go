package protobuf

import "errors"

// CryptoType Define a custom type for cryptographic algorithm types
type CryptoType string

// Define constants for the supported cryptographic types
const (
	ECDSA CryptoType = "ecdsa"
	EDDSA CryptoType = "eddsa"
)

func ParseTransactionType(s string) (CryptoType, error) {
	switch s {
	case string(ECDSA):
		return ECDSA, nil
	case string(EDDSA):
		return EDDSA, nil
	default:
		return "", errors.New("unknown transaction type")
	}
}
