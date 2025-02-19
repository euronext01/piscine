package main

import (
	"github.com/01-edu/z01"
)

func main() {

	for i := '0'; i <= '7'; i++ {
		for z := '1'; z <= '8'; z++ {
			for l := '2'; l <= '9'; l++ {
				z01.PrintRune(i)
				z01.PrintRune(z)
				z01.PrintRune(l)
				if !(i == '7' && z == '8' && l == '9') {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
}
