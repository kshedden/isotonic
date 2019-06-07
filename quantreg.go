// https://arxiv.org/pdf/1901.02398.pdf

package isotonic

import (
	"math"

	"gonum.org/v1/gonum/floats"
)

// Estimate a conditional quantile of y given x under monotonicity
// constraints.
type QuantReg struct {

	// The covariate
	x []float64

	// The response
	y []float64

	// If true, P(Y<y | X=x) is increasing in x, for fixed y, else
	// it is decreasing.
	increasing bool

	// Number of equally-spaced points between min(y) and max(y).
	npt int

	f  [][]float64
	qp []float64
}

// NewQuantReg creates a QuantReg value for estimating conditional
// quantiles of y given x.  The value of npt determines the number of
// points at which the each CDF is evaluated to approximate the
// conditional quantiles.
func NewQuantReg(x, y []float64, increasing bool, npt int) *QuantReg {

	x, y = sortTwo(x, y)

	qr := &QuantReg{
		y:          y,
		x:          x,
		increasing: increasing,
		npt:        npt,
	}

	qr.fitiso()

	return qr
}

// Fit returns estimates of the p^th conditional quantile
// of Y given X, evaluated at the training data points.
func (qr *QuantReg) Fit(p float64) []float64 {

	n := len(qr.x)
	qtx := make([]float64, n)

	for i := range qr.x {

		var i0 int
		var v float64
		for j := range qr.f {
			d := math.Abs(qr.f[j][i] - p)
			if j == 0 || d < v {
				i0 = j
				v = d
			}
		}
		qtx[i] = qr.qp[i0]
	}

	return qtx
}

func (qr *QuantReg) fitiso() {

	ymin := floats.Min(qr.y)
	ymax := floats.Max(qr.y)
	ind := make([]float64, len(qr.y))
	qp := make([]float64, qr.npt)
	qr.f = qr.f[0:0]

	for j := 0; j < qr.npt; j++ {
		qp[j] = ymin + (ymax-ymin)*float64(j)/(float64(qr.npt-1))

		for i := range qr.y {
			if qr.y[i] <= qp[j] {
				ind[i] = 1
			} else {
				ind[i] = 0
			}
		}

		// These functions are decreasing in y.
		ir := NewIsotonicRegression(qr.x, ind, !qr.increasing)
		yf := ir.Fit()
		qr.f = append(qr.f, yf)
	}

	qr.qp = qp
}
