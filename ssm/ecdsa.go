package ssm

import (
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
)

func CreateECDSAKeyPair() (string, string, string, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Error("generate key fail", "err", err)
		return EmptyHexString, EmptyHexString, EmptyHexString, err
	}
	priKeyStr := hex.EncodeToString(crypto.FromECDSA(privateKey))
	pubKeyStr := hex.EncodeToString(crypto.FromECDSAPub(&privateKey.PublicKey))
	compressPubkeyStr := hex.EncodeToString(crypto.CompressPubkey(&privateKey.PublicKey))

	return priKeyStr, pubKeyStr, compressPubkeyStr, nil
}

func SignECDSAMessage(privKey string, txMsg string) (string, error) {
	hash := common.HexToHash(txMsg)
	fmt.Println(hash.Hex())
	privByte, err := hex.DecodeString(privKey)
	if err != nil {
		log.Error("decode private key fail", "err", err)
		return EmptyHexString, err
	}
	privKeyEcdsa, err := crypto.ToECDSA(privByte)
	if err != nil {
		log.Error("Byte private key to ecdsa key fail", "err", err)
		return EmptyHexString, err
	}
	signatureByte, err := crypto.Sign(hash[:], privKeyEcdsa)
	if err != nil {
		log.Error("sign transaction fail", "err", err)
		return EmptyHexString, err
	}
	return hex.EncodeToString(signatureByte), nil
}

func VerifyEcdsaSignature(publicKey, txHash, signature string) (bool, error) {
	// Convert public key from hexadecimal to bytes
	pubKeyBytes, err := hex.DecodeString(publicKey)
	if err != nil {
		log.Error("Error converting public key to bytes", err)
		return false, err
	}

	// Convert transaction string from hexadecimal to bytes
	txHashBytes, err := hex.DecodeString(txHash)
	if err != nil {
		log.Error("Error converting transaction hash to bytes", err)
		return false, err
	}

	// Convert signature from hexadecimal to bytes
	sigBytes, err := hex.DecodeString(signature)
	if err != nil {
		log.Error("Error converting signature to bytes", err)
		return false, err
	}

	// Verify the transaction signature using the public key
	return crypto.VerifySignature(pubKeyBytes, txHashBytes, sigBytes[:64]), nil
}
