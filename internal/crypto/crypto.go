package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"io"
)

func processKey(key string) []byte {
	hash := sha256.Sum256([]byte(key))
	return hash[:]
}

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
