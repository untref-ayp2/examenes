# Primer Parcial Resuelto / 2024 TN

1. Implementar la primitiva (método) de CircularList, `func (l *CircularList[T])InsertAt(data T, pos int){...}` que inserte el dato dado en la posición pos % size, dónde size es el tamaño actual de la lista. Indicar el orden del método. La lista está implementada sobre nodos doblemente enlazados.

```go
	type DoubleLinkedNode[T comparable] struct {
		data T
		next *DoubleLinkedNode[T]
		prev *DoubleLinkedNode[T]
	}
	func NewDoubleLinkedListNode[T comparable](data T) *DoubleLinkedNode[T]
	func (n *DoubleLinkedNode[T]) SetData(data T)
	func (n *DoubleLinkedNode[T]) Data() T
	func (n *DoubleLinkedNode[T]) SetNext(next *DoubleLinkedNode[T])
	func (n *DoubleLinkedNode[T]) HasNext() bool
	func (n *DoubleLinkedNode[T]) SetPrev(newPrev *DoubleLinkedNode[T])
	func (n *DoubleLinkedNode[T]) Prev() *DoubleLinkedNode[T]
	func (n *DoubleLinkedNode[T]) HasPrev() bool
```

2. Implementar una primitiva (método) de BitMap, llamada `String`, que devuelva una cadena de caracteres con la representación binaria del mapa del bit (la cadena siempre debe tener 32 caracteres, ceros y/o unos).

```go
	type BitMap struct {
		bits uint32
	}
	func NewBitMap() *BitMap
	func (bm *BitMap) On(pos uint8) error
	func (bm *BitMap) Off(pos uint8) error
```

3. Escribir una función recursiva que reciba como parámetro una cola y devuelva la cola invertida, es decir, el último elemento de la cola pasa a ser el primero y viceversa. No utilizar ninguna estructura de datos auxiliar.

4. Dado un arreglo de enteros ordenados, cuyos elementos se pueden repetir, escribir una función que utilice división y conquista para devolver el índice de la primera aparición de un número buscado. Por ejemplo si a = [1,1,2,5,5,5,6,8] y se busca el 5 debe devolver 3. Si el elemento no se encuentra debe devolver -1. Calcular su orden
