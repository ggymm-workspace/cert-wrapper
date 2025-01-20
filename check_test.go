package main

import (
	"testing"
)

func Test_Check(t *testing.T) {
	req := &CheckReq{
		Filepath: "cert/certificate.pem",
	}
	resp := Check(req)
	t.Logf("resp: %v", resp)
}
