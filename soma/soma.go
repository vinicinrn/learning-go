package main

func Soma(numeros []int) int {
	soma := 0
	for _, numeros := range numeros { //sempre que é chamado, retorna dois valores: indíce e valor
		soma += numeros
	}
	return soma
}
