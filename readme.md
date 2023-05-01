[![CI](https://github.com/nikitagriaznov/plt/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/nikitagriaznov/plt/actions/workflows/ci.yml)
# plt - Svg Plot Making Library for Go
plt is a simple to use and low weight library to make a svg plot
It contains two entry points:
- `plot, err := plt.DrawAngular(PlotHeight, PlotWidth, xDivisionsQty, yDivisionsQty, xArray, yArray, NameOfX, NameOfY)`
- `plot, err := plt.DrawAngularFrom0(PlotHeight, PlotWidth, xDivisionsQty, yDivisionsQty, xArray, yArray, NameOfX, NameOfY)`
## Parameters description
**TotalHeight** and **TotalWidth** is the size of resulting picture in px.

**xDivisionsQty** and **yDivisionsQty** is the number of divisions on the x and y scale.

**xArray** and **yArray** is the parallel arrays of x and y coordinates, it should be equal length. Following example describes parabola plot
```go
xArray:=[]int{-3, -2, -1, 0, 1, 2, 3}
yArray:=[]int{ 9,  4,  1, 0, 1, 4, 9}
```
**NameOfX** and **NameOfY** is the axis labels

## DrawAngular
DrawAngular makes a plot where points are joined with strait lines
The greed of the plot starts from the smallest values of xArray and yArray

## DrawAngularFrom0
DrawAngular makes a plot where points are joined with strait lines
The greed of the plot starts from zero point

## Usage example

```go
package main

import (
	"github.com/nikitagryaznov/plt"
	"log"
	"os"
)

func main() {
	totalHeight := uint(200)
	totalWidth := uint(400)
	xDivisionsQty := uint(10)
	yDivisionsQty := uint(10)
	xLabel := string("x")
	yLabel := string("y")
	xArray := []float64{-3, -2, -1, 0, 1, 2, 3}
	yArray := []float64{9, 4, 1, 0, 1, 4, 9}
	plot, err := plt.DrawAngular(totalHeight, totalWidth, xDivisionsQty, yDivisionsQty, xArray, yArray, xLabel, yLabel)
	if err != nil {
		log.Fatalln(err)
	}
	err = os.WriteFile("plot.svg", plot, 777)
	if err != nil {
		log.Fatalln(err)
	}
}
```