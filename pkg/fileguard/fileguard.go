package fileguard

import "github.com/seponik/fileguard/internal/crypto"

func EncryptFile(filePath, key string) error {
	return crypto.EncryptFile(filePath, key)
}

func DecryptFile(filePath, key string) error {
	return crypto.DecryptFile(filePath, key)
}
