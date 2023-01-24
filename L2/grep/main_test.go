package main

import (
	"reflect"
	"testing"
)

func TestGrep(t *testing.T) {
	searchStr := "example"
	filename := "text.txt"

	got, err := grep(searchStr, filename)
	if err != nil {
		t.Error("We dont expect error but got: ", err)
	}

	want := []string{
		"I found example line",
		"Here is another one line with example word",
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("We espected %v but got %v ", want, got)
	}
}
