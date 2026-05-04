package main

import (
	"strings"
	"tdas/cola"
	"tdas/pila"
)

func ConvertirAPostfija(infija string) string {
	numeros := cola.CrearColaEnlazada[rune]()
	signos := pila.CrearPilaDinamica[rune]()
	postfija := ""
	for _, caracter := range infija {
		manejoDeCaracteres(numeros, signos, caracter, &postfija)
		switch caracter {
		case '(':
			signos.Apilar(caracter)
		case ')':
			postfija = manejoDeParentesis(numeros, signos, postfija)
		}
	}
	//Lo retorno parcheando los " "
	return strings.Join(strings.Fields(vacioDeListas(numeros, signos, postfija)), " ")
}
