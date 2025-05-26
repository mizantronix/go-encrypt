package main

import (
	"flag"
	"fmt"
	"log"
	"time"
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

	start := time.Now()
	if *mode == "encrypt" {
		err := EncryptDirectory(*inputDir, *outputDir, key)
		if err != nil {
			log.Fatalf("Encryption failed: %v", err)
		}

		t := time.Now()
		fmt.Printf("Encryption complete (%.03fs).", t.Sub(start).Seconds())
	} else if *mode == "decrypt" {
		err := DecryptDirectory(*inputDir, *outputDir, key)
		if err != nil {
			log.Fatalf("Decryption failed: %v", err)
		}

		t := time.Now()
		fmt.Printf("Decryption complete (%.03fs).", t.Sub(start).Seconds())
	} else {
		log.Fatalf("Unknown mode: %s", *mode)
	}
}
