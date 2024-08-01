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

package iterator

import "github.com/fitm-elite/grafik"

// breadthFirstIterator is an implementation of the Iterator interface
// for traversing a graph using a breadth-first search (BFS) algorithm.
type breadthFirstIterator[T comparable] struct {
	iteratorProperties[T] // base properties for bread first iterator

	queue []T // a slice that represents the queue of vertices to visit in BFS traversal order.
	head  int // the current head of the queue.
}

// NewBreadthFirstIterator creates a new instance of breadthFirstIterator
// and returns it as the Iterator interface.
func NewBreadthFirstIterator[T comparable](g grafik.Grafik[T], start T) (Iterator[T], error) {
	v := g.GetVertexByLabel(start)
	if v == nil {
		return nil, grafik.ErrVertexDoesNotExist
	}

	return newBreadthFirstIterator(g, start), nil
}

func newBreadthFirstIterator[T comparable](g grafik.Grafik[T], start T) *breadthFirstIterator[T] {
	return &breadthFirstIterator[T]{
		iteratorProperties: iteratorProperties[T]{
			graph:   g,
			start:   start,
			visited: map[T]bool{start: true},
		},

		queue: []T{start},
		head:  -1,
	}
}

// HasNext returns a boolean indicating whether there are more vertices
// to be visited in the BFS traversal. It returns true if the head index
// is in the range of the queue indices.
func (d *breadthFirstIterator[T]) HasNext() bool {
	return d.head < len(d.queue)-1
}

// Next returns the next vertex to be visited in the BFS traversal.
// It dequeues the next vertex from the queue and updates the head field.
// If the HasNext is false, returns nil.
func (d *breadthFirstIterator[T]) Next() *grafik.Vertex[T] {
	if !d.HasNext() {
		return nil
	}

	d.head++

	// get the next vertex from the queue
	currentNode := d.graph.GetVertexByLabel(d.queue[d.head])

	// add unvisited neighbors to the queue
	neighbors := currentNode.Neighbors()
	for _, neighbor := range neighbors {
		if !d.visited[neighbor.Label()] {
			d.visited[neighbor.Label()] = true
			d.queue = append(d.queue, neighbor.Label())
		}
	}

	return currentNode
}

// Iterate iterates through all the vertices in the BFS traversal order
// and applies the given function to each vertex. If the function returns
// an error, the iteration stops and the error is returned.
func (d *breadthFirstIterator[T]) Iterate(f func(v *grafik.Vertex[T]) error) error {
	for d.HasNext() {
		if err := f(d.Next()); err != nil {
			return err
		}
	}

	return nil
}

// Reset resets the iterator by setting the initial state of the iterator.
func (d *breadthFirstIterator[T]) Reset() {
	d.queue = []T{d.start}
	d.head = -1
	d.visited = map[T]bool{d.start: true}
}
