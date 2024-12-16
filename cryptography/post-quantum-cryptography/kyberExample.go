package postquantumcryptography

import (
	"fmt"
	"log"

	"github.com/cloudflare/circl/kem/schemes"
)

func KyberExample() {
	// Choose the Kyber mode (Kyber512, Kyber768, or Kyber1024)
	schemeName := "Kyber768" // Use a valid scheme name
	scheme := schemes.ByName(schemeName)

	// Check if the scheme was found
	if scheme == nil {
		log.Printf("Scheme %s not found", schemeName)
	}

	// Key Pair Generation (for the recipient)
	fmt.Println("Generating keypair for recipient...")
	publicKey, privateKey, err := scheme.GenerateKeyPair()
	if err != nil {
		log.Printf("Key generation failed: %v", err)
	}
	fmt.Println("Recipient's keypair generated successfully!")

	// Encapsulation: Sender generates shared secret and ciphertext
	fmt.Println("Sender is encapsulating shared secret...")
	sharedSecretSender, ciphertext, err := scheme.Encapsulate(publicKey)
	if err != nil {
		log.Printf("Encapsulation failed: %v", err)
	}
	fmt.Println("Shared secret encapsulated successfully!")

	// Check the size of the ciphertext to debug the error
	expectedCiphertextSize := scheme.EncapsulationSeedSize() // Correct method to get ciphertext size
	fmt.Printf("Expected Ciphertext Size: %d bytes\n", expectedCiphertextSize)
	fmt.Printf("Actual Ciphertext Size: %d bytes\n", len(ciphertext))

	// Ensure the ciphertext size matches the expected size for the Kyber768 scheme
	if len(ciphertext) != expectedCiphertextSize {
		log.Printf("Ciphertext size mismatch: Expected %d, got %d", expectedCiphertextSize, len(ciphertext))
	}

	// Decapsulation: Recipient retrieves the shared secret using the private key
	fmt.Println("Recipient is decapsulating shared secret...")
	sharedSecretRecipient, err := scheme.Decapsulate(privateKey, ciphertext)
	if err != nil {
		log.Printf("Decapsulation failed: %v", err)
	}
	fmt.Println("Shared secret decapsulated successfully!")

	// Comparing shared secrets
	if string(sharedSecretSender) == string(sharedSecretRecipient) {
		fmt.Println("Shared secrets match! Key exchange successful!")
	} else {
		fmt.Println("Shared secrets do not match! Key exchange failed!")
	}
}
