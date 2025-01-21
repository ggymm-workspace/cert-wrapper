package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	op := os.Args[1]
	args := decode(os.Args[2])
	if op == "check" {
		req := new(CheckReq)
		err := json.Unmarshal(args, req)
		if err != nil {
			fmt.Print("error#", err)
		}
		fmt.Print(Check(req))
	} else if op == "parse" {
		req := new(ParseReq)
		err := json.Unmarshal(args, req)
		if err != nil {
			fmt.Print("error#", err)
		}
		fmt.Print(Parse(req))
	} else if op == "create" {
		req := new(CreateReq)
		err := json.Unmarshal(args, req)
		if err != nil {
			fmt.Print("error#", err)
		}
		fmt.Print(Create(req))
	} else if op == "verify" {
		req := new(VerifyReq)
		err := json.Unmarshal(args, req)
		if err != nil {
			fmt.Print("error#", err)
		}
		fmt.Print(Verify(req))
	} else if op == "signature" {
		req := new(SignatureReq)
		err := json.Unmarshal(args, req)
		if err != nil {
			fmt.Print("error#", err)
		}
		fmt.Print(Signature(req))
	}
}
