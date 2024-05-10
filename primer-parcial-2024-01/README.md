# Primer Parcial Resuelto / 2024 TM

1. Implementar la primitiva (método) de DoubleLinkedList, `func (l *DoubleLinkedList[T])Extend(other *DoubleLinkedList[T]){...}` que extienda la lista con todos los elementos de la otra lista que se pasa por parámetro. Indicar y justificar el orden del algoritmo. Al finalizar la ejecución la lista sobre la que se ejecutó `Extend` debe tener al final de la misma concatenada la lista `other` y la lista `other` debe quedar vacía. El método debe ser O(1).

Ejemplo:
```go
	l1 := NewDoubleLinkedList(1, 3, 5)
	l2 := NewDoubleLinkedList(1, 2, 3)
	l1.Extend(l2) // l1 = {1, 3, 5, 1, 2, 3}; l2 = {}
```

Código suministrado:
```go
	type DoubleLinkedList[T comparable] struct {
		head *DoubleLinkedNode[T]
		tail *DoubleLinkedNode[T]
		size int
	}
	func NewDoubleLinkedList[T comparable](elements ...T) *DoubleLinkedList[T]
	func (l *DoubleLinkedList[T]) Head() *DoubleLinkedNode[T]}
	func (l *DoubleLinkedList[T]) Tail() *DoubleLinkedNode[T]
	func (l *DoubleLinkedList[T]) Size() int
	func (l *DoubleLinkedList[T]) IsEmpty() bool
	func (l *DoubleLinkedList[T]) Clear() {
		l.head = nil
		l.tail = nil
		l.size = 0
	}
	func (l *DoubleLinkedList[T]) Prepend(data T)
	func (l *DoubleLinkedList[T]) Append(data T)
	func (l *DoubleLinkedList[T]) Find(data T) *DoubleLinkedNode[T]
	func (l *DoubleLinkedList[T]) RemoveFirst()
	func (l *DoubleLinkedList[T]) RemoveLast()
	func (l *DoubleLinkedList[T]) Remove(data T)
```

2. Escribir una función que reciba como parámetro un diccionario cuyas claves son títulos de libros (strings) y cada clave tiene asociada un conjunto de autores (string). La función debe devolver un diccionario cuyas claves sean los autores y cuyo valor una lista enlazada simple de títulos de libros.

3. Dada una tabla de hash de tamaño 11 mostrar paso a paso las siguientes operaciones. Insertar 11, Insertar 17, Insertar 0, Eliminar 11, Insertar 28, Insertar 33. Resolver las colisiones con el método de hashing cerrado. la función de hash es `h(x) = x mod 11`

4. Implementar por división y conquista una función que reciba un arreglo de números enteros positivos desordenados y devuelva el máximo del arreglo. Aplicar el teorema del maestro y calcular el orden.
