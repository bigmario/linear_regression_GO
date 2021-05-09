package main

import (
	"fmt"
	"image/color"

	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

type xy struct {
	x []float64
	y []float64
}

func (d xy) Len() int {
	return len(d.x)
}

func (d xy) XY(i int) (x, y float64) {
	x = d.x[i]
	y = d.y[i]
	return
}

func summatory(arrayX []float64, arrayY []float64, meanX float64, meanY float64) (float64, float64) {
	resultXX := 0.0
	resultXY := 0.0

	for x := 0; x < len(arrayX); x++ {
		for y := 0; y < len(arrayY); y++ {
			if x == y {
				resultXY += (arrayX[x] - meanX) * (arrayY[y] - meanY)
			}
		}
		resultXX += (arrayX[x] - meanX) * (arrayX[x] - meanX)
	}

	return resultXY, resultXX
}

func estimateB0B1(x []float64, y []float64) (float64, float64) {
	var meanX float64
	var meanY float64
	var sumXY float64
	var sumXX float64

	meanX = stat.Mean(x, nil)
	meanY = stat.Mean(y, nil)

	sumXY, sumXX = summatory(x, y, meanX, meanY)

	b1 := sumXY / sumXX
	b0 := meanY - b1*meanX

	return b0, b1
}

func plotRegression(x []float64, y []float64, b0 float64, b1 float64) {

	data := xy{
		x: x,
		y: y,
	}

	p := plot.New()

	p.Title.Text = "Regresion Lineal"
	p.X.Label.Text = "X - Independiente"
	p.Y.Label.Text = "Y - Dependiente"

	scatter, _ := plotter.NewScatter(data)
	scatter.Color = color.Black

	regression := plotter.NewFunction(func(x float64) float64 { return b0 + b1*x })
	regression.Color = color.RGBA{B: 255, A: 255}

	p.Add(plotter.NewGrid(), regression, scatter)

	p.X.Min = 2011
	p.X.Max = 2022
	p.Y.Min = 1390
	p.Y.Max = 1600

	p.Save(8*vg.Inch, 8*vg.Inch, "RegresionLineal.png")
}

func main() {
	X := []float64{2011, 2012, 2013, 2014, 2015, 2016, 2017, 2018, 2019, 2020}
	Y := []float64{1390, 1408, 1415, 1445, 1432, 1445, 1478, 1495, 1530, 1520}

	b0, b1 := estimateB0B1(X, Y)

	fmt.Printf("Los valores de b0=%v , b1=%v\n", b0, b1)

	plotRegression(X, Y, b0, b1)
}
