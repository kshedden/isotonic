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

func gendat() ([]float64, []float64) {

	n := 1000
	x := make([]float64, n)
	y := make([]float64, n)
	for i := range x {
		x[i] = 16*float64(i)/float64(n-1) - 8
		e := rand.NormFloat64()
		y[i] = math.Exp(x[i])/(1+math.Exp(x[i])) + 0.05*e
	}

	return x, y
}

func main() {

	x, y := gendat()
	qi := isotonic.NewIsotonicRegression(x, y, true)
	qy := qi.Fit()

	pts := make(plotter.XYs, len(x))
	fpts := make(plotter.XYs, len(x))
	for i := range qy {
		pts[i].X = x[i]
		pts[i].Y = y[i]
		fpts[i].X = x[i]
		fpts[i].Y = qy[i]
	}

	plt, err := plot.New()
	if err != nil {
		panic(err)
	}

	s, err := plotter.NewScatter(pts)
	if err != nil {
		panic(err)
	}
	plt.Add(s)

	l, err := plotter.NewLine(fpts)
	if err != nil {
		panic(err)
	}
	l.LineStyle.Width = vg.Points(3)
	l.LineStyle.Color = color.RGBA{R: 255, G: 40, B: 40, A: 255}
	plt.Add(l)

	plt.X.Label.Text = "X"
	plt.Y.Label.Text = "Y"

	err = plt.Save(4*vg.Inch, 4*vg.Inch, "iso.pdf")
	if err != nil {
		panic(err)
	}
}
