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

package pathfinder

import (
	"math"

	"github.com/fitm-elite/grafik"
	"github.com/fitm-elite/grafik/queue"
)

// DijkstraOptionFunc represent an alias of function type that modifies the specified dijkstra properties.
type DijkstraOptionFunc func(properties *DijkstraProperties)

// DijkstraProperties represents the properties of an dijkstra.
type DijkstraProperties struct {
	useStandard bool
}

// WithDijkstraStandard sets the standard algorithm for the specified dijkstra properties in the returned DijkstraOptionFunc.
func WithDijkstraStandard() DijkstraOptionFunc {
	return func(properties *DijkstraProperties) {
		properties.useStandard = true
	}
}

// dijkstraVertex represents dijkstra vertex.
type dijkstraVertex[T comparable] struct {
	label    T
	dist     float64
	visited  bool
	previous T
}

func newDijkstraVertex[T comparable](label T) *dijkstraVertex[T] {
	return &dijkstraVertex[T]{
		label:   label,
		dist:    math.MaxFloat64,
		visited: false,
	}
}

// Dijkstra has two implementation of Dijkstra's algorithm.
//
// The first one is simple dijkstra, It's using
// a simple slice to keep track of unvisited vertices, and selected the vertex
// with the smallest tentative distance using linear search.
//
// And the second is standard dijkstra, That uses a
// min heap as a priority queue to find the shortest path between start vertex
// and all other vertices in the specified graph.
//
// The time complexity of the simple Dijkstra's algorithm implementation is O(V^2).
//
// It returns the shortest distances from the starting vertex to all other vertices
// in the graph.
func Dijkstra[T comparable](g grafik.Grafik[T], start T, opts ...DijkstraOptionFunc) map[T]float64 {
	var properties DijkstraProperties
	for _, opt := range opts {
		opt(&properties)
	}

	dist := make(map[T]float64)

	startVertex := g.GetVertexByLabel(start)
	if startVertex == nil {
		return dist
	}

	// useStandard checker
	if !properties.useStandard {
		vertices := g.GetAllVertices()
		for _, v := range vertices {
			dist[v.Label()] = math.MaxFloat64
		}

		dist[start] = 0
		visited := make(map[T]bool)
		for len(visited) < len(vertices) {
			var u *grafik.Vertex[T]
			for _, v := range vertices {
				if !visited[v.Label()] && (u == nil || dist[v.Label()] < dist[u.Label()]) {
					u = v
				}
			}

			visited[u.Label()] = true

			neighbors := u.Neighbors()
			for _, neighbor := range neighbors {
				edge := g.GetEdge(u, neighbor)
				if alt := dist[u.Label()] + edge.Weight(); alt < dist[edge.Destination().Label()] {
					dist[edge.Destination().Label()] = alt
				}
			}
		}

		return dist
	}

	// Initialize the heap and the visited map
	pq := queue.NewVertexPriorityQueue[T]()
	visited := make(map[T]bool)

	// Initialize the start vertex
	dVertices := make(map[T]*dijkstraVertex[T])
	vertices := g.GetAllVertices()
	for _, v := range vertices {
		dVertices[v.Label()] = newDijkstraVertex(v.Label())
	}

	dVertices[start].dist = 0

	// Add the start vertex to the heap
	pq.Push(queue.NewVertexWithPriority(g.GetVertexByLabel(start), dVertices[start].dist))

	// Main loop
	for pq.Len() > 0 {
		// Extract the vertex with the smallest tentative distance from the heap
		curr := pq.Pop()
		visited[curr.Vertex().Label()] = true

		// Update the distances of its neighbors
		neighbors := curr.Vertex().Neighbors()
		for i, v := range neighbors {
			if !visited[v.Label()] {
				neighbor := dVertices[v.Label()]
				newDist := curr.Priority() + g.GetEdge(curr.Vertex(), v).Weight()
				if newDist < neighbor.dist {
					neighbor.dist = newDist
					neighbor.previous = curr.Vertex().Label()
					pq.Push(queue.NewVertexWithPriority(neighbors[i], dVertices[v.Label()].dist))
				}
			}
		}
	}

	// Return the distances from the start vertex to each other vertex
	distances := make(map[T]float64)
	for _, v := range dVertices {
		distances[v.label] = v.dist
	}

	return distances
}
