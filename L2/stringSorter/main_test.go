package main

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

var in []string = []string{
	`hello   1 World Muhammed 29232      Oct 21 23:40  tmc.int`,
	`hello   1 World Muhammed 29232      Aug 21 23:40  tmc.int`,
	`salam   1 World Uefa 2023      Sep 21 23:40  tmc.int`,
	`asdd   1 Kumir s 1         aug 26 13:46  доки`,
	`vv     1 Georgiy Saha -1        oct 14 23:13  Games`,
	`asd    1 Conor Mcgregor ilya 0         nov 10 08:04 'UFC'`,
	`asd    1 Habib ilya 0         nov 10 08:04 'задания реп'`,
}

func TestCut(t *testing.T) {
	f, err := os.OpenFile("example.txt", os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range in {
		fmt.Fprintln(f, v)
	}

	f.Close()

	testCases := []struct {
		arg []string
	}{
		{
			arg: []string{"run", "main.go", "example.txt"},
		},
		{
			arg: []string{"run", "main.go", "-k", "5", "-n", "example.txt"},
		},
		{
			arg: []string{"run", "main.go", "-u", "example.txt"},
		},
		{
			arg: []string{"run", "main.go", "-u", "-r", "example.txt"},
		},
	}

	for _, v := range testCases {
		b1, _ := exec.Command("go", v.arg...).Output()
		b2, _ := exec.Command("sort", v.arg[2:]...).Output()
		if string(b1) != string(b2) {
			t.Fatal("\ncut output: ", string(b2), "\nmy cut output: ", string(b1))
		}
	}

	exec.Command("rm", "input").Run()
}
