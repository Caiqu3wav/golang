package main

import ("fmt")

func main() {
	var chave =  [][][]int{
	{{1, 1, 1}, {3, 2, 0}, {1, 1, 0}},
	}
	texto := "eu não quero mais ir pra lá"
	fmt.Println(cifrar(texto, chave))
}

func convertToMatriz(mensagem string) [][][]int{
	tamanhoBloco := 3
	tamanhoMatriz := (len(mensagem) + tamanhoBloco - 1) / tamanhoBloco
	matriz := make([][][]int, tamanhoMatriz)

	for i := 0; i < len(mensagem); i += tamanhoBloco {
		fim := i + tamanhoBloco
		if fim > len(mensagem) {
			fim = len(mensagem)
		}
		bloco := mensagem[i:fim]

		matriz[i/tamanhoBloco] = make([][]int, tamanhoBloco)
		for j, char := range bloco {
			matriz[i/tamanhoBloco][j] = []int{int(char)}
		}
	}
	return matriz
}

func multiplicarMatrizes(mensagemMatriz [][][]int, chave[][][]int) [][][]int {
	if len(mensagemMatriz[0]) != len(chave){
        panic("As dimensões das matrizes não são compatíveis para multiplicação")
	}

	resultado := make([][][]int, len(mensagemMatriz))
	for i := range resultado {
		resultado[i] = make([][]int, len(chave[0]))
		for j := range resultado[i]{
		resultado[i][j] = make([]int, len(chave[0][0]))
	}}

	for i, linhaMensagem := range mensagemMatriz {
		for j, colunaChave := range chave[0] {
			for k := range colunaChave {
				for l := range linhaMensagem {
					resultado[i][j][k] += mensagemMatriz[i][l][0] * chave[l][j][k]
				}
			}
		}
	}
	return resultado
}

func cifrar(mensagem string, key [][][]int) [][][]int{
	mensagemMatriz := convertToMatriz(mensagem)
	mensagemCifrada := multiplicarMatrizes(mensagemMatriz, key)
	return mensagemCifrada
}

