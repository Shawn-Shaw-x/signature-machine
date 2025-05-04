package ssm

import (
	"crypto/ed25519"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"testing"
)

func TestCreateEdDSAKeyPair(t *testing.T) {
	privKey, pubKey, _ := CreateEdDSAKeyPair()
	fmt.Println("privKey=", privKey)
	fmt.Println("pubKey=", pubKey)
}

func TestSignEdDSAMessage(t *testing.T) {
	privKey := "09fa5c99a11f3857dccfede0b9f6ead29bc2f5757b43b336796d64d2cdacf74a39f523de37c1218d28ca467a6e0ea0aa0a603064ab402983829513a0feca0039"
	txMsg, _ := SignEdDSAMessage(privKey, common.Hash{}.String())
	fmt.Println("txMsg=", txMsg)
}

func TestSignEdDSAMessageV2(t *testing.T) {
	seed := ""
	message := ""

	seedBytes, err := hex.DecodeString(seed)
	if err != nil {
		t.Fatalf("DecodeString: %v", err)
	}

	privateKey := ed25519.NewKeyFromSeed(seedBytes)

	fullPrivateKey := hex.EncodeToString(privateKey)

	signature, err := SignEdDSAMessage(fullPrivateKey, message)

	if err != nil {
		t.Errorf("SignEdDSAMessage() fail: %v", err)
		return
	}

	t.Logf("signature: %s", signature)

	if signature == "" {
		t.Error("SignEdDSAMessage() not nil")
	}

	decodedSig, err := hex.DecodeString(signature)
	if err != nil {
		t.Errorf("signature DecodeString fail: %v", err)
	}
	if len(decodedSig) != 64 {
		t.Errorf("decodedSig len error: got %d, want 64", len(decodedSig))
	}
}

func TestVerifyEdDSASign(t *testing.T) {
	signature := "9594742341865c897b066714f10bc90f4c2afebba986c3fef500d9d69ea8eb6df19a19160c210dcf65becbe68c1885820915565f3a84277efc34027b73c89608"
	pubKey := "39f523de37c1218d28ca467a6e0ea0aa0a603064ab402983829513a0feca0039"
	ok := VerifyEdDSASign(pubKey, common.Hash{}.String(), signature)
	fmt.Println("ok=", ok)
}
