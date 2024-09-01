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
	"testing"

	"github.com/fitm-elite/grafik/options"
)

func TestVertexLabelFunc(t *testing.T) {
	label := "A"

	vA := NewVertex(label)

	vALabel := vA.Label()
	if vALabel != label {
		t.Errorf(testErrMsgNotEqual, label, vALabel)
	}
}

func TestVertexWeightFunc(t *testing.T) {
	label := "B"
	weight := 4.00

	vB := NewVertex(label, options.WithVertexWeight(weight))

	vBWeight := vB.Weight()
	if vBWeight != weight {
		t.Errorf(testErrMsgNotEqual, weight, vBWeight)
	}
}
