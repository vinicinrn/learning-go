package main

import "fmt"

const prefixoOlaPortugues = "Olá, " //PRA ESSE CASO, CONST É INÚTIL.

func Ola(nome string) string {
	if nome == "" {
		nome = "Mundo"
	}
	return prefixoOlaPortugues + nome
}

func main() {
	fmt.Println(Ola("mundo"))
}
