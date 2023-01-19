package facade

import (
	"fmt"
	"testing"
)

func TestFacade(t *testing.T) {
	expect := "Build System Block\nMonitor turn on\nKeyboard connect"
	pc := NewPC()
	res := pc.Start()
	if res != expect {
		fmt.Printf("Expect result to equal %s, but %s.\n\n", expect, res)
	}
}
