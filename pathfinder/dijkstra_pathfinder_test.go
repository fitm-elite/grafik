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
	"testing"

	"github.com/fitm-elite/grafik"
)

func TestSimpleDijkstra(t *testing.T) {
	g := grafik.New[string]()

	vA := g.AddVertexByLabel("A")
	vB := g.AddVertexByLabel("B")
	vC := g.AddVertexByLabel("C")
	vD := g.AddVertexByLabel("D")

	_, _ = g.AddEdge(vA, vB, grafik.WithEdgeWeight(4))
	_, _ = g.AddEdge(vA, vC, grafik.WithEdgeWeight(3))
	_, _ = g.AddEdge(vB, vC, grafik.WithEdgeWeight(1))
	_, _ = g.AddEdge(vB, vD, grafik.WithEdgeWeight(2))
	_, _ = g.AddEdge(vC, vD, grafik.WithEdgeWeight(4))

	// use not existing vertex
	dist := Dijkstra(g, "X")
	if len(dist) > 0 {
		t.Errorf("Expected dist map length be 0, got %d", len(dist))
	}

	dist = Dijkstra(g, "A")

	if dist[vA.Label()] != 0 {
		t.Errorf("Expected distance from A to %s to be 0, got %f", vA.Label(), dist[vA.Label()])
	}
	if dist[vB.Label()] != 4 {
		t.Errorf("Expected distance from A to %s  to be 4, got %f", vB.Label(), dist[vB.Label()])
	}
	if dist[vC.Label()] != 3 {
		t.Errorf("Expected distance from A to %s  to be 3, got %f", vC.Label(), dist[vC.Label()])
	}
	if dist[vD.Label()] != 6 {
		t.Errorf("Expected distance from A to %s  to be 6, got %f", vD.Label(), dist[vD.Label()])
	}
}

func TestStandardDijkstra(t *testing.T) {
	g := grafik.New[int]()

	v1 := g.AddVertexByLabel(1)
	v2 := g.AddVertexByLabel(2)
	v3 := g.AddVertexByLabel(3)
	v4 := g.AddVertexByLabel(4)

	_, _ = g.AddEdge(v1, v2, grafik.WithEdgeWeight(4))
	_, _ = g.AddEdge(v1, v3, grafik.WithEdgeWeight(3))
	_, _ = g.AddEdge(v2, v3, grafik.WithEdgeWeight(1))
	_, _ = g.AddEdge(v2, v4, grafik.WithEdgeWeight(2))
	_, _ = g.AddEdge(v3, v4, grafik.WithEdgeWeight(4))

	// use not existing vertex
	dist := Dijkstra(g, 0, WithDijkstraStandard())
	if len(dist) > 0 {
		t.Errorf("Expected dist map length be 0, got %d", len(dist))
	}

	dist = Dijkstra(g, 1)

	if dist[v1.Label()] != 0 {
		t.Errorf("Expected distance from 1 to %d to be 0, got %f", v1.Label(), dist[v1.Label()])
	}
	if dist[v2.Label()] != 4 {
		t.Errorf("Expected distance from 1 to %d to be 4, got %f", v2.Label(), dist[v2.Label()])
	}
	if dist[v3.Label()] != 3 {
		t.Errorf("Expected distance from 1 to %d to be 3, got %f", v3.Label(), dist[v3.Label()])
	}
	if dist[v4.Label()] != 6 {
		t.Errorf("Expected distance from 1 to %d to be 6, got %f", v4.Label(), dist[v4.Label()])
	}
}
