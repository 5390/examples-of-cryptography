package keyexchangealgorithms

import (
	"crypto/ecdh"
	"crypto/rand"
	"fmt"
)

func KeyExchangeExampleECDH() {
	// Generate private/public key pair for party A
	privateKeyA, _ := ecdh.P256().GenerateKey(rand.Reader)
	publicKeyA := privateKeyA.PublicKey()

	// Generate private/public key pair for party B
	privateKeyB, _ := ecdh.P256().GenerateKey(rand.Reader)
	publicKeyB := privateKeyB.PublicKey()

	// Derive shared secrets
	sharedSecretA, _ := privateKeyA.ECDH(publicKeyB)
	sharedSecretB, _ := privateKeyB.ECDH(publicKeyA)

	fmt.Printf("Shared Secret A: %x\n", sharedSecretA)
	fmt.Printf("Shared Secret B: %x\n", sharedSecretB)
}
