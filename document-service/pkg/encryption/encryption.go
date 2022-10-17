package encryption

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
)

var (
	ErrWrongIVSize = errors.New("iv length must equal block size")
)

const BlockSize = aes.BlockSize

// Encrypt "text" string with AES
//
// "key" - should be the AES key, either 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256.
//
// "iv" - len(iv) should equal BlockSize (default: 16). IV stands for Initialization vector.
// IV should be randomly generated for each encryption.
// IV are used to ensure that the same value encrypted N times, will not always result in the same encrypted value.
func EncryptAES(key []byte, iv []byte, text string) (string, error) {
	if BlockSize != len(iv) {
		return "", ErrWrongIVSize
	}

	bytesText := PKCS5Padding([]byte(text))
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, len(bytesText))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, bytesText)
	return hex.EncodeToString(ciphertext), nil
}

// Decrypt "cipherText" string with AES
//
// "key" - should be the AES key, either 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256.
//
// "iv" - len(iv) should equal BlockSize (default: 16). IV stands for Initialization vector.
// IV should be randomly generated for each encryption.
// IV are used to ensure that the same value encrypted N times, will not always result in the same encrypted value.
func DecryptAES(key []byte, iv []byte, cipherText string) (string, error) {
	if BlockSize != len(iv) {
		return "", ErrWrongIVSize
	}

	cipherTextDecoded, err := hex.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks([]byte(cipherTextDecoded), []byte(cipherTextDecoded))
	return string(PKCS5UnPadding(cipherTextDecoded)), nil
}

// Add padding bytes for the message to make it
// divisible by the block size before encryption.
func PKCS5Padding(src []byte) []byte {
	padding := BlockSize - len(src)%BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

// Remove padding bytes (usually after decode)
func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}

// Generate pseudorandom IV bytes array
func GenerateRandomIV(length uint) ([]byte, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)

	if err != nil {
		return nil, err
	}

	return b, nil
}
