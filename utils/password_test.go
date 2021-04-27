package utils

import (
	"math/rand"
	"testing"
)

// check if hashpass returns a hash
func TestHashPass(t *testing.T) {
	// generate a random 31bit string
	pass := string(rand.Int31())

	result := HashPass(pass)

	if result == "" {
		t.Errorf("Hash of %#v was empty", pass)
		return
	}
	t.Logf("HashPass sucessfully returns something from %#v: %#v", pass, result)
}

// check if CompareHashPass returns the correct result
func TestCompareHashPass(t *testing.T) {
	pass := "hello"

	result := HashPass(pass)

	if ans := CompareHashPass(pass, result); ans {
		t.Logf("CompareHashPass returned the correct result: pass=%#v, hash=%#v", pass, result)
	} else {
		t.Errorf("CompareHashPass didnt return the correct result: pass=%#v, hash=%#v", pass, result)
	}
}
