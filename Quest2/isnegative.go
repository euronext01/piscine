package main

import "github.com/01-edu/z01"

func IsNegative(nb int) {

	if nb < 0 {
		z01.PrintRune('T')
		//ila kan 9al mn, 0 kitba3 True

	} else {
		z01.PrintRune('F')  // ila kan kbar mn 0 kitbazza3 False
		z01.PrintRune('\n')  // hadi kirja3 lstar
	}
}

func main() {
	IsNegative(1)
	IsNegative(0)
	IsNegative(1)
}
