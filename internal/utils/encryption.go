package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
)

func Encrypt(plaintext string, DefaultKey string) (string, error){
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from panic in Encrypt")
		}
	}()
	bDK := []byte(DefaultKey)
	bIV := []byte("420#abA%,ZfE79@M")
	hexD, _ := hex.DecodeString(plaintext)
	block, _ := aes.NewCipher(bDK)
	ciphertext := make([]byte, len(hexD))
	mode := cipher.NewCBCEncrypter(block, bIV)
	mode.CryptBlocks(ciphertext, hexD)
	return hex.EncodeToString(ciphertext), nil
}

func Decrypt(plaintext []byte, DefaultKey string) (string, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from panic in Decrypt")
		}
	}()
	bDK := []byte(DefaultKey)
	bIV := []byte("420#abA%,ZfE79@M")
	bPlaintext := []byte(plaintext)
	block, _ := aes.NewCipher(bDK)
	ciphertext := make([]byte, len(bPlaintext))
	mode := cipher.NewCBCDecrypter(block, bIV)
	mode.CryptBlocks(ciphertext, bPlaintext)
	fmt.Println("plaintextDed", ciphertext)
	return hex.EncodeToString(ciphertext), nil
}