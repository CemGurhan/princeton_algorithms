package deque

import "errors"

type Deque[T any] []T

// return the number of items on the deque
func (d *Deque[T]) Size() int {
	return len(*d)
}

// add item to the front of the deque
func (d *Deque[T]) AddFirst(item T) {
	newDeque := make(Deque[T], len(*d)+1)

	for i := len(newDeque) - 1; i > -1; i-- {
		if i == 0 {
			newDeque[i] = item
			continue
		}
		newDeque[i] = (*d)[i-1]
	}

	*d = newDeque
}

// add item to the back of the deque
func (d *Deque[T]) AddLast(item T) {
	newDeque := make(Deque[T], len(*d)+1)

	for i := range newDeque {
		if i == len(newDeque)-1 {
			newDeque[i] = item
			continue
		}
		newDeque[i] = (*d)[i]
	}

	*d = newDeque
}

// remove and return the item from the front
func (d *Deque[T]) RemoveFirst() (T, error) {
	var removedItem T
	if d.IsEmpty() {
		return removedItem, errors.New("cannot remove from empty deque")
	}

	newDeque := make(Deque[T], len(*d)-1)

	for i := 0; i < len(*d); i++ {
		if i == 0 {
			removedItem = (*d)[i]
			continue
		}
		newDeque[i-1] = (*d)[i]
	}

	*d = newDeque

	return removedItem, nil
}

// remove and return the item from the back of the deque
func (d *Deque[T]) RemoveLast() (T, error) {
	var removed T

	if d.IsEmpty() {
		return removed, errors.New("cannot remove from empty deque")
	}

	removed = (*d)[len(*d)-1]
	*d = (*d)[:len(*d)-1]

	return removed, nil
}

// verify whether the deque is empty or not
func (d *Deque[T]) IsEmpty() bool {
	if len(*d) == 0 {
		return true
	}

	return false
}
