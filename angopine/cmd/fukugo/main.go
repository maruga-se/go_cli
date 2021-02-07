package main

import (
	"angopine"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	gcm, err := angopine.Prepare()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	nonce := make([]byte, 12)
	n, err := io.ReadFull(os.Stdin, nonce)
	if n != 12 || err != nil {
		fmt.Fprintf(os.Stderr, "Can't read nonce: %v\n", err)
		os.Exit(1)
	}
	encrypted, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Decryption Error: %v\n", err)
		os.Exit(1)
	}
	result, err := gcm.Open(nil, nonce, encrypted, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Open Error: %v\n", err)
		os.Exit(1)
	}
	os.Stdout.Write(result)
}
