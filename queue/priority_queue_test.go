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
	"reflect"
	"testing"

	"github.com/fitm-elite/grafik"
)

func TestPriorityQueue(t *testing.T) {
	// create a new priority queue
	pq := make(priorityQueue[string], 0)

	// push some items with different priorities to the queue
	heap.Push(&pq, NewVertexWithPriority(grafik.NewVertex("A"), 3))
	heap.Push(&pq, NewVertexWithPriority(grafik.NewVertex("B"), 1))
	heap.Push(&pq, NewVertexWithPriority(grafik.NewVertex("C"), 5))
	heap.Push(&pq, NewVertexWithPriority(grafik.NewVertex("D"), 2))
	heap.Push(&pq, NewVertexWithPriority(grafik.NewVertex("E"), 4))

	// push different type
	heap.Push(&pq, 123)

	// check that the length of the priority queue is 5
	if len(pq) != 5 {
		t.Errorf("priorityQueue length = %d; want 5", len(pq))
	}

	// pop items from the queue and check that they are in the correct order
	items := make([]string, 0)
	for pq.Len() > 0 {
		item := heap.Pop(&pq)
		vp, ok := item.(*VertexWithPriority[string])
		if !ok {
			t.Errorf("Expected *VertexWithPriority[string] type, but got %+v", reflect.TypeOf(item))
		}
		items = append(items, vp.Vertex().Label())
	}
	expected := []string{"B", "D", "A", "E", "C"}
	if !reflect.DeepEqual(items, expected) {
		t.Errorf("PriorityQueue Pop() order = %v; want %v", items, expected)
	}
}

func TestVertexPriorityQueue(t *testing.T) {
	// create a new vertex priority queue
	vpq := NewVertexPriorityQueue[string]()

	// push some items with different priorities to the queue
	vpq.Push(NewVertexWithPriority(grafik.NewVertex("A"), 2))
	vpq.Push(NewVertexWithPriority(grafik.NewVertex("B"), 1))
	vpq.Push(NewVertexWithPriority(grafik.NewVertex("C"), 3))

	if vpq.Peek().vertex.Label() != "B" {
		t.Errorf("Expected Peek returns B, but got %v", vpq.Peek().vertex.Label())
	}

	// check that the length of the priority queue is 5
	if vpq.Len() != 3 {
		t.Errorf("VertexPriorityQueue length = %d; want 5", len(vpq.pq))
	}

	// pop items from the queue and check that they are in the correct order
	items := make([]string, 0)
	priorities := make([]float64, 0)
	for vpq.pq.Len() > 0 {
		item := vpq.Pop()
		items = append(items, item.Vertex().Label())
		priorities = append(priorities, item.Priority())
	}
	expectedVertices := []string{"B", "A", "C"}
	if !reflect.DeepEqual(items, expectedVertices) {
		t.Errorf("VertexPriorityQueue Pop() order = %v; want %v", items, expectedVertices)
	}

	expectedPriorities := []float64{1, 2, 3}
	if !reflect.DeepEqual(priorities, expectedPriorities) {
		t.Errorf("VertexPriorityQueue Pop() order = %v; want %v", priorities, expectedPriorities)
	}

	if vpq.Peek() != nil {
		t.Errorf("Expected Peek returns nil, but got %v", vpq.Peek())
	}
}
