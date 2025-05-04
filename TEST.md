## 1.链支持
### request
```bash
grpcurl -plaintext -d '{}' 127.0.0.1:8983 wallet.WalletService.getSupportSignWay
```
### response
```bash
{
  "Code": "SUCCESS",
  "msg": "get sign way success",
  "sign_way": [
    {
      "schema": "ecdsa"
    },
    {
      "schema": "eddsa"
    }
  ]
}
```

## 2.批量地址生成
### request
```bash
grpcurl -plaintext -d '{
  "type": "ecdsa",
  "number": "3"
}' 127.0.0.1:8983 wallet.WalletService.exportPublicKeyList
```
### response
```bash
{
  "Code": "SUCCESS",
  "msg": "create key pair success",
  "public_key": [
    {
      "compress_pubkey": "0287916d6ff3cc4312c8e996bfbff4f8d2f7cf44da480f3a2e70be64bdebaa68f0",
      "pubkey": "0487916d6ff3cc4312c8e996bfbff4f8d2f7cf44da480f3a2e70be64bdebaa68f0e7ca2c7409fc0edb4d5f0496c3a9de1a0904deb48fe329f9a9130e5a70f35e76"
    },
    {
      "compress_pubkey": "031427c8d6062b79072a0b941a87f776f4b862e6ab9f9b7fa5a61f5b7b3722e9ab",
      "pubkey": "041427c8d6062b79072a0b941a87f776f4b862e6ab9f9b7fa5a61f5b7b3722e9ab1b3ec09bf22c123f2398bab7ecab516f7508c25bc0d506c52ee6924403eb701b"
    },
    {
      "compress_pubkey": "03ed79ebca46736aa456729d38d00f87d5c435987adcebc2d6404f4fefc13bb72a",
      "pubkey": "04ed79ebca46736aa456729d38d00f87d5c435987adcebc2d6404f4fefc13bb72a6811f36574c2fcc42c66903c9309a94dadc4d2e74d39b4c58e5979fca189c905"
    }
  ]
}
```

## 3.签名
### request
```bash
grpcurl -plaintext -d '{
  "type": "ecdsa",
  "publicKey": "04ed79ebca46736aa456729d38d00f87d5c435987adcebc2d6404f4fefc13bb72a6811f36574c2fcc42c66903c9309a94dadc4d2e74d39b4c58e5979fca189c905",
  "messageHash": "0x9ca77bd43a45da2399da96159b554bebdd89839eec73a8ff0626abfb2fb4b538"
}' 127.0.0.1:8983 wallet.WalletService.signTxMessage
```
### response
```bash
{
  "Code": "SUCCESS",
  "msg": "sign tx message success",
  "signature": "d703844145ef3277b9c8abf6e878590b69122684b2985bb75fd5454177a9ccfb19efcca6c063ce94ca71f3164d9d83d082fe7df73a6afbde579ff29fce4738e800"
}
```