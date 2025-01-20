package main

import (
	"crypto/rand"
	"math/big"
	"testing"
)

func Test_SerialNumber(t *testing.T) {
	sn, err := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 128))
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%d", sn)
	t.Logf("%x", sn)
}

func Test_Created(t *testing.T) {
	req := &CreateReq{
		Country:      "CN",
		Organization: "CPO",
		CommonName:   "JX_CB_097",
	}
	resp := Create(req)
	t.Logf("resp: %v", resp)
}
