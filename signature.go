package main

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"os"
)

type SignatureReq struct {
	Key  string `json:"key"`
	File string `json:"file"`
}

func Signature(req *SignatureReq) string {
	key, err := os.ReadFile(req.Key)
	if err != nil {
		return "error#" + err.Error()
	}
	block, _ := pem.Decode(key)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return "error#failed to decode PEM block"
	}

	// 解析私钥
	private, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "error#" + err.Error()
	}

	// 读取文件
	file, err := os.ReadFile(req.File)
	if err != nil {
		return "error#" + err.Error()
	}
	hashed := sha256.Sum256(file)

	// 计算摘要
	signature, err := rsa.SignPKCS1v15(nil, private, crypto.SHA256, hashed[:])
	if err != nil {
		return "error#" + err.Error()
	}
	return encode(signature)
}
