package parcial

// Nodo de la Lista
type node struct {
	value int
	next  *node
}

func newNode(value int) *node {
	return &node{value: value, next: nil}
}

// Lista Simplemente enlazada y ordenada
type OrdLinkList struct {
	head *node // puntero al primer nodo
	tail *node // puntero al último nodo
	size int
}

func NewOrdLinkList() *OrdLinkList {
	return &OrdLinkList{head: nil, tail: nil, size: 0}
}

// Inserta un nuevo nodo en la lista ordenada
func (l *OrdLinkList) Insert(value int) {
	newNode := newNode(value)
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		if value < l.head.value {
			newNode.next = l.head
			l.head = newNode
		} else if value > l.tail.value {
			l.tail.next = newNode
			l.tail = newNode
		} else {
			aux := l.head
			for aux.next.value < value {
				aux = aux.next
			}
			newNode.next = aux.next
			aux.next = newNode
		}
	}
	l.size++
}
