package main

import (
	"tdas/cola"
	"tdas/pila"
)

func manejoDeParentesis(numeros cola.Cola[rune], signos pila.Pila[rune], postfija string) string {
	if !numeros.EstaVacia() {
		postfija = vaciarNumeros(numeros, postfija) + " "
	}
	postfija = vaciarSignos(signos, postfija)
	if !signos.EstaVacia() && signos.VerTope() == '(' {
		signos.Desapilar()
	}
	return postfija
}
