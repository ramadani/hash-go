package bcrypt

import (
	hash "github.com/garavan/hash-go"
	"golang.org/x/crypto/bcrypt"
)

type bcryptHash struct {
	cost int
}

// DefaultBcryptHash create bcryptHash using default cost
func DefaultBcryptHash() hash.Hash {
	return &bcryptHash{bcrypt.DefaultCost}
}

// NewBcryptHash create bcryptHash using a given cost
func NewBcryptHash(cost int) hash.Hash {
	return &bcryptHash{cost}
}

func (h *bcryptHash) Make(s string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(s), h.cost)
	return string(bytes), err
}

func (h *bcryptHash) Check(s, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(s))
	return err == nil
}
