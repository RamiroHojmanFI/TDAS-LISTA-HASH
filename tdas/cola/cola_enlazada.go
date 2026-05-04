package cola

type nodoCola[T any] struct {
	dato T
	prox *nodoCola[T]
}
type colaEnlazada[T any] struct {
	primero *nodoCola[T]
	ultimo  *nodoCola[T]
}

func CrearColaEnlazada[T any]() Cola[T] {
	return &colaEnlazada[T]{
		primero: nil,
		ultimo:  nil,
	}
}



func (c *colaEnlazada[T]) EstaVacia() bool {

	return c.primero == nil
}

func (c *colaEnlazada[T]) VerPrimero() T {
	if c.EstaVacia() {
		panic("La cola esta vacia")
	}
	return c.primero.dato

}

func crearNodo[T any](elem T) *nodoCola[T] {
	return &nodoCola[T]{
		dato: elem,
		prox: nil,
	}
}

func (c *colaEnlazada[T]) Encolar(elem T) {
	aux := crearNodo(elem)
	if c.primero == nil {
		c.primero = aux
		c.ultimo = aux
		return
	}
	c.ultimo.prox = aux
	c.ultimo = aux

}

func (c *colaEnlazada[T]) Desencolar() T {
	if c.EstaVacia() {
		panic("La cola esta vacia")
	}
	aux := c.primero.dato
	if c.primero.prox != nil {
		c.primero = c.primero.prox
	} else {
		c.primero = nil
		c.ultimo = nil
	}
	return aux
}
func (cola *colaEnlazada[T]) Multiprimeros(k int) []T {
	devolucion := make([]T, 0, k)
    
    for i := 0; i < k && !cola.EstaVacia(); i++ {
        devolucion = append(devolucion, cola.Desencolar())
    }
    
    return devolucion
}
