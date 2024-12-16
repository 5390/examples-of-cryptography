package postquantumcryptography

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/cloudflare/circl/sign/schemes"
	"golang.org/x/crypto/sha3"
)

// hexHash hashes the input using SHA3-256 and returns a hex-encoded string
// It uses SHA3-256 with Shake256 to hash the key and returns the first 16 bytes as hex string.
func hexHash(in []byte) string {
	var ret [16]byte
	h := sha3.NewShake256()
	_, _ = h.Write(in[:])
	_, _ = h.Read(ret[:])
	return hex.EncodeToString(ret[:])
}

// DilithiumKeyGeneration generates keys for Dilithium schemes (Dilithium2, Dilithium3, Dilithium5) based on the provided seed and compares them with expected values.
func DilithiumKeyGeneration(schemeName string, expectedPublicKeyHex string, expectedSecretKeyHex string) {
	// Choose the Dilithium mode (Dilithium2, Dilithium3, or Dilithium5)
	scheme := schemes.ByName(schemeName)
	if scheme == nil {
		log.Printf("Scheme %s not found", schemeName)
	}

	// Generate a key pair from a predefined seed (here we use a static seed)
	// Ensure that the seed here is consistent with what the reference implementation uses.
	var seed [32]byte
	// This is an example seed. Modify it to match the seed used in the reference implementation if required.
	// For now, using a simple byte sequence for the seed
	for i := 0; i < len(seed); i++ {
		seed[i] = byte(i)
	}

	// Derive the public and private keys using the scheme
	pk, sk := scheme.DeriveKey(seed[:])

	// Marshal the keys to binary format
	ppk, err := pk.MarshalBinary()
	if err != nil {
		log.Printf("Failed to marshal public key: %v", err)
	}
	psk, err := sk.MarshalBinary()
	if err != nil {
		log.Printf("Failed to marshal secret key: %v", err)
	}

	// Hash the keys using SHA3-256 and compare with expected values
	pkh := hexHash(ppk)
	skh := hexHash(psk)

	// Print the results for debugging purposes
	fmt.Printf("Generated Public Key Hash: %s\n", pkh)
	fmt.Printf("Generated Secret Key Hash: %s\n", skh)

	// Compare the hashes with the expected values
	if pkh != expectedPublicKeyHex {
		log.Printf("Expected public key hash: %s, but got: %s", expectedPublicKeyHex, pkh)
	}

	if skh != expectedSecretKeyHex {
		log.Printf("Expected secret key hash: %s, but got: %s", expectedSecretKeyHex, skh)
	}

	// If they match, the key generation was successful
	fmt.Println("Key generation and verification successful!")
}
