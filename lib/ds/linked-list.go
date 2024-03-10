package ds

type LinkedList[V any] struct {
	size  int
	front *LinkedListNode[V]
	back  *LinkedListNode[V]
}

func NewLinkedList[V any]() *LinkedList[V] {
	return &LinkedList[V]{size: 0}
}

func (list *LinkedList[V]) Front() V {
	return list.front.value
}

func (list *LinkedList[V]) Back() V {
	return list.back.value
}

func (list *LinkedList[V]) Add(value V) *LinkedListNode[V] {
	return list.back.AddAfter(value)
}

func (list *LinkedList[V]) AddBack(value V) *LinkedListNode[V] {
	return list.back.AddAfter(value)
}

func (list *LinkedList[V]) AddFront(value V) *LinkedListNode[V] {
	return list.front.AddBefore(value)
}

func (list *LinkedList[V]) RemoveFront() V {
	return list.front.Remove()
}

func (list *LinkedList[V]) RemoveBack() V {
	return list.back.Remove()
}

func (list *LinkedList[V]) Len() int {
	return list.size
}

func (list *LinkedList[V]) ToArray() []V {
	ret := make([]V, 0, list.size)

	for p := list.front; p != nil; p = p.next {
		ret = append(ret, p.value)
	}

	return ret
}

type LinkedListNode[V any] struct {
	value  V
	parent *LinkedList[V]
	prev   *LinkedListNode[V]
	next   *LinkedListNode[V]
}

func (node *LinkedListNode[V]) AddAfter(value V) *LinkedListNode[V] {
	next := &LinkedListNode[V]{
		value:  value,
		parent: node.parent,
		prev:   node,
		next:   node.next,
	}

	if node.next != nil {
		node.next.prev = next
	}

	node.next = next

	node.parent.size++
	if node.parent.back == node {
		node.parent.back = next
	}

	return next
}

func (node *LinkedListNode[V]) AddBefore(value V) *LinkedListNode[V] {
	prev := &LinkedListNode[V]{
		value:  value,
		parent: node.parent,
		prev:   node.prev,
		next:   node,
	}

	if node.prev != nil {
		node.prev.next = prev
	}

	node.prev = prev

	node.parent.size++
	if node.parent.front == node.prev {
		node.parent.front = prev
	}

	return prev
}

func (node *LinkedListNode[V]) Remove() V {
	prev := node.prev
	next := node.next

	if prev != nil {
		prev.next = next
	}
	if next != nil {
		next.prev = prev
	}

	node.parent.size--
	if node.parent.front == node {
		node.parent.front = next
	}
	if node.parent.back == node {
		node.parent.back = prev
	}

	return node.value
}
