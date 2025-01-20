package main

import (
	"testing"
)

func Test_Verify(t *testing.T) {
	req := &VerifyReq{
		Cert:      "cert/certificate.pem",
		File:      "cert-wrapper.exe",
		Signature: "Pmo+lZJFlkhVVLtFe9k65RWuFa9hwJMyLTa1lj7OPN1xIkKo3lbkwjQhoUZKQdywu9ZX6j0/b3ETy5yO62QiBjRE/zblG3Ws6894hRSNpB7fEaHJax/QQ5yugzl0wdUyGRMEIOkU1o1evg2Mdkzc2YAu6ArcRDazgEwXQuzg/1hJz09fOlolW1fTNjBg8u0pfBIee5um7vQAJjsTxm4xklV5saK6CL4Deeq8rvYLZikKQVuqKF/hblvufS9frlorL2gCvPxkr2WqOWAESBNNvZOS0ruuUEsGzVX1Yu2+hIGAvS2jRdiZgQ67eelcKONMKBrSzF1CxuOw0op0dKwaxA==",
	}
	resp := Verify(req)
	t.Logf("resp: %v", resp)
}
