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

import (
	"errors"
	"reflect"
	"testing"

	"github.com/fitm-elite/grafik"
)

func TestBreadFirstIterator(t *testing.T) {
	g := grafik.New[string]()

	/*
		     *  TODO: Create a graph
		     *  & And use the iterator
		     *
			 *  A -> B -> C
			 *  |    |    |
			 *  v    v    v
			 *  D -> E -> F
	*/

	vertices := map[string]*grafik.Vertex[string]{
		"A": g.AddVertexByLabel("A"),
		"B": g.AddVertexByLabel("B"),
		"C": g.AddVertexByLabel("C"),
		"D": g.AddVertexByLabel("D"),
		"E": g.AddVertexByLabel("E"),
		"F": g.AddVertexByLabel("F"),
	}

	// add some edges
	_, _ = g.AddEdge(vertices["A"], vertices["B"])
	_, _ = g.AddEdge(vertices["A"], vertices["D"])
	_, _ = g.AddEdge(vertices["B"], vertices["C"])
	_, _ = g.AddEdge(vertices["B"], vertices["E"])
	_, _ = g.AddEdge(vertices["C"], vertices["F"])
	_, _ = g.AddEdge(vertices["D"], vertices["E"])
	_, _ = g.AddEdge(vertices["E"], vertices["F"])

	// create an iterator with a vertex that doesn't exist
	_, err := NewBreadthFirstIterator(g, "X")
	if err == nil {
		t.Error("Expect NewBreadthFirstIterator returns error, but got nil")
	}

	// test depth first iteration
	iterator, err := NewBreadthFirstIterator[string](g, "A")
	if err != nil {
		t.Errorf("Expect NewBreadthFirstIterator doesn't return error, but got %s", err)
	}

	expected := []string{"A", "B", "D", "C", "E", "F"}
	for i, label := range expected {
		if !iterator.HasNext() {
			t.Errorf("Expected iterator.HasNext() to be true, but it was false for label %s", label)
		}

		v := iterator.Next()
		if v.Label() != expected[i] {
			t.Errorf("Expected iterator.Next().Label() to be %s, but got %s", expected[i], v.Label())
		}
	}

	if iterator.HasNext() {
		t.Error("Expected iter.HasNext() to be false, but it was true")
	}

	v := iterator.Next()
	if v != nil {
		t.Errorf("Expected nil, but got %+v", v)
	}

	// test the Reset method
	iterator.Reset()
	if !iterator.HasNext() {
		t.Error("Expected iter.HasNext() to be true, but it was false after reset")
	}

	v = iterator.Next()
	if v.Label() != "A" {
		t.Errorf("Expected iter.Next().Label() to be %s, but got %s", "A", v.Label())
	}

	// test Iterate method
	iterator.Reset()
	var ordered []string
	err = iterator.Iterate(func(vertex *grafik.Vertex[string]) error {
		ordered = append(ordered, vertex.Label())
		return nil
	})
	if err != nil {
		t.Errorf("Expect iter.Iterate(func) returns no error, but got one %s", err)
	}

	if !reflect.DeepEqual(expected, ordered) {
		t.Errorf("Expect same vertex order, but got different one expected: %v, actual: %v", expected, ordered)
	}

	iterator.Reset()
	expectedErr := errors.New("something went wrong")
	err = iterator.Iterate(func(vertex *grafik.Vertex[string]) error {
		return expectedErr
	})
	if err == nil {
		t.Error("Expect iter.Iterate(func) returns error, but got nil")
	}

	if !errors.Is(err, expectedErr) {
		t.Errorf("Expect %+v error, but got %+v", expectedErr, err)
	}
}
