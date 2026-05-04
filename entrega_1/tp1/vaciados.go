package main

import (
	"tdas/cola"
	"tdas/pila"
)

func vaciarNumeros(numeros cola.Cola[rune], postfija string) string {
	for !numeros.EstaVacia() {
		postfija += string(numeros.Desencolar())
	}
	return postfija
}
func vaciarSignos(signos pila.Pila[rune], postfija string) string {
	for !signos.EstaVacia() && signos.VerTope() != '(' {
		postfija += string(signos.Desapilar()) + " "
	}
	return postfija
}

func vacioDeListas(numeros cola.Cola[rune], signos pila.Pila[rune], postfija string) string {
	if !numeros.EstaVacia() {
		postfija = vaciarNumeros(numeros, postfija) + " "
	}
	postfija = vaciarSignos(signos, postfija)
	return postfija
}
