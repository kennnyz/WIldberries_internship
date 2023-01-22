package main

import (
	"testing"
	"time"
)

func TestGetTheCurrentTime(t *testing.T) {

	input := GetTheCurrentTime()[0:16]

	expectOutput := time.Now().String()[0:16]

	if input != expectOutput {
		t.Errorf("expected output to be %s but got %s ", expectOutput, input)
	}

}
