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

package options

// VertexOptionFunc represent an alias of function type that modifies the specified vertex properties.
type VertexOptionFunc func(properties *VertexProperties)

// VertexProperties represents the properties of an edge.
type VertexProperties struct {
	weight float64
}

// Weight returns v.weight from VertexProperties
func (v VertexProperties) Weight() float64 {
	return v.weight
}

// WithVertexWeight sets the edge weight for the specified vertex properties in the returned VertexOptionFunc.
func WithVertexWeight(weight float64) VertexOptionFunc {
	return func(properties *VertexProperties) {
		properties.weight = weight
	}
}
