package parcial

// Esta es la única función a desarrollar
func (l *DoubleLinkedList[T]) Extend(other *DoubleLinkedList[T]) {
	l.tail.next = other.head
	other.head.prev = l.tail
	l.tail = other.tail
	l.size += other.size
	other.Clear()
}

// ==============================================
// Todo esto se agrega para que el código compile
// ==============================================

type DoubleLinkedList[T comparable] struct {
	head *DoubleLinkedNode[T]
	tail *DoubleLinkedNode[T]
	size int
}

func (l *DoubleLinkedList[T]) Clear() {
	l.head = nil
	l.tail = nil
	l.size = 0
}

type DoubleLinkedNode[T comparable] struct {
	data T
	next *DoubleLinkedNode[T]
	prev *DoubleLinkedNode[T]
}
