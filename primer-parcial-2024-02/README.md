# Primer Parcial Resuelto / 2024 TN

Comisión Ma/Ju Tarde (2/5/2024)

### Ejercicio 1

Implementar la primitiva (método) de `CircularList`:

```go
func (l *CircularList[T]) InsertAt(data T, pos int)
```

que inserte el dato dado en la posición `pos % size`, dónde _size_ es el tamaño
actual de la lista. Indicar el orden del método. La lista está implementada
sobre nodos doblemente enlazados.

```go
type DoubleLinkedNode[T comparable] struct {
   data T
   next *DoubleLinkedNode[T]
   prev*DoubleLinkedNode[T]
}
func NewDoubleLinkedListNode[T comparable](data T) *DoubleLinkedNode[T]
func (n*DoubleLinkedNode[T]) SetData(data T)
func (n *DoubleLinkedNode[T]) Data() T
func (n*DoubleLinkedNode[T]) SetNext(next *DoubleLinkedNode[T])
func (n*DoubleLinkedNode[T]) HasNext() bool
func (n *DoubleLinkedNode[T]) SetPrev(newPrev*DoubleLinkedNode[T])
func (n *DoubleLinkedNode[T]) Prev()*DoubleLinkedNode[T]
func (n \*DoubleLinkedNode[T]) HasPrev() bool
```

#### Solución 1

Una primera versión de la solución podría ser la siguiente:

```go
func (l *CircularList[T]) InsertAt(data T, pos int) {
    node := NewDoubleLinkedNode(data)

    if l.Size() == 0 {
        l.head = node
        node.SetPrev(node)
        node.SetNext(node)
        l.size++
        return
    }

    pos = pos % l.Size()
    current := l.Head()
    for i := 0; i < pos; i++ {
        current = current.Next()
    }

    if pos == 0 {
        l.head = node
    }

    node.SetNext(current)
    node.SetPrev(current.Prev())
    node.Prev().SetNext(node)
    current.SetPrev(node)
    l.size++
}
```

Sobre la base de la solución anteior, luego de iterar y extraer algunos método y
utilizando otros existentes en `CircularList`, se puede llegar a la siguiente
solución:

```go
func (n *DoubleLinkedNode[T]) InsertBefore(data T) {
    node := NewDoubleLinkedNode(data)
    node.SetNext(n)
    node.SetPrev(n.Prev())
    node.Prev().SetNext(node)
    n.SetPrev(node)
}

func (l *CircularList[T]) NodeAt(pos int) *DoubleLinkedNode {
    current := l.Head()
    for i := 0; i < pos; i++ {
        current = current.Next()
    }
    return current
}

func (l *CircularList[T]) InsertAt(data T, pos int) {
    pos = pos % l.Size()

    if l.Size() == 0 || pos == 0 {
        l.Prepend(data)
        return
    }

    current := l.NodeAt(pos)
    current.InsertBefore(data)
    l.size++
}
```

### Ejercicio 2

Implementar una primitiva (método) de `BitMap`, llamada `String`, que devuelva
una cadena de caracteres con la representación binaria del mapa de bits (la
cadena siempre debe tener 32 caracteres, ceros y/o unos).

```go
type BitMap struct {
    bits uint32
}
func NewBitMap() *BitMap
func (bm *BitMap) On(pos uint8) error
func (bm *BitMap) Off(pos uint8) error
```

#### Solución 2

Podemos ver este ejercicio de 2 formas, por un lado como un ejercicio de
`BitMap` donde iteramos cada bit y chequeando si está prendido o apagado,
mostraremos un `1` o un `0`:

La complicación de este enfoque es que el `BitMap` del parcial no tenía
implementado un método `IsOn` cómo el visto en clase:

```go
func (bm *BitMap) String1() string {
    str := ""
    for pos := uint8(0); pos < 32; pos++ {
        if bm.bits&(1<<pos) != 0 {
            str = "1" + str
        } else {
            str = "0" + str
        }
    }
    return str
}
```

o bien, podríamos extraer la condición dentro del `if` como un método de
`BitMap`, de la siguiente forma:

```go
func (bm *BitMap) IsOn(pos uint8) bool {
    mask := uint32(1 << pos)
    return bm.bits&mask != 0b0
}

func (bm *BitMap) String1() string {
    str := ""
    for pos := uint8(0); pos < 32; pos++ {
        if bm.IsOn(pos) { // <<< Esta es la única línea que difiere
                          //     con la versión anterior
            str = "1" + str
        } else {
            str = "0" + str
        }
    }
    return str
}
```

El otro enfoque es el de entender que lo que se nos pide es la representación en
binario del campo `bits` del `BitMap`, por lo que nuestro método podría ser
simplemente un conversor de `int` a su representación binaria en `string`:

```go
func (bm BitMap) String() string {
    str := ""
    bits := bm.bits // copiamos el valor para no modificar
                    // los valores del mapa de bits
    for pos := uint8(0); pos < 32; pos++ {
        resto := bits % 2
        bits /= 2

        if resto == 1 {
            str = "1" + str
        } else {
            str = "0" + str
        }
    }
    return str
}
```

(recomiendamos mucho repasar sistemas de numeración y cambio de bases entre los
distintos sistemas, YouTube tiene montones de videos cortos explicando estos
temas).

> **Yapa!**
>
> Si hubiera que implementar esto en la realidad sería algo como esto, que
> claramente no es un hack válido para el parcial:
>
> ```go
> func (bm *BitMap) String() string {
>    return fmt.Sprintf("%032b", bm.bits)
> }
> ```

### Ejercicio 3

Escribir una función recursiva que reciba como parámetro una cola y devuelva la
cola invertida, es decir, el último elemento de la cola pasa a ser el primero y
viceversa. No utilizar ninguna estructura de datos auxiliar.

#### Solución 3

```go
func InvertirCola(q *queue.Queue) {
    if q.IsEmpty() {
        return
    }

    elem, _ := q.Dequeue()
    InvertirCola(q)
    q.Enqueue(elem)
}
```

### Ejercicio 4

Dado un arreglo de enteros ordenados, cuyos elementos se pueden repetir,
escribir una función que utilice división y conquista para devolver el índice de
la primera aparición de un número buscado. Por ejemplo si `a = [1, 1, 2, 5, 5,
5, 6, 8]` y se busca el `5` debe devolver `3`. Si el elemento no se encuentra
debe devolver `-1`. Calcular su orden.

#### Solución 4

```go
func buscar(lista []int, inicio int, fin int, buscado int) int {
    if inicio > fin {
        return -1
    }

    medio := (inicio + fin) / 2

    if lista[medio] > buscado {
        return buscar(lista, inicio, medio-1, buscado)
    }

    if lista[medio] < buscado {
        return buscar(lista, medio+1, fin, buscado)
    }

    if medio > 0 && lista[medio-1] == buscado {
        return buscar(lista, inicio, medio-1, buscado)
    }

    return medio
}