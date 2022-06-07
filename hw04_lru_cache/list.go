package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	head  *ListItem
	tail  *ListItem
	count int
}

func NewList() List {
	return new(list)
}

func (l *list) Len() int {
	return l.count
}

func (l *list) Remove(i *ListItem) {
	if i.Prev != nil {
		i.Prev.Next = i.Next
	} else {
		l.head = i.Next
	}

	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.tail = i.Prev
	}

	if l.count == 0 {
		return
	}
	l.count--
}

func (l *list) Front() *ListItem {
	return l.head
}

func (l *list) Back() *ListItem {
	return l.tail
}

func (l *list) PushFront(v interface{}) *ListItem {
	item := ListItem{
		Value: v,
		Next:  l.head,
	}

	if item.Next != nil {
		l.head.Prev = &item
	}

	if l.tail == nil {
		l.tail = &item
	}

	l.count++
	l.head = &item
	return &item
}

func (l *list) PushBack(v interface{}) *ListItem {
	item := ListItem{
		Value: v,
		Prev:  l.tail,
	}

	if item.Prev != nil {
		l.tail.Next = &item
	} else {
		l.head = &item
	}

	l.count++
	l.tail = &item
	return &item
}

func (l *list) MoveToFront(i *ListItem) {
	if l.head == i {
		return
	}

	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.tail = i.Prev
	}

	if i.Prev != nil {
		i.Prev.Next = i.Next
		i.Next = l.head
		l.head.Prev = i
	}

	i.Prev = nil
	l.head = i
}
