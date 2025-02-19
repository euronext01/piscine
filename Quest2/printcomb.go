package main

import (
	"github.com/01-edu/z01"
)

func main() {
	for i := '0'; i <= '7'; i++ {
		for j := '1'; j <= '8'; j++ {
			for z := '2'; z <= '9'; z++ {
				z01.PrintRune(i)
				z01.PrintRune(j)
				z01.PrintRune(z)
				if i != '7' || j != '8' || z != '9' {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
}

// 1 hadou code dv zb
