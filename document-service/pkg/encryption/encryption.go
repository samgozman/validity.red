package encryption

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
)

var (
	ErrWrongIVSize = errors.New("iv length must equal block size")
)

// Encrypt "text" string with AES
//
// "key" - should be the AES key, either 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256.
//
// "iv" - len(iv) should equal aes.BlockSize (default: 16). IV stads for Initialization vector.
// IV should be randomly generated for each encryption.
// IV are used to ensure that the same value encrypted N times, will not always result in the same encrypted value.
func EncryptAES(key []byte, iv []byte, text string) (string, error) {
	bytesText := PKCS5Padding([]byte(text), len(text))
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	if block.BlockSize() != len(iv) {
		return "", ErrWrongIVSize
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
// "iv" - len(iv) should equal aes.BlockSize (default: 16). IV stads for Initialization vector.
// IV should be randomly generated for each encryption.
// IV are used to ensure that the same value encrypted N times, will not always result in the same encrypted value.
func DecryptAES(key []byte, iv []byte, cipherText string) (string, error) {
	cipherTextDecoded, err := hex.DecodeString(cipherText)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	if block.BlockSize() != len(iv) {
		return "", ErrWrongIVSize
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks([]byte(cipherTextDecoded), []byte(cipherTextDecoded))
	return string(PKCS5UnPadding(cipherTextDecoded)), nil
}

// Add padding bytes for the message to transform
// it into multiple 8-byte blocks.
func PKCS5Padding(src []byte, after int) []byte {
	padding := aes.BlockSize - len(src)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

// Remove padding bytes (usually after decode)
func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}
