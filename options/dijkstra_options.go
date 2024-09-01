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

// DijkstraOptionFunc represent an alias of function type that modifies the specified dijkstra properties.
type DijkstraOptionFunc func(properties *DijkstraProperties)

// DijkstraProperties represents the properties of an dijkstra.
type DijkstraProperties struct {
	useStandard bool
}

// UseStandard set use standard to true
func (dj *DijkstraProperties) UseStandard() {
	dj.useStandard = true
}

// GetUseStandard return dj.useStandard from DijkstraProperties.
func (dj DijkstraProperties) GetUseStandard() bool {
	return dj.useStandard
}

// WithDijkstraStandard sets the standard algorithm for the specified dijkstra properties in the returned DijkstraOptionFunc.
func WithDijkstraStandard() DijkstraOptionFunc {
	return func(properties *DijkstraProperties) {
		properties.useStandard = true
	}
}
