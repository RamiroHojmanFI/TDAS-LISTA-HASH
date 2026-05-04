package main

import (
	"tdas/pila"
)

func prioridad(caracter rune) int {
	switch caracter {
	case '+':
		return 1

	case '-':
		return 1

	case '*':
		return 2

	case '/':
		return 2

	case '^':
		return 3

	}
	return 0
}

func manejoDeSignos(signos pila.Pila[rune], caracter rune, postfija *string) {
	for !signos.EstaVacia() && (prioridad(signos.VerTope()) >= prioridad(caracter) && caracter != '^') {
		*postfija = *postfija + string(signos.Desapilar()) + " "
	}
	signos.Apilar(caracter)
}
