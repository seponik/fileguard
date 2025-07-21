package crypto

import (
	"bytes"
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	key := "SuperSecretKey"
	data := []byte("Hey there, I'm Sepehr.")

	encrypted, err := encrypt(data, key)
	if err != nil {
		t.Fatalf("Encryption failed: %v", err)
	}

	decrypted, err := decrypt(encrypted, key)
	if err != nil {
		t.Fatalf("Decryption failed: %v", err)
	}

	if !bytes.Equal(data, decrypted) {
		t.Errorf("Decrypted text doesn't match. Got %s, expected %s", decrypted, data)
	}
}

// func TestEncryptDecryptFile(t *testing.T) {}
