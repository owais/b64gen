package main

import (
	"crypto/rand"
	b64 "encoding/base64"
	"fmt"
	"log"
	"os"
	"strconv"
)

const usage = `
Generates a base64 encoded string from random bytes of a given size

Usage:

	go-rand-key <key_size_in_bytes>

Examples:

	go-rand-key 32
	JP9/jNtHjIxiCHXtg4LSfgk8SlxucDKJu9c/zt2Ee6g=

	go-rand-key 64
	7FrsXoPDJA420b09IssIRisAPrdzY+LOBMd9aknh8E6Ajp4FMYYma/V+rxSVPzhCFX9t5F0wZiCH/p6Sb8bJ2A==
`

func main() {
	if len(os.Args) < 2 {
		fmt.Println(usage)
		return
	}

	size := os.Args[1]
	sizeInt, err := strconv.Atoi(size)
	if err != nil {
		fmt.Println(usage)
		return
	}

	key := make([]byte, sizeInt)
	_, err = rand.Read(key)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(b64.StdEncoding.EncodeToString([]byte(key)))
}
