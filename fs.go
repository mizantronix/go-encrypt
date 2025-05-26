package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func EncryptDirectory(input, output string, key []byte) error {
	manifest := make(Manifest)
	counter := 0

	err := filepath.Walk(input, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if path == input {
			return nil
		}

		relPath, _ := filepath.Rel(input, path)
		shortName := fmt.Sprintf("%08d", counter)

		manifest[shortName] = relPath
		counter++

		subdir := shortName[:4]

		encryptedPath := filepath.Join(output, subdir, shortName)

		if info.IsDir() {
			return os.MkdirAll(encryptedPath, 0700)
		} else {
			return EncryptFile(path, encryptedPath, key)
		}
	})

	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(manifest, "", " ")
	if err != nil {
		return err
	}

	manifestPath := filepath.Join(output, "MANIFEST")
	fmt.Printf("manifest content: %s \r\n", data)
	return EncryptBytesToFile(data, manifestPath, key)
}

func DecryptDirectory(input, output string, key []byte) error {
	manifestPath := filepath.Join(input, "MANIFEST")
	data, err := DecryptFileToBytes(manifestPath, key)
	if err != nil {
		return err
	}

	var manifest Manifest
	if err := json.Unmarshal(data, &manifest); err != nil {
		return nil
	}

	for shortName, relPath := range manifest {
		subdir := shortName[:4]
		inputPath := filepath.Join(input, subdir, shortName)
		outPath := filepath.Join(output, relPath)

		info, err := os.Stat(inputPath)
		if err != nil {
			return err
		}

		if info.IsDir() {
			err = os.MkdirAll(outPath, 0700)
		} else {
			err = DecryptFile(inputPath, outPath, key)
		}

		if err != nil {
			return err
		}
	}

	return nil
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
