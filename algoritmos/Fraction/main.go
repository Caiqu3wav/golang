package main

import "fmt"

func main() {
a:= 8
b:= 12
fmt.Println(Fraction(a, b))
}

func Fraction(a, b int) int {
	rest := Mdc(a, b)

	reducedFraction := []int{a / rest, b / rest}

	result:= reducedFraction[0] + reducedFraction[1]

	return result
}

func Mdc(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}

	var dividendo, divisor, rest int
		
	if a > b {
		dividendo = a
		divisor = b
	} else {
		dividendo = b
		divisor = a
	}

	for dividendo % divisor != 0 {
		rest = dividendo % divisor
		dividendo = divisor
		divisor = rest
	} 
	return divisor
}