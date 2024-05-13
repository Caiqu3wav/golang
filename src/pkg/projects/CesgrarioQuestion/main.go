package main

import ("fmt")

func main(){
	var n int = 1
	var x int = 0

	fmt.Println(calc(n, x))
}

func calc(n, x int) int{
	var i int

	for n < 10 {
		if n % 2 != 0 {
			for i = 3; i * i <= n; i += 2{
				if n % i == 0{
				break
			}
			}
			if i < n {
				x++
			}
		}
		n++
	}
	return x + n
}