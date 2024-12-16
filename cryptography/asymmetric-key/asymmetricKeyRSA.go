package asymmetrickey

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

func RSAExample() {
	// Generate RSA Key Pair
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)

	// Encode Private Key to PEM
	privatePEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})

	fmt.Println("Generated RSA Private Key:")
	fmt.Println(string(privatePEM))

	// Public Key
	publicKey := &privateKey.PublicKey

	// Encrypt a message using the public key
	message := []byte("Hello, RSA Encryption!")
	cipherText, _ := rsa.EncryptPKCS1v15(rand.Reader, publicKey, message)
	fmt.Printf("Encrypted Message: %x\n", cipherText)

	// Decrypt the message using the private key
	decryptedMessage, _ := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherText)
	fmt.Printf("Decrypted Message: %s\n", decryptedMessage)
}
