# Primer Parcial Resuelto 3 de mayo de 2024 Comisión Miércoles y Viernes



1. Implementar una primitiva (método) de LinkedList, Interleave[T comparable](other
*LinkedList[T]), que intercale los nodos de la lista (receptor) y other en la lista actual. Indicar y
justificar el orden del algoritmo. Al finalizar la ejecución la lista sobre la que se ejecutó Interaleave
debe tener el tamaño de la lista original más el tamaño de other y la lista other debe quedar vacía.

Ejemplo:
```go
    l1 := NewLinkedList(1, 3, 5)
    l2 := NewLinkedList(1, 2, 3)
    l1.Interleave(l2) // l1 = {1, 1, 3, 2, 5, 3}; l2 = {}
```

Código suministrado:
```go
    type LinkedList[T comparable] struct {
        head *LinkedNode[T]
        tail *LinkedNode[T]
        size int
    }
    func NewLinkedList[T comparable](elements ...T) *LinkedList[T]
    func (l *LinkedList[T]) Head() *LinkedNode[T]
    func (l *LinkedList[T]) Tail() *LinkedNode[T]
    func (l *LinkedList[T]) Size() int
    func (l *LinkedList[T]) IsEmpty() bool
    func (l *LinkedList[T]) Clear()
    func (l *LinkedList[T]) Prepend(data T)
    func (l *LinkedList[T]) Append(data T)
    func (l *LinkedList[T]) Find(data T) *LinkedNode[T]
    func (l *LinkedList[T]) RemoveFirst()
    func (l *LinkedList[T]) RemoveLast()
    func (l *LinkedList[T]) Remove(data T)
```

2. Implementar el tipo de dato ArraySet, conjunto que utiliza un slice de GO como contenedor de datos.
Los métodos que debe tener el ArraySet son:

Código suministrado:
```go
    func NewArraySet[T types.Ordered](elements ...T)*ArraySet[T]
    func (s *ArraySet[T]) Contains(element T) bool
    func (s *ArraySet[T]) Add(elements ...T)
    func (s *ArraySet[T]) Remove(element T)
    func (s *ArraySet[T]) Size() int
    func (s *ArraySet[T]) Values() []T
    func (s *ArraySet[T]) String() string
```
Nota: Los conjuntos no pueden contener elementos repetidos


3. Escribir una función recursiva que reciba un arreglo de números enteros y devuelva otro arreglo con
los elementos elevados al cuadrado

4. Si un problema determinado se puede solucionar con un enfoque iterativo con un algoritmo de orden
O(n^2 ) y con otro algoritmo por división y conquista que subdivide el problema original en dos
subproblemas y realiza dos llamadas recursivas y el orden para dividir el problema y luego juntar las
soluciones parciales es O(n) ¿Cuál de los dos algoritmos es más eficiente? Justificar aplicando el
teorema del maestro
