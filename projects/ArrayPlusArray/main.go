package main

import ("fmt")

func main(){
	var firstArr = []int{1, 2, 3}
	var secondArr = []int{4, 5, 6}

	fmt.Println(arrayPlusArray(firstArr, secondArr))
}

func arrayPlusArray(arr1, arr2 []int) int{
	sum := 0
	combinedArray := append(arr1, arr2...)

	for i := 0; i < len(combinedArray); i++ {
		sum += combinedArray[i]
	}
	return sum
}