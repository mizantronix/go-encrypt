package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func EncryptFile(inputPath, outputPath string, key []byte) error {
	plaintext, err := os.ReadFile(inputPath)
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath.Dir(outputPath), 0700)
	if err != nil {
		return err
	}

	return EncryptBytesToFile(plaintext, outputPath, key)
}

func DecryptFile(inputPath, outputPath string, key []byte) error {
	err := os.MkdirAll(filepath.Dir(outputPath), 0700)
	if err != nil {
		return err
	}

	plaintext, err := DecryptFileToBytes(inputPath, key)

	if err != nil {
		return err
	}

	return os.WriteFile(outputPath, plaintext, 0600)
}

func DeriveKey(password string) []byte {
	return Sha256([]byte(password))
}

func Sha256(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}


func EncryptBytesToFile(data []byte, outputPath string, key []byte) error {
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return err
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return os.WriteFile(outputPath, ciphertext, 0600)
}


func DecryptFileToBytes(inputPath string, key []byte) ([]byte, error) {
	ciphertext, err := os.ReadFile(inputPath)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	return gcm.Open(nil, nonce, ciphertext, nil)
}