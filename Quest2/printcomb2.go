package main

import (
	"github.com/01-edu/z01"
)

func main() {

	 for a := 0; a <= 98; a++ { // hada for baxytba3na joj ar9am lwlin 00
		for b := a + 1; b <= 99; b++ { // hadi bax ytba3na joj taniyn  00
			z01.PrintRune(rune((a / 10) + '0')) // had rune bax fixiw beli ga33 les nmbrs lwl fo9 hiya rune +'0' bax nhawllo l3dd ila rune 
			z01.PrintRune(rune((a%10 + '0')))
			z01.PrintRune(rune(' '))
			z01.PrintRune(rune((b / 10) + '0'))
			z01.PrintRune(rune((b % 10) + '0'))
			if !(a == 98 && b == 99) {  //
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}

		}
	}
}
