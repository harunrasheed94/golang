// You can edit this code!
// Click here and start typing.
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

// AesEncryptor is an implementation for EncryptDecryptInterface
type AesEncryptor struct {
	cipher.AEAD
}

// NewAesEncryptor is a convenience function to create a UserController
func NewAesEncryptor(aesSecret []byte) (*AesEncryptor, error) {
	block, err := aes.NewCipher(aesSecret)
	if err != nil {
		return nil, err
	}
	aesGcm, err := cipher.NewGCM(block)
	return &AesEncryptor{aesGcm}, err
}

// EncryptDecryptInterface interface describes all operations
type EncryptDecryptInterface interface {
	Encrypt(clearTxt string) (encryptedTxt, cryptoNonce string, encryptErr error)
	Decrypt(encryptedTxt string, cryptoNonce string) (clearTxt string, decryptErr error)
}

// Encrypt encrypts clearTxt content and returns a hex of it along with a nonce.
func (e AesEncryptor) Encrypt(clearTxt string) (string, string, error) {

	inputContents := []byte(clearTxt)

	nonce := make([]byte, 12)

	_, _ = io.ReadFull(rand.Reader, nonce)

	ciphertext := e.AEAD.Seal(nil, nonce, inputContents, nil)

	return hex.EncodeToString(ciphertext), hex.EncodeToString(nonce), nil
}

// Decrypt takes in hexed encryptedTxt content, nonce and returns in plain text.
func (e AesEncryptor) Decrypt(encryptedTxt string, cryptoNonce string) (string, error) {

	ciphertext, err := hex.DecodeString(encryptedTxt)

	if err != nil {
		return "", err
	}

	nonce, err := hex.DecodeString(cryptoNonce)

	if err != nil {
		return "", err
	}

	plainTextContents, err := e.AEAD.Open(nil, nonce, ciphertext, nil)

	return string(plainTextContents), err
}

func main() {
	password := "FBN9xRSJEZB4lrtEFKxA"
	aesKey := "PHgnyMjkcFzNrmzKe9vT4MZA945WvJWS"
	e, _ := NewAesEncryptor([]byte(aesKey))
	encryptedPwd, nonce, _ := e.Encrypt(password)
	fmt.Println("EncryptedPassword = ", encryptedPwd)
	fmt.Println("Nonce = ", nonce)
	pwd, _ := e.Decrypt(encryptedPwd, nonce)
	fmt.Println("Password = ", pwd)
}
