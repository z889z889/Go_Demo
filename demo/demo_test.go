package demo_test

import (
	"testing"

	"example.com/go-demo-1/demo"
)

func TestTell(t *testing.T) {
	if demo.Tell() != "me" {
		t.Fatal("Wrong one")
	}
}
