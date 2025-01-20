package main

import (
	"testing"
)

func Test_Parse(t *testing.T) {
	req := &ParseReq{
		Filepath: "cert/certificate.pem",
	}
	resp := Parse(req)
	t.Logf("resp: %v", resp)
}
