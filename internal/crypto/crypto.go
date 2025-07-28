package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"os"
	"strings"

	"golang.org/x/crypto/argon2"
)

// fileguardPayload is salt + nonce + ciphertext.

const (
	NonceSize = 12 // Size of the nonce in bytes (used for AES-GCM)
	SaltSize  = 16 // Size of the salt in bytes (used for Argon2id)
)

// generateRandomBytes returns a slice of securely generated random bytes with the specified length.
// Returns an error if the random byte generation fails.
func generateRandomBytes(length int) ([]byte, error) {
	randomBytes := make([]byte, length)

	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}

	return randomBytes, nil
}

// hashKey securely hashes the given key with Argon2id.
// This ensures the key length is valid, even if it’s too short or too long.
// Returns a 32-byte hash or an error if hashing fails.
func hashKey(key string, salt []byte) ([]byte, error) {
	keyHash := argon2.IDKey([]byte(key), salt, 3, 64*1024, 2, 32)

	return keyHash, nil
}

// encrypt encrypts the give data using AES-GCM with the given key.
// Returns the combined salt, nonce ciphertext (fileguardPayload), or an error if encryption fails.
func encrypt(plaintext []byte, key string) ([]byte, error) {
	keySalt, err := generateRandomBytes(SaltSize)
	if err != nil {
		return nil, err
	}

	derivedKey, err := hashKey(key, keySalt)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(derivedKey)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce, err := generateRandomBytes(NonceSize)
	if err != nil {
		return nil, err
	}

	cipherText := gcm.Seal(nil, nonce, plaintext, nil)

	fileguardPayload := append(append(keySalt, nonce...), cipherText...)

	return fileguardPayload, nil
}

// decrypt decrypts the given fileguardPayload (salt + nonce + ciphertext) using AES-GCM with the given key.
// The input must start with a 16-byte salt, followed by a 12-byte nonce, and then the ciphertext.
// Returns the plaintext or an error if decryption fails.
func decrypt(fileguardPayload []byte, key string) ([]byte, error) {
	if len(fileguardPayload) < SaltSize+NonceSize {
		return nil, errors.New("decrypt: input data is too short; expected at least 28 bytes for salt, nonce, and ciphertext")
	}

	salt, nonce, ciphertext := fileguardPayload[:SaltSize], fileguardPayload[SaltSize:SaltSize+NonceSize], fileguardPayload[SaltSize+NonceSize:]

	derivedKey, err := hashKey(key, salt)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(derivedKey)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	return gcm.Open(nil, nonce, ciphertext, nil)
}

// EncryptFile encrypts the given file using encrypt function.
// Returns error if encryption fails.
func EncryptFile(filePath, key string) error {
	plaintext, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	fileguardPayload, err := encrypt(plaintext, key)
	if err != nil {
		return err
	}

	// WARNING: Modifying the encrypted file will break the authentication system.
	// Even a small change will cause decryption to fail.
	return os.WriteFile(filePath+".fg", fileguardPayload, 0444)
}

// DecryptFile decrypts the given file using decrypt function.
// Returns error if decryption fails.
func DecryptFile(filePath, key string) error {
	fileguardPayload, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	plaintext, err := decrypt(fileguardPayload, key)
	if err != nil {
		return err
	}

	return os.WriteFile(strings.Replace(filePath, ".fg", "", 1), plaintext, 0664)
}
