package main

import ("fmt")

func main() {
	var chave = [][]int{{1, 1, 1}, {3, 2, 0}, {1, 1, 0}}
	texto := "eu não quero mais ir pra lá"
	fmt.Println(cifrar(texto, chave))
}

func convertToMatriz(mensagem string){
	matriz := make([][]int, len(mensagem))

	for i := range mensagem {
		matriz[i, 0] := mensagem[i]
	}
}

func cifrar(mensagem string, key [][]int) string{
	mensagemMatriz := convertToMatriz(mensagem)
	mensagemCifrada := multiplicarMatrizes(mensagemMatriz, key)
	return mensagemCifrada
}

func decifrar(mensagemCifrada [][]int, key [][]int) string{
	chaveInversa := calcMatrizInversa(key)
	mensagemDecifradaMatriz := multiplicarMatrizes(mensagemCifrada, key)
	mensagemDecifrada := convertToStr(mensagemDecifradaMatriz, key)
	return mensagemDecifrada
}