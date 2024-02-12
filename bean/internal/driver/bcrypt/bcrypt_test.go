package bcrypt

import (
	"testing"
)

func TestHasher(t *testing.T) {
	hasher := New()

	hash, err := hasher.Hash("test")
	if err != nil {
		t.Fatalf("error: %s", err)
	}

	if string(hash) == "test" {
		t.Error("hash is not hashed")
	}

	if len(hash) != 60 {
		t.Error("hash length is not 60")
	}
}

func TestCompare(t *testing.T) {
	hasher := New()

	hashed, err := hasher.Hash("test")
	if err != nil {
		t.Fatalf("error: %s", err)
	}

	err = hasher.Compare("test", hashed)
	if err != nil {
		t.Fatalf("error: %s", err)
	}

	err = hasher.Compare(hashed, "wrong")
	if err == nil {
		t.Fatal("error: expected error, got nil")
	}
}
