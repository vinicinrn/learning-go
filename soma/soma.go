package main

func Soma(numeros [5]int) int {
	soma := 0
	for _, numeros := range numeros {
		soma += numeros
	}
	return soma
}
