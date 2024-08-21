package main

import (
	"fmt"
	"strings"
	"regexp"
	"strconv"
)

func Order(sentence string) string {
	stringsArr := strings.Fields(sentence)
	 quantity := len(stringsArr)
	 orderedString := make([]string, quantity)

	for _, v := range stringsArr {
		re := regexp.MustCompile(`\d+`)
		stringNumber := re.FindString(v)
		num, err := strconv.Atoi(stringNumber)

		if err != nil {
			fmt.Println("Erro ao converter:", err)
			continue
		}
		orderedString[num-1] = v
	}
	return strings.Join(orderedString, " ")
}


func main() {
	text := "is2 Thi1s T4est 3a"
	fmt.Println(Order(text))
}