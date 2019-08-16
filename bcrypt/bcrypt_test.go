package bcrypt

import "testing"

func TestGenerateAndCheckHash(t *testing.T) {
	hash := NewBcryptHash(12)
	s := "password"
	hashed, err := hash.Make(s)
	if err != nil {
		t.Errorf("error when make a string hash: %v", err)
	}

	isEq := hash.Check(s, hashed)
	if !isEq {
		t.Errorf("string hashed %s is not equal from %s", hashed, s)
	}
}

func TestGenerateAndCheckHashUsingDefault(t *testing.T) {
	hash := DefaultBcryptHash()
	s := "password"
	hashed, err := hash.Make(s)
	if err != nil {
		t.Errorf("error when make a string hash: %v", err)
	}

	isEq := hash.Check(s, hashed)
	if !isEq {
		t.Errorf("hashed string %s is not equal from %s", hashed, s)
	}
}
