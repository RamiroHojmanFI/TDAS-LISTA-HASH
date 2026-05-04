package cola_test

import (
	"github.com/stretchr/testify/require"
	TDAcola "tdas/cola"
	"testing"
)

func TestCrearCola(t *testing.T) {
	c := TDAcola.CrearColaEnlazada[int]()

	require.True(t, c.EstaVacia())

	require.Panics(t, func() { c.VerPrimero() })
	require.Panics(t, func() { c.Desencolar() })
}

func TestEncolar(t *testing.T) {
	c := TDAcola.CrearColaEnlazada[int]()

	c.Encolar(10)
	c.Encolar(20)
	c.Encolar(30)
	c.Encolar(40)
	c.Encolar(50)

	if c.EstaVacia() {
		t.Error()
	}

	if c.VerPrimero() != 10 {
		t.Error()
	}
}

func TestDesencolar(t *testing.T) {
	c := TDAcola.CrearColaEnlazada[int]()
	c.Encolar(10)
	c.Encolar(20)
	c.Encolar(30)
	c.Encolar(40)
	c.Encolar(50)
	require.Equal(t, 10, c.Desencolar())
	require.Equal(t, 20, c.VerPrimero())
	require.Equal(t, 20, c.Desencolar())
	require.Equal(t, 30, c.VerPrimero())
	require.Equal(t, 30, c.Desencolar())
	require.Equal(t, 40, c.VerPrimero())
	require.Equal(t, 40, c.Desencolar())
	require.Equal(t, 50, c.VerPrimero())
	require.Equal(t, 50, c.Desencolar())
	require.Panics(t, func() { c.Desencolar() })
}

func TestVolumen(t *testing.T) {
	c := TDAcola.CrearColaEnlazada[int]()
	volumen := 10000

	for i := 0; i < volumen; i++ {
		c.Encolar(i)
		require.Equal(t, 0, c.VerPrimero())
	}

	for i := 0; i < volumen; i++ {
		require.Equal(t, i, c.VerPrimero())
		require.Equal(t, i, c.Desencolar())
	}

	require.True(t, c.EstaVacia())
}

func TestColaVaciada(t *testing.T) {
	c := TDAcola.CrearColaEnlazada[int]()

	c.Encolar(5)
	c.Encolar(10)
	c.Desencolar()
	c.Desencolar()

	require.True(t, c.EstaVacia())

	require.Panics(t, func() { c.VerPrimero() })
	require.Panics(t, func() { c.Desencolar() })
}

func TestStrings(t *testing.T) {
	c := TDAcola.CrearColaEnlazada[string]()

	c.Encolar("algoritmos")
	c.Encolar("2")
	c.Encolar("fiuba")

	require.Equal(t, "algoritmos", c.Desencolar())
	require.Equal(t, "2", c.Desencolar())
	require.Equal(t, "fiuba", c.Desencolar())

	require.True(t, c.EstaVacia())
}
