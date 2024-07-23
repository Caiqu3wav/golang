package main

import (
	"fmt"
)

func criptografia(texto string, matriz [][]int) string {
	if len(texto) < 15 {
		fmt.Println("Precisa de mais de 15 Caractéres no texto")
	}
	
	byteArray := []byte(texto)

	tamanhoBloco := 3
	blocos := make([][]byte, len(byteArray)/tamanhoBloco)
	for i := 0; i < len(byteArray); i += tamanhoBloco {
		blocos = append(blocos, byteArray[i:i+tamanhoBloco])
	}

	var matrizPermutacao [][]int 
	for _, row := range matriz {
		matrizPermutacao = append(matrizPermutacao, row)
	}

	for i := range blocos {
		bloco := blocos[i]
		row := matrizPermutacao[i % len(matrizPermutacao)]
		for j := range row {
			if j < len(bloco) {
				blocos[i][j] = bloco[row[j]]
			}
		}
	}

	var textoCriptografado []byte 
	for _, bloco := range blocos {
		textoCriptografado = append(textoCriptografado, bloco...)
	}

	return string(textoCriptografado)
}

func main() {
	matriz := [][]int{{1, 1, 1}, {3, 2, 0}, {1, 1, 0}}
	texto := "EUNAOQUEROMAISISSONUNCAMAIS"
	textoCriptografado := criptografia(texto, matriz)
	fmt.Println(textoCriptografado)
}