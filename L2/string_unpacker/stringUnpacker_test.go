package main

import "testing"

func TestUnpack(t *testing.T) {
	expected := "aaabcc"
	s, err := UnpackingString("aaabcc")
	if s != expected && err != nil {
		t.Errorf("Expected {%s} but got {%s} ", expected, s)
	}
	s, err = UnpackingString("")

	if s != "" && err != nil {
		t.Errorf("Expected {%s} but got {%s} ", expected, s)
	}

	s, err = UnpackingString("45")

	if s != "" && err != nil {
		t.Errorf("Expected {%s} but got {%s} ", expected, s)
	}

	s, err = UnpackingString("abcd")

	if s != "abcd" && err != nil {
		t.Errorf("Expected {%s} but got {%s} ", expected, s)
	}
	s, err = UnpackingString("pop2opo6k3t3mnmn6")

	if s != "poppopooooookkktttmnmnnnnnn" && err != nil {
		t.Errorf("Expected {%s} but got {%s} ", expected, s)
	}

}
