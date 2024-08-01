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

package centrality

import (
	"testing"

	"github.com/fitm-elite/grafik"
	"github.com/fitm-elite/grafik/pathfinder"
)

func TestDijkstraCentrality(t *testing.T) {
	g := grafik.New[string]()

	vA := g.AddVertexByLabel("A")
	vB := g.AddVertexByLabel("B")
	vC := g.AddVertexByLabel("C")
	vE := g.AddVertexByLabel("E")
	vJ := g.AddVertexByLabel("J")
	vK := g.AddVertexByLabel("K")
	vQ := g.AddVertexByLabel("Q")

	_, _ = g.AddEdge(vA, vB, grafik.WithEdgeWeight(1))
	_, _ = g.AddEdge(vA, vE, grafik.WithEdgeWeight(5))
	_, _ = g.AddEdge(vA, vJ, grafik.WithEdgeWeight(2))

	_, _ = g.AddEdge(vB, vC, grafik.WithEdgeWeight(3))
	_, _ = g.AddEdge(vB, vE, grafik.WithEdgeWeight(1))

	_, _ = g.AddEdge(vE, vC, grafik.WithEdgeWeight(4))
	_, _ = g.AddEdge(vE, vJ, grafik.WithEdgeWeight(3))
	_, _ = g.AddEdge(vE, vQ, grafik.WithEdgeWeight(3))
	_, _ = g.AddEdge(vE, vK, grafik.WithEdgeWeight(2))

	_, _ = g.AddEdge(vJ, vQ, grafik.WithEdgeWeight(1))

	_, _ = g.AddEdge(vQ, vK, grafik.WithEdgeWeight(2))

	paths := DijkstraCentrality(g, pathfinder.WithDijkstraStandard())

	if len(paths) != 7 {
		t.Errorf("Expected len from paths is %d, got %d", 7, len(paths))
	}

	if paths[0].Label() == vE.Label() && paths[0].AverageLength() == 2.14 {
		t.Errorf("Expected %s (%.2f), got %s (%.2f)", vE.Label(), 2.14, paths[0].Label(), paths[0].AverageLength())
	}

	if paths[1].Label() == vB.Label() && paths[1].AverageLength() == 2.14 {
		t.Errorf("Expected %s (%.2f), got %s (%.2f)", vB.Label(), 2.14, paths[0].Label(), paths[0].AverageLength())
	}
}
