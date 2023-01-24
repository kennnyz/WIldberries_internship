package main

import "testing"

func TestPop(t *testing.T) {
	mytest := [7]string{"Изба", "бАзи", "пЯткА", "тЯпка", "слиток", "столик", "листок"}
	expected := map[string][]string{}
	expected["изба"] = []string{"изба", "бази"}
	expected["пятка"] = []string{"пятка", "тяпка"}
	expected["слиток"] = []string{"слиток", "столик", "листок"}
	if Pop(&mytest) != &expected {
		t.Errorf("Expected %v but got %v", &expected, Pop(&mytest))
	}
}
