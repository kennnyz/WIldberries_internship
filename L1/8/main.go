package main

import (
	"fmt"
	"strconv"
)

func IntToBits(num int) string {
	return strconv.FormatInt(int64(num), 2)
}

func setBit(n int, pos int) int {
	pos = len(IntToBits(n)) - pos
	mask := (1 << pos)
	n = n | mask
	return n
}

// Clears the bit at pos in n.
func clearBit(n int, pos int) int {
	pos = len(IntToBits(n)) - pos
	mask := ^(1 << pos)
	n &= mask
	return n
}

func main() {
	fmt.Println(clearBit(823, 2)) // returns 567
	fmt.Println(setBit(140, 2))   // 204
}
