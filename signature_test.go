package main

import (
	"testing"
)

func Test_Signature(t *testing.T) {
	req := &SignatureReq{
		Key:  "cert/private_key.pem",
		File: "cert-wrapper.exe",
	}
	resp := Signature(req)
	t.Logf("resp: %v", resp)
}
