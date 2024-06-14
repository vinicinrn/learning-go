package main

import "fmt"

const espanhol = "espanhol"
const frances = "frances"
const prefixoOlaPortugues = "Olá, "
const prefixoOlaEspanhol = "Hola, "
const prefixoFrances = "Bonjour, "

func Ola(nome string, idioma string) string {
	if nome == "" {
		nome = "Mundo"
	}

	if idioma == "espanhol" {
		return prefixoOlaEspanhol + nome
	}

	if idioma == "frances" {
		return prefixoFrances + nome
	}

	return prefixoOlaPortugues + nome
}

func main() {
	fmt.Println(Ola("mundo", ""))
}
