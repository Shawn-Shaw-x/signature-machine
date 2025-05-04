package rpc

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/log"
	"signature-machine/leveldb"
	"signature-machine/protobuf"
	"signature-machine/protobuf/wallet"
	"signature-machine/ssm"
)

// 获取加密支持类型
func (r *RpcServer) GetSupportSignWay(ctx context.Context, request *wallet.SupportSignWayRequest) (*wallet.SupportSignWayResponse, error) {
	var signWay []*wallet.SignWay
	signWay = append(signWay, &wallet.SignWay{Schema: "ecdsa"})
	signWay = append(signWay, &wallet.SignWay{Schema: "eddsa"})
	return &wallet.SupportSignWayResponse{
		Code:    wallet.ReturnCode_SUCCESS,
		Msg:     "get sign way success",
		SignWay: signWay,
	}, nil
}

// 批量导出公钥
func (r *RpcServer) ExportPublicKeyList(ctx context.Context, request *wallet.ExportPublicKeyRequest) (*wallet.ExportPublicKeyResponse, error) {
	resp := &wallet.ExportPublicKeyResponse{
		Code: wallet.ReturnCode_SUCCESS,
	}
	cryptoType, err := protobuf.ParseTransactionType(request.Type)
	if err != nil {
		resp.Msg = "input type error"
		return resp, nil
	}
	if request.Number > 10000 {
		resp.Msg = "input number > 10000"
		return resp, nil
	}

	var keyList []leveldb.Key
	var retKeyList []*wallet.PublicKey

	for counter := 0; counter < int(request.Number); counter++ {
		var privateKeyStr, publicKeyStr, compressPublicKeyStr string
		var err error

		switch cryptoType {
		case protobuf.ECDSA:
			privateKeyStr, publicKeyStr, compressPublicKeyStr, err = ssm.CreateECDSAKeyPair()
		case protobuf.EDDSA:
			privateKeyStr, publicKeyStr, err = ssm.CreateEdDSAKeyPair()
			compressPublicKeyStr = publicKeyStr
		default:
			return nil, errors.New("unsupported key type")
		}
		if err != nil {
			log.Error("create key pair error", "err", err)
			return nil, err
		}

		keyItem := leveldb.Key{
			PrivateKey: privateKeyStr,
			PublicKey:  publicKeyStr,
		}
		putItem := &wallet.PublicKey{
			CompressPubkey: compressPublicKeyStr,
			Pubkey:         publicKeyStr,
		}
		retKeyList = append(retKeyList, putItem)
		keyList = append(keyList, keyItem)
	}
	isOk := r.db.StoreKeys(keyList)
	if !isOk {
		log.Error("db store keys error", "err", isOk)
		return nil, errors.New("db store keys error")
	}
	resp.Code = wallet.ReturnCode_SUCCESS
	resp.Msg = "create key pair success"
	resp.PublicKey = retKeyList
	return resp, nil
}

// 通过 messageHash 进行签名
func (r *RpcServer) SignTxMessage(ctx context.Context, request *wallet.SignTxMessageRequest) (*wallet.SignTxMessageResponse, error) {
	resp := &wallet.SignTxMessageResponse{
		Code: wallet.ReturnCode_ERROR,
	}
	cryptoType, err := protobuf.ParseTransactionType(request.Type)
	if err != nil {
		resp.Msg = "input type error"
		return resp, nil
	}
	privateKey, isOk := r.db.GetPrivateKey(request.PublicKey)
	if !isOk {
		return nil, errors.New("get private key error")
	}

	var signature string
	var err2 error
	switch cryptoType {
	case protobuf.ECDSA:
		signature, err2 = ssm.SignECDSAMessage(privateKey, request.MessageHash)
	case protobuf.EDDSA:
		signature, err2 = ssm.SignEdDSAMessage(privateKey, request.MessageHash)
	default:
		return nil, errors.New("unsupported key type")
	}
	if err2 != nil {
		return nil, err2
	}
	resp.Code = wallet.ReturnCode_SUCCESS
	resp.Msg = "sign tx message success"
	resp.Signature = signature
	return resp, nil
}
