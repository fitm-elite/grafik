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

import (
	"errors"
)

var (
	ErrNilVertices        = errors.New("vertices are nil")
	ErrVertexDoesNotExist = errors.New("vertex does not exist")
	ErrEdgeAlreadyExists  = errors.New("edge already exists")
)

type grafik[T comparable] struct {
	vertices map[T]*Vertex[T]
	edges    map[T]map[T]*Edge[T]
}

type VertexFunc[T comparable] interface {
	// AddVertex adds a new vertex with the given label to the graph.
	// Label of the vertex is a comparable type. This method also accepts the
	// vertex properties such as weight.
	//
	// If there is a vertex with the same label in the graph, returns nil.
	// Otherwise, returns the created vertex.
	AddVertexByLabel(label T, options ...VertexOptionFunc) *Vertex[T]

	// AddVertex adds the input vertex to the graph. It doesn't add
	// vertex to the graph if the input vertex label is already exists
	// in the graph.
	AddVertex(v *Vertex[T])

	// GetVertexByID returns the vertex with the input label.
	//
	// If vertex doesn't exist, returns nil.
	GetVertexByLabel(label T) *Vertex[T]

	// GetAllVertices returns a slice of all existing vertices in the graph.
	GetAllVertices() []*Vertex[T]

	// ContainsVertex returns 'true' if this graph contains the specified vertex.
	//
	// If the specified vertex is nil, returns 'false'.
	ContainsVertex(v *Vertex[T]) bool
}

type EdgeFunc[T comparable] interface {
	// AddEdge adds and edge from the vertex with the 'from' label to
	// the vertex with the 'to' label by appending the 'to' vertex to the
	// 'neighbors' slice of the 'from' vertex, in directed graph.
	//
	// In undirected graph, it creates edges in both directions between
	// the specified vertices.
	//
	// It creates the input vertices if they don't exist in the graph.
	// If any of the specified vertices is nil, returns nil.
	// If edge already exist, returns error.
	AddEdge(from, to *Vertex[T], opts ...EdgeOptionFunc) (*Edge[T], error)

	// GetAllEdges returns a slice of all edges connecting source vertex to
	// target vertex if such vertices exist in this graph.
	//
	// In directed graph, it returns a single edge.
	//
	// If any of the specified vertices is nil, returns nil.
	// If any of the vertices does not exist, returns nil.
	// If both vertices exist but no edges found, returns an empty set.
	GetAllEdges(from, to *Vertex[T]) []*Edge[T]

	// GetEdge returns an edge connecting source vertex to target vertex
	// if such vertices and such edge exist in this graph.
	//
	// In undirected graph, returns only the edge from the "from" vertex to
	// the "to" vertex.
	//
	// If any of the specified vertices is nil, returns nil.
	// If edge does not exist, returns nil.
	GetEdge(from, to *Vertex[T]) *Edge[T]

	// ContainsEdge returns 'true' if and only if this graph contains an edge
	// going from the source vertex to the target vertex.
	//
	// If any of the specified vertices does not exist in the graph, or if is nil,
	// returns 'false'.
	ContainsEdge(from, to *Vertex[T]) bool
}

type Grafik[T comparable] interface {
	VertexFunc[T]
	EdgeFunc[T]
}

func New[T comparable]() Grafik[T] {
	return &grafik[T]{
		vertices: make(map[T]*Vertex[T]),
		edges:    make(map[T]map[T]*Edge[T]),
	}
}

//
// Vertex implementations
//

func (g *grafik[T]) addVertex(v *Vertex[T]) *Vertex[T] {
	if _, ok := g.vertices[v.label]; ok {
		return nil
	}

	g.vertices[v.label] = v

	return v
}

func (g *grafik[T]) findVertex(label T) *Vertex[T] {
	return g.vertices[label]
}

// AddVertexByLabel adds a new vertex with the given label to the graph.
// Label of the vertex is a comparable type. This method also accepts the
// vertex properties such as weight.
//
// If there is a vertex with the same label in the graph, returns nil.
// Otherwise, returns the created vertex.
func (g *grafik[T]) AddVertexByLabel(label T, opts ...VertexOptionFunc) *Vertex[T] {
	var properties VertexProperties
	for _, opt := range opts {
		opt(&properties)
	}

	v := g.addVertex(&Vertex[T]{label: label, properties: properties})

	return v
}

// AddVertex adds the input vertex to the graph. It doesn't add
// vertex to the graph if the input vertex label is already exists
// in the graph.
func (g *grafik[T]) AddVertex(v *Vertex[T]) {
	if v == nil {
		return
	}

	g.addVertex(v)
}

// GetVertexByID returns the vertex with the input label.
//
// If vertex doesn't exist, returns nil.
func (g *grafik[T]) GetVertexByLabel(label T) *Vertex[T] {
	return g.findVertex(label)
}

// GetAllVertices returns a slice of all existing vertices in the graph.
func (g *grafik[T]) GetAllVertices() []*Vertex[T] {
	vertices := make([]*Vertex[T], 0, len(g.vertices))
	for _, vertex := range g.vertices {
		vertices = append(vertices, vertex)
	}

	return vertices
}

// ContainsVertex returns 'true' if this graph contains the specified vertex.
//
// If the specified vertex is nil, returns 'false'.
func (g *grafik[T]) ContainsVertex(v *Vertex[T]) bool {
	if v == nil {
		return false
	}

	return g.findVertex(v.label) != nil
}

//
// Edge implementations
//

// addToEdgeMap creates a new edge struct and adds it to the edges map inside
// the baseGraph struct. Note that it doesn't add the neighbor to the source vertex.
//
// It returns the created edge.
func (g *grafik[T]) addToEdgeMap(from, to *Vertex[T], opts ...EdgeOptionFunc) *Edge[T] {
	edge := NewEdge(from, to, opts...)
	if _, ok := g.edges[from.label]; !ok {
		g.edges[from.label] = map[T]*Edge[T]{to.label: edge}
	} else {
		g.edges[from.label][to.label] = edge
	}

	return edge
}

// AddEdge adds and edge from the vertex with the 'from' label to
// the vertex with the 'to' label by appending the 'to' vertex to the
// 'neighbors' slice of the 'from' vertex, in directed graph.
//
// In undirected graph, it creates edges in both directions between
// the specified vertices.
//
// It creates the input vertices if they don't exist in the graph.
// If any of the specified vertices is nil, returns nil.
// If edge already exist, returns error.
func (g *grafik[T]) AddEdge(from, to *Vertex[T], opts ...EdgeOptionFunc) (*Edge[T], error) {
	if from == nil || to == nil {
		return nil, ErrNilVertices
	}

	if g.findVertex(from.label) == nil {
		g.AddVertex(from)
	}

	if g.findVertex(to.label) == nil {
		g.AddVertex(to)
	}

	// prevent edge-multiplicity
	if g.ContainsEdge(from, to) {
		return nil, ErrEdgeAlreadyExists
	}

	from = g.vertices[from.label]
	to = g.vertices[to.label]

	from.neighbors = append(from.neighbors, to)
	to.inDegree++

	// add "from" to the "to" vertex neighbor slice, if graph is undirected.
	to.neighbors = append(to.neighbors, from)
	from.inDegree++

	g.addToEdgeMap(to, from, opts...)

	return g.addToEdgeMap(from, to, opts...), nil
}

// GetEdge returns an edge connecting source vertex to target vertex
// if such vertices and such edge exist in this graph.
//
// In undirected graph, returns only the edge from the "from" vertex to
// the "to" vertex.
//
// If any of the specified vertices is nil, returns nil.
// If edge does not exist, returns nil.
func (g *grafik[T]) GetEdge(from, to *Vertex[T]) *Edge[T] {
	if from == nil || to == nil {
		return nil
	}

	if g.findVertex(from.label) == nil {
		return nil
	}

	if g.findVertex(to.label) == nil {
		return nil
	}

	if destMap, ok := g.edges[from.label]; ok {
		return destMap[to.label]
	}

	return nil
}

// GetAllEdges returns a slice of all edges connecting source vertex to
// target vertex if such vertices exist in this graph.
//
// In directed graph, it returns a single edge.
//
// If any of the specified vertices is nil, returns nil.
// If any of the vertices does not exist, returns nil.
// If both vertices exist but no edges found, returns an empty set.
func (g *grafik[T]) GetAllEdges(from, to *Vertex[T]) []*Edge[T] {
	if from == nil || to == nil {
		return nil
	}

	if g.findVertex(from.label) == nil {
		return nil
	}

	if g.findVertex(to.label) == nil {
		return nil
	}

	edges := make([]*Edge[T], 0, len(g.edges))

	if destMap, ok := g.edges[from.label]; ok {
		if edge, ok := destMap[to.label]; ok {
			edges = append(edges, edge)
		}
	}

	if destMap, ok := g.edges[to.label]; ok {
		if edge, ok := destMap[from.label]; ok {
			edges = append(edges, edge)
		}
	}

	return edges
}

// ContainsEdge returns 'true' if and only if this graph contains an edge
// going from the source vertex to the target vertex.
//
// If any of the specified vertices does not exist in the graph, or if is nil,
// returns 'false'.
func (g *grafik[T]) ContainsEdge(from, to *Vertex[T]) bool {
	if from == nil || to == nil {
		return false
	}

	if g.findVertex(from.label) == nil {
		return false
	}

	if g.findVertex(to.label) == nil {
		return false
	}

	if dest, ok := g.edges[from.label]; ok {
		if _, ok = dest[to.label]; ok {
			return true
		}
	}

	return false
}
