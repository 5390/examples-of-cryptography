package symmetrickey

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

func AESExample() bool {
	key, err := GenerateAESKey()

	if err != nil {
		fmt.Errorf("Got Error while gerating Key :", err)
		return true
	}

	plainText := "Hello, Golang Cryptography!"

	cipherText, err := EncryptAES(key, plainText)
	if err != nil {
		panic(err)
	}
	fmt.Println("Encrypted Text:", cipherText)

	decryptedText, err := DecryptAES(key, cipherText)
	if err != nil {
		panic(err)
	}
	fmt.Println("Decrypted Text:", decryptedText)
	return false
}

func EncryptAES(key []byte, text string) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	cipherText := make([]byte, aes.BlockSize+len(text))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], []byte(text))

	return base64.URLEncoding.EncodeToString(cipherText), nil
}

func DecryptAES(key []byte, cipherText string) (string, error) {
	data, err := base64.URLEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	iv := data[:aes.BlockSize]
	data = data[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(data, data)

	return string(data), nil
}

func GenerateAESKey() ([]byte, error) {
	key := make([]byte, 32) // AES-256
	_, err := rand.Read(key)
	return key, err
}
