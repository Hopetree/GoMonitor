package tool

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

type AESCipher struct {
	key []byte
}

func NewAESCipher(key string) *AESCipher {
	byteKey := []byte(key)
	return &AESCipher{byteKey}
}

func (c *AESCipher) Encrypt(plaintext string) string {
	block, err := aes.NewCipher(c.key)
	if err != nil {
		panic(err)
	}

	plaintextBytes := []byte(plaintext)
	plaintextBytes = pkcs7Pad(plaintextBytes, aes.BlockSize)

	ciphertext := make([]byte, aes.BlockSize+len(plaintextBytes))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintextBytes)

	encoded := base64.StdEncoding.EncodeToString(ciphertext)
	if len(encoded) > 128 {
		panic("encrypted text too long")
	}

	return encoded
}

func (c *AESCipher) Decrypt(ciphertext string) string {
	ciphertextBytes, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		panic(err)
	}

	block, err := aes.NewCipher(c.key)
	if err != nil {
		panic(err)
	}

	if len(ciphertextBytes) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertextBytes[:aes.BlockSize]
	ciphertextBytes = ciphertextBytes[aes.BlockSize:]

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertextBytes, ciphertextBytes)

	return string(pkcs7Unpad(ciphertextBytes))
}

func pkcs7Pad(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padText...)
}

func pkcs7Unpad(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}
