package main

import (
	"encoding/base64"
	"os"
	"path/filepath"
)

func EncryptDirectory(input, output string, key []byte) error {
	return filepath.Walk(input, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, _ := filepath.Rel(input, path)
		encryptedPath := filepath.Join(output, EncodeName(relPath))

		if info.IsDir() {
			return os.MkdirAll(encryptedPath, 0700)
		} else {
			return EncryptFile(path, encryptedPath, key)
		}
	})
}

func DecryptDirectory(input, output string, key []byte) error {
	return filepath.Walk(input, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, _ := filepath.Rel(input, path)
		decryptedPath := filepath.Join(output, DecodeName(relPath))

		if info.IsDir() {
			return os.MkdirAll(decryptedPath, 0700)
		} else {
			return DecryptFile(path, decryptedPath, key)
		}
	})
}

func EncodeName(name string) string {
	return base64.URLEncoding.EncodeToString([]byte(name))
}

func DecodeName(encoded string) string {
	decoded, err := base64.URLEncoding.DecodeString(encoded)
	if err != nil {
		return encoded
	}
	return string(decoded)
}
