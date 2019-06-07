package isotonic

import (
	"gonum.org/v1/gonum/floats"
)

// IsotonicRegression supports fitting a monotone relationship between
// Y (the dependent variable) and X (the independent variable).  If
// X is not sorted, call the Sort method before calling Fit.
type IsotonicRegression struct {
	x          []float64
	y          []float64
	increasing bool
}

func NewIsotonicRegression(x, y []float64, increasing bool) *IsotonicRegression {

	x, y = sortTwo(x, y)
	return &IsotonicRegression{
		x:          x,
		y:          y,
		increasing: increasing,
	}
}

func sortTwo(x, y []float64) ([]float64, []float64) {
	inds := make([]int, len(x))
	floats.Argsort(x, inds)
	z := make([]float64, len(y))
	for j, i := range inds {
		z[j] = y[i]
	}
	return x, z
}

func reverse(x []float64) {
	for i := len(x)/2 - 1; i >= 0; i-- {
		j := len(x) - 1 - i
		x[i], x[j] = x[j], x[i]
	}
}

// Reverse reverses Y and X in the regression struct.
func (ir *IsotonicRegression) Reverse() {
	for _, a := range [][]float64{ir.x, ir.y} {
		reverse(a)
	}
}

func (ir *IsotonicRegression) Fit() []float64 {

	if ir.increasing {
		return ir.fitNonDecreasing()
	}

	ir.Reverse()
	y := ir.fitNonDecreasing()
	ir.Reverse()
	reverse(y)
	return y
}

func (ir *IsotonicRegression) fitNonDecreasing() []float64 {

	y := ir.y
	n := len(y)
	w := make([]float64, n)
	z := make([]float64, n)
	s := make([]int, n+1)

	z[0] = y[0]
	s[0] = -1
	w[0] = 1
	j := 0
	for i := 1; i < n; i++ {
		j++
		z[j] = y[i]
		w[j] = 1
		for j > 0 && z[j] < z[j-1] {
			z[j-1] = (w[j]*z[j] + w[j-1]*z[j-1]) / (w[j] + w[j-1])
			w[j-1] += w[j]
			j--
		}
		s[j+1] = i
	}

	yf := make([]float64, n)
	for k := 0; k <= j; k++ {
		for i := s[k] + 1; i <= s[k+1]; i++ {
			yf[i] = z[k]
		}
	}

	return yf
}
