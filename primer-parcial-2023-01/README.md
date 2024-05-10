# Primer Parcial Resuelto / 2023 TM

```go
// Nodo de la Lista

type node[int] struct {
  value int
	next  *node[int]
}

func newNode[int](value int) *node[int] {
  return &node[int]{value: value, next: nil}
}

// Lista Simplemente enlazada y ordenada
type OrdLinkList[int] struct {
	head *node[int] // puntero al primer nodo
	tail *node[int] // puntero al último nodo
	size int
}

func NewOrdLinkList[int]() *OrdLinkList[int] {
	return &OrdLinkList[int]{head: nil, tail: nil, size: 0}
}

func (l *OrdLinkList[int]) Insert(value int) {
	//TODO
}
```

1. Dada la anterior implementación de una lista enlazada simple, pero que debe mantener el orden de sus elementos. Se pide. Implementar el método Insertar, que reciba como parámetro un elemento y lo inserte en la posición que corresponda. Justificar el orden de Insertar


2. Dado un diccionario que contiene las notas de los estudiantes, escriba una función que devuelva la nota final (promedio de notas de cada alumno). Ej {"Perez" : [4,4,8,6], "Sánchez": [7,5,7,7], "Flores": [4,9,8]}. Debe devolver {"Perez": 5.50, "Sánchez": 6.50, "Flores": 7.0}

3. Escribir una función recursiva que reciba como parámetro una Cola y devuelva una Cola invertida. Por ejemplo, si la Cola que recibe es [1, 2, 3, 4, 5] donde 1 es el primer elemento que entró en la cola, debe devolver [5, 4, 3, 2, 1]

4. Escribir una función recursiva, que use la técnica de división y conquista, que reciba como parámetro una cadena de caracteres y un carácter (también como string) y devuelve la cantidad de veces que aparece dicho carácter en la cadena. Calcular el orden de la misma utilizando el teorema del maestro. Comparar con el orden teórico de una solución trivial que recorre cada elemento de la cadena y va acumulando las repeticiones del carácter buscado
