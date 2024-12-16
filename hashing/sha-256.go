package hashing

import (
	"crypto/sha256"
	"fmt"
)

func Sha256Example() {
	data := "Hello, Hashing with SHA-256!"
	hash := sha256.Sum256([]byte(data))

	fmt.Printf("Original Text: %s\n", data)
	fmt.Printf("SHA-256 Hash: %x\n", hash)
}
