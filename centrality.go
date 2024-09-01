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
	"github.com/fitm-elite/grafik/options"
)

// WithDijkstraStandard sets the standard algorithm for the specified dijkstra properties in the returned DijkstraOptionFunc.
func WithDijkstraStandard() options.DijkstraOptionFunc {
	return func(properties *options.DijkstraProperties) {
		properties.UseStandard()
	}
}

// DijkstraCentrality return slice of Vertex that has result of centrality.
func DijkstraCentrality[T comparable](g Grafik[T], opts ...options.DijkstraOptionFunc) []Vertex[T] {
	return DijkstraCentrality(g, opts...)
}
