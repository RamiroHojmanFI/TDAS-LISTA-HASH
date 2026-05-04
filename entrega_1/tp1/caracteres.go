package main

import (
	"tdas/cola"
	"tdas/pila"
)

func manejoDeCaracteres(numeros cola.Cola[rune], signos pila.Pila[rune], caracter rune, postfija *string) {
	if caracter >= '0' && caracter <= '9' {
		numeros.Encolar(caracter)
	} else if caracter == '+' || caracter == '-' || caracter == '*' || caracter == '/' || caracter == '^' {
		if !numeros.EstaVacia() {
			*postfija = vaciarNumeros(numeros, *postfija) + " "
		}
		manejoDeSignos(signos, caracter, postfija)
	}
}
