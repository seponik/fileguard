package fileguard

import "github.com/seponik/fileguard/internal/crypto"

// EncryptFile encrypts the give file using AES-GCM with the given key.
// Returns error if encryption fails.
// Creates a file with the same name and path, but with a .fg extension.
func EncryptFile(filePath, key string) error {
	return crypto.EncryptFile(filePath, key)
}

// DecryptFile decrypts the give file using AES-GCM with the given key.
// Returns error if decryption fails.
// Creates a file with the same name and path, but without .fg extension.
func DecryptFile(filePath, key string) error {
	return crypto.DecryptFile(filePath, key)
}
