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
			fmt.Println("error#", err)
		}
		fmt.Println(Check(req))
	} else if op == "parse" {
		req := new(ParseReq)
		err := json.Unmarshal(args, req)
		if err != nil {
			fmt.Println("error#", err)
		}
		fmt.Println(Parse(req))
	} else if op == "create" {
		req := new(CreateReq)
		err := json.Unmarshal(args, req)
		if err != nil {
			fmt.Println("error#", err)
		}
		fmt.Println(Create(req))
	} else if op == "verify" {
		req := new(VerifyReq)
		err := json.Unmarshal(args, req)
		if err != nil {
			fmt.Println("error#", err)
		}
		fmt.Println(Verify(req))
	} else if op == "signature" {
		req := new(SignatureReq)
		err := json.Unmarshal(args, req)
		if err != nil {
			fmt.Println("error#", err)
		}
		fmt.Println(Signature(req))
	}
}
