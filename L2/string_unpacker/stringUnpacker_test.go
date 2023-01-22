package main

import "testing"

func TestUnpack(t *testing.T) {
	expected := "aaabcc"
	if Unpack("a3bc2") != expected {
		t.Errorf("Expected {%s} but got {%s} ", expected, Unpack("a3bc2"))
	}

	if Unpack("") != "" {
		t.Errorf("Expected {%s} but got {%s} ", "", Unpack(""))
	}

	if Unpack("45") != "" {
		t.Errorf("Expected {%s} but got {%s} ", expected, Unpack(""))
	}

	if Unpack("abcd") != "abcd" {
		t.Errorf("Expected {%s} but got {%s} ", expected, Unpack(""))
	}

	if Unpack("pop2opo6k3t3mnmn6") != "poppopooooookkktttmnmnnnnnn" {
		t.Errorf("Expected {%s} but got {%s} ", expected, Unpack(""))
	}
}
