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

package grafik

import "github.com/fitm-elite/grafik/options"

// Vertex represents a node or point in a graph
type Vertex[T comparable] struct {
	label    T
	inDegree int

	neighbors []*Vertex[T]

	properties options.VertexProperties
}

func NewVertex[T comparable](label T, opts ...options.VertexOptionFunc) *Vertex[T] {
	v := &Vertex[T]{label: label}
	for _, opt := range opts {
		opt(&v.properties)
	}

	return v
}

// NeighborByLabel iterates over the neighbor slice and returns the
// vertex which its label is equal to the input label.
//
// It returns nil if there is no neighbor with that label.
func (v *Vertex[T]) NeighborByLabel(label T) *Vertex[T] {
	for i := range v.neighbors {
		if v.neighbors[i].label == label {
			return v.neighbors[i]
		}
	}

	return nil
}

// HasNeighbor checks if the input vertex is the neighbor of the
// current node or not. It returns 'true' if it finds the input
// in the neighbors. Otherwise, returns 'false'.
func (v *Vertex[T]) HasNeighbor(vertex *Vertex[T]) bool {
	return v.NeighborByLabel(vertex.label) != nil
}

// Label returns vertex label.
func (v *Vertex[T]) Label() T {
	return v.label
}

// InDegree returns the number of incoming edges to the current vertex.
func (v *Vertex[T]) InDegree() int {
	return v.inDegree
}

// OutDegree returns the number of outgoing edges to the current vertex.
func (v *Vertex[T]) OutDegree() int {
	return len(v.neighbors)
}

// Degree returns the total degree of the vertex which is the sum of in and out degrees.
func (v *Vertex[T]) Degree() int {
	return v.inDegree + v.OutDegree()
}

// Neighbors returns a copy of neighbor slice.
func (v *Vertex[T]) Neighbors() []*Vertex[T] {
	neighbors := make([]*Vertex[T], 0, len(v.neighbors))
	for idx := range v.neighbors {
		clone := &Vertex[T]{}
		*clone = *v.neighbors[idx]
		neighbors = append(neighbors, clone)
	}

	return neighbors
}

// Weight returns vertex weight.
func (v *Vertex[T]) Weight() float64 {
	return v.properties.Weight()
}
