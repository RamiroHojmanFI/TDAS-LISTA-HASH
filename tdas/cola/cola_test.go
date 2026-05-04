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

func TestMultiprimeros(t *testing.T) {
	
	t.Run("Debe devolver los primeros k elementos correctamente", func(t *testing.T) {
		// Accedemos mediante el nombre del paquete importado
		cola := TDAcola.CrearColaEnlazada[int]()
		cola.Encolar(10)
		cola.Encolar(20)
		cola.Encolar(30)

		resultado := cola.Multiprimeros(2)

		// Verificamos con testify
		require.Equal(t, []int{10, 20}, resultado, "Los elementos devueltos no son los correctos")
		require.Equal(t, 30, cola.Desencolar(), "El elemento restante en la cola no es el esperado")
	})

	t.Run("Si k es mayor al largo debe devolver todos los disponibles", func(t *testing.T) {
		cola := TDAcola.CrearColaEnlazada[string]()
		cola.Encolar("Hola")
		cola.Encolar("Mundo")

		resultado := cola.Multiprimeros(100)

		require.Len(t, resultado, 2)
		require.Equal(t, []string{"Hola", "Mundo"}, resultado)
		require.True(t, cola.EstaVacia())
	})

	t.Run("Con k igual a cero no debe desencolar nada", func(t *testing.T) {
		cola := TDAcola.CrearColaEnlazada[int]()
		cola.Encolar(5)

		resultado := cola.Multiprimeros(0)

		require.Empty(t, resultado)
	})
}