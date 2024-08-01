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

// depthFirstIterator  is an implementation of the Iterator interface
// for traversing a graph using a depth-first search (DFS) algorithm.
type depthFirstIterator[T comparable] struct {
	iteratorProperties[T] // base properties for bread first iterator

	stack []T // a slice that represents the stack of vertices to visit in DFS traversal order.
}

// NewDepthFirstIterator creates a new instance of depthFirstIterator
// and returns it as the Iterator interface.
func NewDepthFirstIterator[T comparable](g grafik.Grafik[T], start T) (Iterator[T], error) {
	v := g.GetVertexByLabel(start)
	if v == nil {
		return nil, grafik.ErrVertexDoesNotExist
	}

	return newDepthFirstIterator(g, start), nil
}

func newDepthFirstIterator[T comparable](g grafik.Grafik[T], start T) *depthFirstIterator[T] {
	return &depthFirstIterator[T]{
		iteratorProperties: iteratorProperties[T]{
			graph:   g,
			start:   start,
			visited: map[T]bool{start: true},
		},

		stack: []T{start},
	}
}

// HasNext returns a boolean indicating whether there are more vertices
// to be visited in the DFS traversal. It returns true if the head index
// is in the range of the queue indices.
func (d *depthFirstIterator[T]) HasNext() bool {
	return len(d.stack) > 0
}

// Next returns the next vertex to be visited in the DFS traversal. It
// pops the latest vertex that has been added to the stack.
// If the HasNext is false, returns nil.
func (d *depthFirstIterator[T]) Next() *grafik.Vertex[T] {
	if !d.HasNext() {
		return nil
	}

	// get the next vertex from the queue
	label := d.stack[len(d.stack)-1]
	d.stack = d.stack[:len(d.stack)-1]
	currentNode := d.graph.GetVertexByLabel(label)

	// add unvisited neighbors to the queue
	neighbors := currentNode.Neighbors()
	for _, neighbor := range neighbors {
		if !d.visited[neighbor.Label()] {
			d.stack = append(d.stack, neighbor.Label())
			d.visited[neighbor.Label()] = true
		}
	}

	return currentNode
}

// Iterate iterates through all the vertices in the DFS traversal order
// and applies the given function to each vertex. If the function returns
// an error, the iteration stops and the error is returned.
func (d *depthFirstIterator[T]) Iterate(f func(v *grafik.Vertex[T]) error) error {
	for d.HasNext() {
		if err := f(d.Next()); err != nil {
			return err
		}
	}

	return nil
}

// Reset resets the iterator by setting the initial state of the iterator.
func (d *depthFirstIterator[T]) Reset() {
	d.stack = []T{d.start}
	d.visited = map[T]bool{d.start: true}
}
