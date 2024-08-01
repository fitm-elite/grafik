// Copyright (c) 2024 Faculty of Industrial Technology and Management, KMUTNB (Provided by FITM Elite)
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package queue

import (
	"container/heap"

	"github.com/fitm-elite/grafik"
)

// VertexPriorityQueue wraps the priorityQueue type to decrease the
// exposes methods, and increase the type safety.
type VertexPriorityQueue[T comparable] struct {
	pq priorityQueue[T] // a slice of VertexWithPriority that represents min heap.
}

func NewVertexPriorityQueue[T comparable]() *VertexPriorityQueue[T] {
	pq := make(priorityQueue[T], 0)
	heap.Init(&pq)
	return &VertexPriorityQueue[T]{
		pq: pq,
	}
}

// Push adds new VertexWithPriority to the queue.
func (v *VertexPriorityQueue[T]) Push(in *VertexWithPriority[T]) {
	heap.Push(&v.pq, in)
}

// Pop removes and returns the minimum element (according to Less) from
// the underlying heap.
func (v *VertexPriorityQueue[T]) Pop() *VertexWithPriority[T] {
	out, _ := heap.Pop(&v.pq).(*VertexWithPriority[T])
	return out
}

// Peek is the number of elements in the underlying queue.
func (v *VertexPriorityQueue[T]) Peek() *VertexWithPriority[T] {
	if v.Len() > 0 {
		return v.pq[0]
	}

	return nil
}

// Len is the number of elements in the underlying queue.
func (v *VertexPriorityQueue[T]) Len() int {
	return len(v.pq)
}

// VertexWithPriority is a vertex priority queue item that stores
// vertex along with its priority.
type VertexWithPriority[T comparable] struct {
	vertex   *grafik.Vertex[T]
	priority float64
	index    int
}

func NewVertexWithPriority[T comparable](vertex *grafik.Vertex[T], priority float64) *VertexWithPriority[T] {
	return &VertexWithPriority[T]{vertex: vertex, priority: priority}
}

// Priority returns the priority of the vertex.
func (v VertexWithPriority[T]) Priority() float64 {
	return v.priority
}

// Vertex returns the vertex.
func (v VertexWithPriority[T]) Vertex() *grafik.Vertex[T] {
	return v.vertex
}

// priorityQueue is a priority queue that implements
// heap and sort interfaces. It represents a min heap.
type priorityQueue[T comparable] []*VertexWithPriority[T]

// Len is the number of elements in the collection.
func (pq priorityQueue[T]) Len() int { return len(pq) }

// Less reports whether the element with index i
// must sort before the element with index j.
func (pq priorityQueue[T]) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

// Swap swaps the elements with indexes i and j.
func (pq priorityQueue[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push adds new item to the collection.
func (pq *priorityQueue[T]) Push(x interface{}) {
	item, ok := x.(*VertexWithPriority[T])
	if !ok {
		return
	}

	n := len(*pq)
	item.index = n
	*pq = append(*pq, item)
}

// Pop removes and returns the minimum element (according to Less) from the heap.
// The complexity is O(log n) where n = h.Len().
// Pop is equivalent to Remove(h, 0).
func (pq *priorityQueue[T]) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	item.index = -1
	*pq = old[0 : n-1]
	return item
}
