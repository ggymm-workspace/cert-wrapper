package main

type VerifyReq struct {
	Key  string `json:"key"`
	File string `json:"file"`
}

func Verify(req *VerifyReq) string {
	return "success"
}
