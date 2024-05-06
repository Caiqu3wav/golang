package main

import ("fmt")

func main(){
	numbers := []int{4, 2, 2}
	result := SquareSum(numbers)
	fmt.Println("A soma dos quadrados é:", result)  
}

func SquareSum(number []int) int {

	sum := 0
	for _, num := range number {
		sum += num * num 
	}
	return sum
}
