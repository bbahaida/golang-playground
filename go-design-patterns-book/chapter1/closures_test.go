package main

import "testing"

func TestAddN(t *testing.T) {
	expected := 11
	res := addN(5)(6)
	if res != expected {
		t.Error("result is not correct")
	}
}
