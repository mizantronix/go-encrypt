package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	mode := flag.String("mode", "encrypt", "Mode: encrypt or decrypt")
	inputDir := flag.String("input", "vault", "Input directory")
	outputDir := flag.String("output", "vault_encrypted", "Output directory")
	password := flag.String("password", "", "Encryption password")
	flag.Parse()

	if *password == "" {
		log.Fatal("Password is required")
	}

	key := DeriveKey(*password)

	if *mode == "encrypt" {
		err := EncryptDirectory(*inputDir, *outputDir, key)
		if err != nil {
			log.Fatalf("Encryption failed: %v", err)
		}

		fmt.Println("Encryption complete.")
	} else if *mode == "decrypt" {
		err := DecryptDirectory(*inputDir, *outputDir, key)
		if err != nil {
			log.Fatalf("Decryption failed: %v", err)
		}

		fmt.Println("Decryption complete.")
	} else {
		log.Fatalf("Unknown mode: %s", *mode)
	}
}
