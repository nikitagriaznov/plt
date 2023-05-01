[![CI](https://github.com/nikitagriaznov/plt/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/nikitagriaznov/plt/actions/workflows/ci.yml)
# svgPlot - Svg Plot Making Library for Go
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

Parameters can be stored in Style type variable

```go
type Style struct {
	TotalHeight, TotalWidth      uint   // Size of resulting picture
	XDivisionsQty, YDivisionsQty uint   // Required number of divisions on X and Y scale
	NameOfX, NameOfY             string // Names of X and Y axis. Hidden if empty
}
```

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
	"github.com/nikk-gr/svgPlot"
	"log"
	"os"
)

func main() {
	TotalHeight := uint(200)
	TotalWidth := uint(400)
	xDivisionsQty := uint(10)
	yDivisionsQty := uint(10)
	NameOfX := string("x")
	NameOfY := string("y")
	xArray := []int{-3, -2, -1, 0, 1, 2, 3}
	yArray := []int{9, 4, 1, 0, 1, 4, 9}
	plot, err := svgPlot.DrawAngular(TotalHeight, TotalWidth, xDivisionsQty, yDivisionsQty, xArray, yArray, NameOfX, NameOfY)
	if err != nil {
		log.Fatalln(err)
	}
	err = os.WriteFile("plot1.svg", []byte(plot), 777)
	if err != nil {
		log.Fatalln(err)
	}
}
```
or
```go
package main

import (
	"github.com/nikk-gr/svgPlot"
	"log"
	"os"
)

func main() {
	plotStyle:= svgPlot.Style{
		Height: 200,
		Width: 400,
		XDivisionsQty: 10,
		YDivisionsQty: 10,
		NameOfX: "x",
		NameOfY: "y",
    }
	xArray := []int{-3, -2, -1, 0, 1, 2, 3}
	yArray := []int{9, 4, 1, 0, 1, 4, 9}
	plot, err := plotStyle.DrawAngularInt(xArray, yArray)
	if err != nil {
		log.Fatalln(err)
	}
	err = os.WriteFile("plot1.svg", []byte(plot), 777)
	if err != nil {
		log.Fatalln(err)
	}
}
```
![result](./.github/img/plot1.svg "result")