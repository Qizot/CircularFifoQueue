package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	n := 10
	q := NewCircularFifoQueue(n)
	assert.NotNil(t, q)
	assert.NotNil(t, q.data)
	assert.Equal(t, 0, q.currentSize)
	assert.Equal(t, 0, q.Len())
	assert.Equal(t, n, q.size)
	assert.Equal(t, n, q.Size())
	assert.Equal(t, []interface{}{}, q.GetElements())
	assert.Equal(t, -1, q.headIdx)
	assert.Equal(t, -1, q.tailIdx)
}

func TestFullQueue(t *testing.T) {
	n := 10
	q := NewCircularFifoQueue(n)
	for i := 0; i < 2*n; i++ {
		q.AddElement(i)
	}
	values := q.GetElements()
	for i := n; i < 2*n; i++ {
		assert.Equal(t, values[i-n].(int), i)
	}
}

func TestFlush(t *testing.T) {
	n := 10
	q := NewCircularFifoQueue(n)
	for i := 0; i < n; i++ {
		q.AddElement(i)
	}
	q.Flush()
	assert.Equal(t, 0, q.Len())
	assert.Equal(t, n, q.Size())
	assert.Equal(t, []interface{}{}, q.GetElements())
	assert.Equal(t, -1, q.headIdx)
	assert.Equal(t, -1, q.tailIdx)
}

func TestGetFront(t *testing.T) {
	n := 10
	q := NewCircularFifoQueue(n)
	q.AddElement(n)
	val, err := q.GetFront()
	assert.NoError(t, err)
	assert.Equal(t, val, n)

	q.Flush()
	val, err = q.GetFront()
	assert.Error(t, err)
	assert.Nil(t, val)

}

func TestPopFront(t *testing.T) {
	n := 10
	q := NewCircularFifoQueue(n)
	q.AddElement(n)
	assert.Equal(t, q.Len(), 1)
	val, err := q.PopFront()
	assert.NoError(t, err)
	assert.Equal(t, val, n)
	assert.Equal(t, q.Len(), 0)

	val, err = q.PopFront()
	assert.Error(t, err)
	assert.Nil(t, val)
	assert.Equal(t, q.Len(), 0)
}

func TestPopAndFront(t *testing.T) {
	n := 10
	q := NewCircularFifoQueue(n)
	for i := 0; i < n; i++ {
		q.AddElement(i)
	}

	i := 0
	for q.Len() != 0 {
		val, err := q.PopFront()
		assert.NoError(t, err)
		assert.Equal(t, val, i)
		i++
	}
}
