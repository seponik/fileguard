package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"io"
	"os"
	"strings"
)

// processKey returns a 32-byte hash of the input key to make sure key length is valid.
// This helps if the user gives a key that is too short or too long.
func processKey(key string) []byte {
	hash := sha256.Sum256([]byte(key))
	return hash[:]
}

// encrypt encrypts the give data using AES-GCM with the given key.
// Returns the combined nonce and ciphertext, or an error if encryption fails.
func encrypt(data []byte, key string) ([]byte, error) {
	block, err := aes.NewCipher(processKey(key))
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, data, nil), nil
}

// decrypt decrypts the given data using AES-GCM with given key.
// The input must start with 12-byte nonce followed by the ciphertext.
// Returns the original data or an error if decryption fails.
func decrypt(data []byte, key string) ([]byte, error) {
	if len(data) < 12 {
		return nil, errors.New("decrypt: input data is too short; expected at least 12 bytes for nonce and ciphertext")
	}

	block, err := aes.NewCipher(processKey(key))
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce, ciphertext := data[:12], data[12:]

	return gcm.Open(nil, nonce, ciphertext, nil)
}

// EncryptFile encrypts the given file using encrypt function.
// Returns error if encryption fails.
func EncryptFile(filePath, key string) error {
	originalData, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	encryptedData, err := encrypt(originalData, key)
	if err != nil {
		return err
	}

	// WARNING: Modifying the encrypted file will break the authentication system.
	// Even a small change will cause decryption to fail.
	return os.WriteFile(filePath+".fg", encryptedData, 0444)
}

// DecryptFile decrypts the given file using decrypt function.
// Returns error if decryption fails.
func DecryptFile(filePath, key string) error {
	encryptedData, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	originalData, err := decrypt(encryptedData, key)
	if err != nil {
		return err
	}

	return os.WriteFile(strings.Replace(filePath, ".fg", "", 1), originalData, 0664)
}
