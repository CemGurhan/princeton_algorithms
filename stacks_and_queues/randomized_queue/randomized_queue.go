package randomizedqueue

type RandomizedQueue[T any] []T

// verify whether randomized queue is empty or not
func (rq *RandomizedQueue[T]) IsEmpty() bool {
	if len(*rq) == 0 {
		return true
	}

	return false
}

// return the number of items on the randomized queue
func (rq *RandomizedQueue[T]) Size() int {
	return len(*rq)
}

// add item to end of queue
func (rq *RandomizedQueue[T]) Enqueue(item T) {
	randQueueLength := len(*rq)
	if cap(*rq) == randQueueLength {
		newRandQueue := make(RandomizedQueue[T], randQueueLength+1, 2*(randQueueLength+1))
		for i := 0; i < randQueueLength; i++ {
			newRandQueue[i] = (*rq)[i]
		}
		newRandQueue[randQueueLength] = item
		*rq = newRandQueue
		return
	}
	*rq = (*rq)[:randQueueLength+1]
	(*rq)[randQueueLength] = item
}

// remove and return a random item from the queue
// func (rq *RandomizedQueue[T]) Dequeue() (T, error) {
// 	var item T
// 	randQueueLength := len(*rq)
// 	if randQueueLength == 0 {
// 		return item, errors.New("randomized queue is empty")
// 	}

// 	randomIndex := rand.Intn(len(*rq))
// 	if randomIndex == 0 {
// 		item = (*rq)[0]
// 		*rq = (*rq)[1:]
// 		return item, nil
// 	}
// }
