package queue

import (
	"errors"
)

type CircularFifoQueue struct {
	data []interface{}
	size int
	currentSize int
	headIdx int
	tailIdx int
}

var (
	ErrQueueEmpty = errors.New("queue is empty")
)


func NewCircularFifoQueue(n int) *CircularFifoQueue {
	return &CircularFifoQueue{make([]interface{}, n), n, 0, -1, -1}
}

// adding new element to the end of queue
func (queue *CircularFifoQueue) AddElement(value interface{}) {

	if queue.currentSize == 0 {
		queue.headIdx = 0;
		queue.tailIdx = 0;
		queue.data[0] = value
		queue.currentSize++
		return
	}

	queue.tailIdx = (queue.tailIdx + 1) % queue.size

	// this if statement means that we reached head from our tail
	// so its time to replace head (oldest element) with new one
	// simply overwrite head and move head index
	if queue.tailIdx == queue.headIdx {
		queue.headIdx = (queue.headIdx + 1) % queue.size
	}
	queue.data[queue.tailIdx]  = value
	if queue.currentSize < queue.size {
		queue.currentSize++
	}
}

// returns front element of queue, returns error when list is empty
func (queue *CircularFifoQueue) GetFront() (interface{}, error) {
	if queue.currentSize == 0 {
		return nil, ErrQueueEmpty
	}
	return queue.data[queue.headIdx], nil
}

// pops front element
func (queue *CircularFifoQueue) PopFront() (interface{}, error) {
	if queue.currentSize == 0 {
		return nil, ErrQueueEmpty
	}
	value := queue.data[queue.headIdx]
	queue.headIdx = (queue.headIdx + 1) % queue.size
	queue.currentSize--
	return value, nil
}

// returns slice containing all queue elements in order from head to tail
func (queue *CircularFifoQueue) GetElements() []interface{} {
	size := queue.currentSize
	values := make([]interface{}, 0, size)
	for i, count := queue.headIdx, 0; count < size; i, count = (i + 1) % size, count + 1 {
		values = append(values, queue.data[i])
	}
	return values
}

// cleans queue
func (queue *CircularFifoQueue) Flush() {
	queue.currentSize = 0
	queue.headIdx = -1
	queue.tailIdx = -1
	queue.data = make([]interface{}, queue.size)
}

// returns number of existing elements
func (queue *CircularFifoQueue) Len() int {
	return queue.currentSize
}

func (queue *CircularFifoQueue) Size() int {
	return queue.size
}