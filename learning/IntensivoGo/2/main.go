package main

import "fmt"

func count(tipo string) {
	for i := 0; i < 10; i++ {
		fmt.Println("%s: %d\n, ", tipo, i)
	}
}

func main() {
	canal := make(chan string)

	go processa(canal)
	result := <-canal
	fmt.Println(result)
}

func processa(canal chan string)  {
	canal <- "processando"
}