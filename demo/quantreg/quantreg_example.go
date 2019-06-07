package main

import (
	"image/color"
	"math"
	"math/rand"

	"github.com/kshedden/isotonic"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func gendat(n int, incr bool) ([]float64, []float64) {

	x := make([]float64, n)
	y := make([]float64, n)
	for i := range x {
		x[i] = 2*float64(i)/float64(n-1) - 1
		e := math.Abs(x[i]) * rand.NormFloat64()
		if incr {
			y[i] = x[i]*x[i]*x[i] + 0.5*e
		} else {
			y[i] = -x[i]*x[i]*x[i] + 0.5*e
		}
	}

	return x, y
}

func main() {

	for _, incr := range []bool{false, true} {

		x, y := gendat(10000, incr)
		qr := isotonic.NewQuantReg(x, y, incr, 1000)

		plt, err := plot.New()
		if err != nil {
			panic(err)
		}

		pts := make(plotter.XYs, len(x))
		for i := range y {
			pts[i].X = x[i]
			pts[i].Y = y[i]
		}

		s, err := plotter.NewScatter(pts)
		if err != nil {
			panic(err)
		}
		plt.Add(s)

		for k, p := range []float64{0.25, 0.75} {

			qy := qr.Fit(p)

			fpts := make(plotter.XYs, len(x))
			for i := range qy {
				fpts[i].X = x[i]
				fpts[i].Y = qy[i]
			}

			l, err := plotter.NewLine(fpts)
			if err != nil {
				panic(err)
			}
			l.LineStyle.Width = vg.Points(3)
			if k == 0 {
				l.LineStyle.Color = color.RGBA{R: 255, G: 40, B: 40, A: 255}
			} else {
				l.LineStyle.Color = color.RGBA{R: 40, G: 40, B: 255, A: 255}
			}
			plt.Add(l)
		}

		plt.X.Label.Text = "X"
		plt.Y.Label.Text = "Y"

		fname := "qr_incr.pdf"
		if !incr {
			fname = "qr_decr.pdf"
		}

		err = plt.Save(4*vg.Inch, 4*vg.Inch, fname)
		if err != nil {
			panic(err)
		}
	}
}
