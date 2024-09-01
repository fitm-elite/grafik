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
	"sort"
	"sync"

	"github.com/fitm-elite/grafik"
	"github.com/fitm-elite/grafik/entity"
	"github.com/fitm-elite/grafik/options"
	"github.com/fitm-elite/grafik/pathfinder"
)

// DijkstraCentrality It's using a dijkstra method to find shortest path in each vertex
// and calculate to find an average value in each path to find a centroid.
//
// Returns []VertexPath[T]
func DijkstraCentrality[T comparable](g entity.Grafik[T], opts ...options.DijkstraOptionFunc) []grafik.VertexPath[T] {
	vertices := g.GetAllVertices()
	vertexPaths := make([]grafik.VertexPath[T], 0, len(vertices))

	var mu sync.Mutex
	var wg sync.WaitGroup

	results := make(chan grafik.VertexPath[T], len(vertices))

	wg.Add(len(vertices))
	for _, v := range vertices {
		go func(v *grafik.Vertex[T]) {
			defer wg.Done()
			label := v.Label()
			pathLengths := pathfinder.Dijkstra(g, label, opts...)

			var totalLength float64
			for _, length := range pathLengths {
				totalLength += length
			}

			averageLength := totalLength / float64(len(pathLengths))
			result := grafik.VertexPath[T]{
				VertexLabel:   label,
				AverageLength: averageLength,
			}

			results <- result
		}(v)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		mu.Lock()
		vertexPaths = append(vertexPaths, result)
		mu.Unlock()
	}

	sort.Slice(vertexPaths, func(i, j int) bool {
		return vertexPaths[i].AverageLength < vertexPaths[j].AverageLength
	})

	return vertexPaths
}
