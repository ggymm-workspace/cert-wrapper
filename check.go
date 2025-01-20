package main

import (
	"crypto/x509"
	"encoding/pem"
	"os"
)

type CheckReq struct {
	Filepath string `json:"filepath"`
}

func Check(req *CheckReq) string {
	cert, err := os.ReadFile(req.Filepath)
	if err != nil {
		return "error#" + err.Error()
	}
	block, _ := pem.Decode(cert)
	if block == nil || block.Type != "CERTIFICATE" {
		return "error#certificate data is not a valid PEM block"
	}
	certificate, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return "error#" + err.Error()
	}

	// 验证是否过期
	if certificate.NotAfter.Before(certificate.NotBefore) {
		return "error#certificate is expired"
	}

	// 验证证书
	roots := x509.NewCertPool()
	roots.AppendCertsFromPEM(cert)
	_, err = certificate.Verify(x509.VerifyOptions{
		Roots: roots,
	})
	if err != nil {
		return "error#" + err.Error()
	}
	return "success"
}
