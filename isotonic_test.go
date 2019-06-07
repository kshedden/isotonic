package isotonic

import (
	"testing"

	"gonum.org/v1/gonum/floats"
)

func TestReverse(t *testing.T) {

	for _, m := range []struct {
		x  []float64
		y  []float64
		xr []float64
		yr []float64
	}{
		{
			x:  []float64{1, 2, 3, 4, 5, 6},
			y:  []float64{3, 1, 5, 4, 6, 8},
			xr: []float64{6, 5, 4, 3, 2, 1},
			yr: []float64{8, 6, 4, 5, 1, 3},
		},
		{
			x:  []float64{1, 2, 3, 4},
			y:  []float64{5, 4, 3, 2},
			xr: []float64{4, 3, 2, 1},
			yr: []float64{2, 3, 4, 5},
		},
		{
			x:  []float64{1, 2, 3, 4},
			y:  []float64{5, 4, 3, 2},
			xr: []float64{4, 3, 2, 1},
			yr: []float64{2, 3, 4, 5},
		},
		{
			x:  []float64{4, 7, 9},
			y:  []float64{2, 3, 4},
			xr: []float64{9, 7, 4},
			yr: []float64{4, 3, 2},
		},
	} {
		ir := NewIsotonicRegression(m.x, m.y, true)
		ir.Reverse()

		if !floats.EqualApprox(ir.x, m.xr, 1e-8) {
			t.Fail()
		}
		if !floats.EqualApprox(ir.y, m.yr, 1e-8) {
			t.Fail()
		}
	}
}

func TestIso(t *testing.T) {

	for _, m := range []struct {
		x          []float64
		y          []float64
		f          []float64
		increasing bool
	}{
		{
			x:          []float64{1, 2, 3, 4, 5, 6},
			y:          []float64{2, 3, 4, 5, 6, 7},
			f:          []float64{2, 3, 4, 5, 6, 7},
			increasing: true,
		},
		{
			x:          []float64{1, 2, 3, 4, 5},
			y:          []float64{2, 3, 4, 5, 6},
			f:          []float64{2, 3, 4, 5, 6},
			increasing: true,
		},
		{
			x:          []float64{1, 2, 3, 4, 5, 6},
			y:          []float64{7, 6, 5, 4, 3, 2},
			f:          []float64{4.5, 4.5, 4.5, 4.5, 4.5, 4.5},
			increasing: true,
		},
		{
			x:          []float64{1, 2, 3, 4, 5},
			y:          []float64{6, 5, 4, 3, 2},
			f:          []float64{4, 4, 4, 4, 4},
			increasing: true,
		},
		{
			x:          []float64{1, 2, 3, 4, 5, 6},
			y:          []float64{2, 3, 6, 5, 4, 1},
			f:          []float64{2, 3, 4, 4, 4, 4},
			increasing: true,
		},
		{
			x:          []float64{1, 2, 3, 4, 5, 6},
			y:          []float64{5, 4, 6, 5, 7, 8},
			f:          []float64{4.5, 4.5, 5.5, 5.5, 7, 8},
			increasing: true,
		},
		{
			x:          []float64{1, 2, 3, 4, 5, 6},
			y:          []float64{6, 5, 8, 7, 9, 9},
			f:          []float64{5.5, 5.5, 7.5, 7.5, 9, 9},
			increasing: true,
		},
		{
			x:          []float64{1, 2, 3, 4, 5, 6},
			y:          []float64{1, 5, 4, 6, 6, 5},
			f:          []float64{1, 4.5, 4.5, 5 + 2.0/3.0, 5 + 2.0/3.0, 5 + 2.0/3.0},
			increasing: true,
		},
		{
			x:          []float64{1, 2, 3, 4, 5, 6},
			y:          []float64{1, 2, 1, 3, 5, 4},
			f:          []float64{1, 1.5, 1.5, 3, 4.5, 4.5},
			increasing: true,
		},
		{
			x:          []float64{1, 2, 3, 4, 5, 6},
			y:          []float64{6, 5, 4, 3, 2, 1},
			f:          []float64{6, 5, 4, 3, 2, 1},
			increasing: false,
		},
		{
			x:          []float64{1, 2, 3, 4, 5, 6},
			y:          []float64{6, 4, 5, 2, 3, 2},
			f:          []float64{6, 4.5, 4.5, 2.5, 2.5, 2},
			increasing: false,
		},
	} {
		ir := NewIsotonicRegression(m.x, m.y, m.increasing)

		f := ir.Fit()
		if !floats.EqualApprox(f, m.f, 1e-6) {
			t.Fail()
		}
	}
}
