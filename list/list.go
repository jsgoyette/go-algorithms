package list

import "sync"

type List struct {
	head   *Item
	last   *Item
	len    int
	locker sync.RWMutex
}

type Item struct {
	Val  interface{}
	next *Item
	prev *Item
	list *List
}

func New() *List {
	list := &List{}
	list.len = 0
	return list
}

func (list *List) First() *Item {
	return list.head
}

func (list *List) Last() *Item {
	return list.last
}

func (list *List) Length() int {
	return list.len
}

func (item *Item) Prev() *Item {
	return item.prev
}

func (item *Item) Next() *Item {
	return item.next
}

func (list *List) Insert(value interface{}) *List {
	newItem := &Item{value, list.head, list.last, list}
	list.locker.Lock()
	defer list.locker.Unlock()

	if list.head == nil {
		list.head = newItem
		list.last = newItem
	} else {
		list.head.prev = newItem
		list.last.next = newItem
		list.last = newItem
	}

	list.len++

	return list
}

func (list *List) Has(value interface{}) bool {
	if list.head == nil {
		return false
	}
	first := list.First()

	for {
		if first.Val == value {
			return true
		} else {
			if first == list.Last() {
				return false
			}

			first = first.next
		}
	}

	return false
}

func (list *List) Remove(value interface{}) *List {
	list.locker.Lock()
	defer list.locker.Unlock()

	if list.head == nil {
		return list
	}

	first := list.First()

	for {
		if first.Val == value {
			first.prev.next = first.next
			first.next.prev = first.prev
			first.prev = nil
			first.next = nil
			first.Val = nil
			first.list = nil
			list.len--
			return list
		}

		first = first.next
	}
}
