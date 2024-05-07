package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
)

func Encrypt(plaintext string, DefaultKey string) ([]byte, error){
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from panic in Encrypt", r)
		}
	}()
	bDK := []byte(DefaultKey)
	bIV := []byte("420#abA%,ZfE79@M")
	hexD, _ := hex.DecodeString(plaintext)
	fmt.Println(hexD)

	block, _ := aes.NewCipher(bDK)
	ciphertext := make([]byte, len(hexD))
	mode := cipher.NewCBCEncrypter(block, bIV)
	mode.CryptBlocks(ciphertext, hexD)
	// return hex.EncodeToString(ciphertext), nil
	return ciphertext, nil
}

func EncryptPadding(plaintext string, DefaultKey string) ([]byte, error){
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from panic in Encrypt", r)
		}
	}()
	bDK := []byte(DefaultKey)
	bIV := []byte("420#abA%,ZfE79@M")
	hexD, _ := hex.DecodeString(plaintext)
	hexD = zeroPad(hexD, aes.BlockSize)
	block, _ := aes.NewCipher(bDK)
	ciphertext := make([]byte, len(hexD))
	mode := cipher.NewCBCEncrypter(block, bIV)
	mode.CryptBlocks(ciphertext, hexD)
	// return hex.EncodeToString(ciphertext), nil
	return ciphertext, nil
}

func zeroPad(input []byte, blockSize int) []byte {
	padSize := blockSize - (len(input)  % blockSize)
	padText := bytes.Repeat([]byte{0}, padSize)
	return append(input, padText...)
}



func Decrypt(plaintext []byte, DefaultKey string) (string, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from panic in Decrypt")
		}
	}()
	bDK := []byte(DefaultKey)
	bIV := []byte("420#abA%,ZfE79@M")
	// fmt.Printf("pltxt: %", string(plaintext))
	bPlaintext := []byte(plaintext)
	block, _ := aes.NewCipher(bDK)
	ciphertext := make([]byte, len(bPlaintext))
	mode := cipher.NewCBCDecrypter(block, bIV)
	mode.CryptBlocks(ciphertext, bPlaintext)
	return hex.EncodeToString(ciphertext), nil
}

func DecryptByte(plaintext []byte, DefaultKey string) ([]byte, error) {
	fmt.Println("420plaintext", DefaultKey)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from panic in Decrypt",r)
		}
	}()
	bDK := []byte(DefaultKey)
	bIV := []byte("420#abA%,ZfE79@M")
	bPlaintext := []byte(plaintext)
	block, _ := aes.NewCipher(bDK)
	ciphertext := make([]byte, len(bPlaintext))
	mode := cipher.NewCBCDecrypter(block, bIV)
	fmt.Printf("pltxt123: %s\n", string(plaintext))
	mode.CryptBlocks(ciphertext, bPlaintext)
	return ciphertext, nil
}