package main

import "fmt"

const espanhol = "espanhol"
const frances = "frances"
const ingles = "ingles"
const prefixoOlaPortugues = "Olá, "
const prefixoOlaEspanhol = "Hola, "
const prefixoOlaFrances = "Bonjour, "
const prefixoOlaIngles = "Hi, "

func verificaPrefixo(idioma string) (prefixo string) {
	switch idioma {
	case espanhol:
		prefixo = prefixoOlaEspanhol
	case frances:
		prefixo = prefixoOlaFrances
	case ingles:
		prefixo = prefixoOlaIngles
	default:
		prefixo = prefixoOlaPortugues
	}
	return
}

func Ola(nome string, idioma string) string {
	if nome == "" {
		nome = "Mundo"
	}

	return verificaPrefixo(idioma) + nome
}
func main() {
	fmt.Println(Ola("Jaa", ""))
}
