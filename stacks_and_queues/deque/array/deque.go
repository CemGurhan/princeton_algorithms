package deque

import "errors"

type Deque[T any] []T

// return the number of items on the deque
func (d *Deque[T]) Size() int {
	return len(*d)
}

// add item to the front of the deque
func (d *Deque[T]) AddFirst(item T) {
	dequeLength := len(*d)
	if cap(*d) == dequeLength {
		newDeque := make(Deque[T], dequeLength+1, 2*(dequeLength+1))
		for i := 0; i < dequeLength; i++ {
			newDeque[i+1] = (*d)[i]
		}
		newDeque[0] = item
		*d = newDeque
		return
	}
	newDeque := make(Deque[T], dequeLength+1)
	newDeque[0] = item
	copy(newDeque[1:], *d)

	*d = newDeque
}

// add item to the back of the deque
func (d *Deque[T]) AddLast(item T) {
	dequeLength := len(*d)
	if cap(*d) == dequeLength {
		newDeque := make(Deque[T], dequeLength+1, 2*(dequeLength+1))
		for i, originalItem := range *d {
			newDeque[i] = originalItem
		}
		newDeque[dequeLength] = item
		(*d) = newDeque
		return
	}
	*d = (*d)[:dequeLength+1]
	(*d)[dequeLength] = item
}

// remove and return the item from the front
func (d *Deque[T]) RemoveFirst() (T, error) {
	var removedItem T

	if d.IsEmpty() {
		return removedItem, errors.New("cannot remove from empty deque")
	}

	removedItem = (*d)[0]
	*d = (*d)[1:]

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
