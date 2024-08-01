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
	"reflect"
	"slices"
	"testing"
)

const (
	testErrMsgError    = "Expected no error, but got %s"
	testErrMsgNoError  = "Expected error, but got no error"
	testErrMsgWrongLen = "Expected len %d, but got %d"
	testErrMsgNotFalse = "Expected false, but got true"
	testErrMsgNotTrue  = "Expected true, but got false"
	testErrMsgNotEqual = "Expected %+v, but got %+v"
)

func TestAddVertexByLabel(t *testing.T) {
	g := New[string]()

	vA := g.AddVertexByLabel("A")

	if !g.ContainsVertex(vA) {
		t.Error(testErrMsgNotTrue)
	}
}

func TestGetVertexByLabel(t *testing.T) {
	g := New[string]()

	vA := g.AddVertexByLabel("A")

	v := g.GetVertexByLabel("A")
	if v.Label() != "A" {
		t.Errorf(testErrMsgNotEqual, vA, v)
	}
}

func TestGetAllVertices(t *testing.T) {
	g := New[string]()

	labels := []string{"A", "B", "C", "D", "E", "F"}
	for _, label := range labels {
		_ = g.AddVertexByLabel(label)
	}

	vertices := g.GetAllVertices()
	if len(vertices) != len(labels) {
		t.Errorf(testErrMsgNotEqual, len(labels), len(vertices))
	}

	for _, vertex := range vertices {
		if matched := slices.Contains(labels, vertex.Label()); !matched {
			t.Error(testErrMsgNotTrue)
		}
	}
}

func TestAddEdge(t *testing.T) {
	g := New[string]()

	vA := g.AddVertexByLabel("A")
	vB := g.AddVertexByLabel("B")

	_, err := g.AddEdge(vA, vB)
	if err != nil {
		t.Errorf(testErrMsgError, err)
	}

	if !g.ContainsEdge(vA, vB) {
		t.Error(testErrMsgNotTrue)
	}
}

func TestGetAllEdges(t *testing.T) {
	g := New[string]()

	vertices := map[string]*Vertex[string]{
		"A": g.AddVertexByLabel("A"),
		"B": g.AddVertexByLabel("B"),
		"C": g.AddVertexByLabel("C"),
	}

	// add some edges
	_, err := g.AddEdge(vertices["A"], vertices["B"])
	if err != nil {
		t.Errorf(testErrMsgError, err)
	}

	_, err = g.AddEdge(vertices["A"], vertices["C"])
	if err != nil {
		t.Errorf(testErrMsgError, err)
	}

	edges := g.GetAllEdges(vertices["A"], vertices["B"])
	if len(edges) != 2 {
		t.Errorf(testErrMsgWrongLen, 2, len(edges))
	}

	edges = g.GetAllEdges(vertices["B"], vertices["A"])
	if len(edges) != 2 {
		t.Errorf(testErrMsgWrongLen, 2, len(edges))
	}

	edges = g.GetAllEdges(vertices["B"], vertices["C"])
	if len(edges) != 0 {
		t.Errorf(testErrMsgWrongLen, 0, len(edges))
	}

	edges = g.GetAllEdges(nil, vertices["B"])
	if edges != nil {
		t.Errorf("Expected nil, but got %+v", edges)
	}

	edges = g.GetAllEdges(vertices["B"], NewVertex("D"))
	if edges != nil {
		t.Errorf("Expected nil, but got %+v", edges)
	}

	edges = g.GetAllEdges(NewVertex("D"), vertices["A"])
	if edges != nil {
		t.Errorf("Expected nil, but got %+v", edges)
	}
}

func TestNeighbors(t *testing.T) {
	g := New[string]()

	vA := g.AddVertexByLabel("A")
	vB := g.AddVertexByLabel("B")
	vC := g.AddVertexByLabel("C")

	_, err := g.AddEdge(vA, vB)
	if err != nil {
		t.Errorf(testErrMsgError, err)
	}

	_, err = g.AddEdge(vA, vC)
	if err != nil {
		t.Errorf(testErrMsgError, err)
	}

	v := vA.NeighborByLabel("B")
	if !reflect.DeepEqual(vB, v) {
		t.Errorf(testErrMsgNotEqual, vB, v)
	}

	if !vA.HasNeighbor(vC) {
		t.Error(testErrMsgNotTrue)
	}

	if vA.HasNeighbor(NewVertex("D")) {
		t.Error(testErrMsgNotFalse)
	}

	if vA.OutDegree() != 2 {
		t.Errorf(testErrMsgNotEqual, 2, vA.OutDegree())
	}

	// test cloning neighbors
	neighbors := vA.Neighbors()
	if len(neighbors) != len(vA.neighbors) {
		t.Errorf(testErrMsgNotEqual, len(neighbors), len(vA.neighbors))
	}

	neighbors[0].label = "D"
	if neighbors[0].Label() == vA.neighbors[0].Label() {
		t.Errorf(testErrMsgNotFalse)
	}
}
