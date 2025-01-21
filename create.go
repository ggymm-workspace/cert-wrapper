package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"os"
	"path/filepath"
	"time"
)

type CreateReq struct {
	Output string `json:"output"`

	Country      string `json:"country"`
	Organization string `json:"organization"`
	CommonName   string `json:"common_name"`
}

func Create(req *CreateReq) string {
	// 生成私钥
	private, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return "error#" + err.Error()
	}

	// 生成序列号
	sn, err := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 128))
	if err != nil {
		return "error#" + err.Error()
	}

	// 创建模板
	template := &x509.Certificate{
		SerialNumber: sn, // 证书序列号
		Subject: pkix.Name{
			Country:      []string{req.Country},
			Organization: []string{req.Organization},
			CommonName:   req.CommonName,
		},
		NotBefore:             time.Now(),                                 // 证书有效期开始时间
		NotAfter:              time.Now().Add(100 * 365 * 24 * time.Hour), // 证书有效期结束时间（100年）
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	// 自签名证书
	certificate, err := x509.CreateCertificate(rand.Reader, template, template, &private.PublicKey, private)
	if err != nil {
		return "error#" + err.Error()
	}

	// 保存私钥
	key, err := os.Create(filepath.Join(req.Output, "private_key.pem"))
	if err != nil {
		return "error#" + err.Error()
	}
	defer func() {
		_ = key.Close()
	}()
	err = pem.Encode(key, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(private)})
	if err != nil {
		return "error#" + err.Error()
	}

	// 保存证书
	cert, err := os.Create(filepath.Join(req.Output, "certificate.pem"))
	if err != nil {
		return "error#" + err.Error()
	}
	defer func() {
		_ = cert.Close()
	}()
	err = pem.Encode(cert, &pem.Block{Type: "CERTIFICATE", Bytes: certificate})
	if err != nil {
		return "error#" + err.Error()
	}
	return "success"
}
