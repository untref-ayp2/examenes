package parcial

// Esta es la única función a desarrollar
func (l *LinkedList[T]) Interleave(other *LinkedList[T]) {
	aux := NewLinkedList[T]()
	if !l.IsEmpty() && !other.IsEmpty() {
		aux.Append(l.head.data)
		l.RemoveFirst()
		aux.Append(other.head.data)
		other.RemoveFirst()
	}
	if !other.IsEmpty() {
		aux.Append(other.head.data)
		other.RemoveFirst()
	}
	if !l.IsEmpty() {
		aux.Append(l.head.data)
		l.RemoveFirst()
	}
	for !aux.IsEmpty() {
		l.Append(aux.head.data)
		aux.RemoveFirst()
	}
}

// ==============================================
// Todo esto se agrega para que el código compile
// ==============================================
type LinkedNode[T comparable] struct {
	data T
	next *LinkedNode[T]
}

func (n *LinkedNode[T]) SetNext(newNext *LinkedNode[T]) {
	n.next = newNext
}

func NewLinkedListNode[T comparable](data T) *LinkedNode[T] {
	return &LinkedNode[T]{data: data}
}

func (n *LinkedNode[T]) Next() *LinkedNode[T] {
	return n.next
}

type LinkedList[T comparable] struct {
	head *LinkedNode[T]
	tail *LinkedNode[T]
	size int
}

func NewLinkedList[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (l *LinkedList[T]) IsEmpty() bool {
	return l.size == 0
}

func (l *LinkedList[T]) Append(data T) {
	newNode := NewLinkedListNode(data)
	if l.IsEmpty() {
		l.head = newNode
	} else {
		l.tail.SetNext(newNode)
	}
	l.tail = newNode
	l.size++
}

func (l *LinkedList[T]) RemoveFirst() {
	if l.IsEmpty() {
		return
	}

	l.head = l.head.Next()

	if l.head == nil {
		l.tail = nil
	}

	l.size--
}
