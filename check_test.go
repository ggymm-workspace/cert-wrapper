package main

import (
	"testing"
)

func Test_Check(t *testing.T) {
	req := &CheckReq{
		Cert: "cert/certificate.pem",
	}
	resp := Check(req)
	t.Logf("resp: %v", resp)
}
