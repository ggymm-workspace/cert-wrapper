package main

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"os"
)

type VerifyReq struct {
	Cert      string `json:"cert"`
	File      string `json:"file"`
	Signature string `json:"signature"`
}

func Verify(req *VerifyReq) string {
	cert, err := os.ReadFile(req.Cert)
	if err != nil {
		return "error#" + err.Error()
	}

	// 解析证书
	block, _ := pem.Decode(cert)
	if block == nil || block.Type != "CERTIFICATE" {
		return "error#certificate data is not a valid PEM block"
	}
	certificate, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return "error#" + err.Error()
	}

	// 获取公钥
	public, ok := certificate.PublicKey.(*rsa.PublicKey)
	if !ok {
		return "error#public key is not RSA"
	}

	// 读取文件
	file, err := os.ReadFile(req.File)
	if err != nil {
		return "error#" + err.Error()
	}
	hashed := sha256.Sum256(file)

	// 验证签名
	err = rsa.VerifyPKCS1v15(public, crypto.SHA256, hashed[:], decode(req.Signature))
	if err != nil {
		return "error#" + err.Error()
	}
	return "success"
}
