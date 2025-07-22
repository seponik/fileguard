package fileguard

import (
	"bytes"
	"os"
	"testing"
)

func TestEncryptDecryptFile(t *testing.T) {
	key := "ThisIsMySecretKeyForEncryption!!"

	inputFile := "test_input.txt"
	encryptedFile := "test_input.txt.fg"

	inputData := []byte("Hello, I'm Sepehr.")

	if err := os.WriteFile(inputFile, inputData, 0664); err != nil {
		t.Fatalf("Failed to write input file: %v", err)
	}

	if err := EncryptFile(inputFile, key); err != nil {
		t.Fatalf("Encryption failed: %v", err)
	}

	if err := os.Remove(inputFile); err != nil {
		t.Fatalf("Failed to remove input file: %v", err)
	}

	if err := DecryptFile(encryptedFile, key); err != nil {
		t.Fatalf("Decryption failed: %v", err)
	}

	decryptedData, err := os.ReadFile(inputFile)
	if err != nil {
		t.Fatalf("Failed to read decrypted file: %v", err)
	}

	if !bytes.Equal(decryptedData, inputData) {
		t.Fatalf("Decrypted data does not match original data. Got: %s, Want: %s", decryptedData, inputData)
	}

	os.Remove(inputFile)
	os.Remove(encryptedFile)
}
