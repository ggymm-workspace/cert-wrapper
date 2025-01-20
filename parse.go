package main

import (
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/asn1"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"os"
)

type ParseReq struct {
	Cert string `json:"cert"`
}

func Parse(req *ParseReq) string {
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

	// 计算 issuerNameHash
	name, err := asn1.Marshal(certificate.Issuer.ToRDNSequence())
	if err != nil {
		return "error#" + err.Error()
	}
	nameHex := sha256.Sum256(name)

	// 计算 issuerKeyHash
	public, ok := certificate.PublicKey.(*rsa.PublicKey)
	if !ok {
		return "error#public key is not RSA"
	}
	keyHex := sha256.Sum256(x509.MarshalPKCS1PublicKey(public))

	// 返回结果
	return fmt.Sprintf("success#%s|%s|%s|%x",
		"SHA256",
		hex.EncodeToString(nameHex[:]),
		hex.EncodeToString(keyHex[:]),
		certificate.SerialNumber,
	)
}
