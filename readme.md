
[![CI](https://github.com/nikitagriaznov/plt/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/nikitagriaznov/plt/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/nikk-gr/svgPlot/branch/main/graph/badge.svg?token=2WSYNM93G5)](https://codecov.io/gh/nikk-gr/svgPlot)
![GitHub](https://img.shields.io/github/license/nikk-gr/svgPlot)

# svgPlot - Svg Plot Making Library for Go 
svgPlot is an easy to use and lightweight svg plotting library.
Use the following command to get it:
```bash
go get github.com/nikk-gr/svgPlot@v1.0.0
```
## Parameters description
`TotalHeight` and `TotalWidth` are the sizes of the resulting picture in px.

`xDivisionsQty` and `yDivisionsQty` are the numbers of divisions on the x and y scale.

`X` and `Y` are the parallel arrays of x and y coordinates, it should be equal length. Following example describes parabola plot
```go
X := []float64{-3, -2, -1, 0, 1, 2, 3}
Y := []float64{ 9,  4,  1, 0, 1, 4, 9}
```
`NameOfX` and `NameOfY` are the axis labels. Max len is 6

## Functions description
### ConvertSliceToFloat64
**ConvertSliceToFloat64** convert slices of any `numeric` type to `[]float64`
```go
intSlice := []int{-3, -2, -1, 0, 1, 2, 3}
float64Slice := svgPlot.ConvertSliceToFloat64(intSlice)
```
### DrawSmooth
**DrawSmooth** makes a plot where points are joined with a smooth curve
The greed of the plot starts from the smallest values of xArray and yArray.
```go
package main

import (
	"github.com/nikk-gr/svgPlot"
	"os"
)

func main() {
	TotalHeight := uint(200) // px
	TotalWidth := uint(400)  // px
	xDivisionsQty := uint(10)
	yDivisionsQty := uint(10)
	NameOfX := string("x")
	NameOfY := string("y")
	xArray := []float64{-3, -2, -1, 0, 1, 2, 3}
	yArray := []float64{9, 4, 1, 0, 1, 4, 9}
	plot, _ := svgPlot.DrawSmooth(TotalHeight, TotalWidth, xDivisionsQty, yDivisionsQty, xArray, yArray, NameOfX, NameOfY)
	_ = os.WriteFile("smooth.svg", []byte(plot), 777)
}
```
![DrawSmooth](./.github/img/smooth.svg "DrawSmooth")
### DrawSmoothFromZero
**DrawSmoothFromZero** makes a plot where points are joined with a smooth curve
The greed of the plot starts from the zero point.
Only positive values of X and Y arrays are allowed
```go
package main

import (
	"github.com/nikk-gr/svgPlot"
	"os"
)

func main() {
	TotalHeight := uint(200) // px
	TotalWidth := uint(400)  // px
	xDivisionsQty := uint(10)
	yDivisionsQty := uint(10)
	NameOfX := string("x")
	NameOfY := string("y")
	xArray := []float64{1, 2, 3, 4, 5, 6, 7}
	yArray := []float64{9, 4, 1, 0, 1, 4, 9}
	plot, _ := svgPlot.DrawSmoothFromZero(TotalHeight, TotalWidth, xDivisionsQty, yDivisionsQty, xArray, yArray, NameOfX, NameOfY)
	_ = os.WriteFile("smooth_from_zero.svg", []byte(plot), 777)
}
```
![DrawSmoothFromZero](./.github/img/smooth_from_zero.svg "DrawSmoothFromZero")
### DrawPoint
**DrawPoint** makes a plot where points are not joined and indicated by the small circles
The greed of the plot starts from the smallest values of xArray and yArray.
```go
package main

import (
	"github.com/nikk-gr/svgPlot"
	"os"
)

func main() {
	TotalHeight := uint(200) // px
	TotalWidth := uint(400)  // px
	xDivisionsQty := uint(10)
	yDivisionsQty := uint(10)
	NameOfX := string("x")
	NameOfY := string("y")
	xArray := []float64{-3, -2, -1, 0, 1, 2, 3}
	yArray := []float64{9, 4, 1, 0, 1, 4, 9}
	plot, _ := svgPlot.DrawPoint(TotalHeight, TotalWidth, xDivisionsQty, yDivisionsQty, xArray, yArray, NameOfX, NameOfY)
	_ = os.WriteFile("point.svg", []byte(plot), 777)
}
```
![DrawPoint](./.github/img/point.svg "DrawPoint")
### DrawPointFromZero
**DrawPointFromZero** makes a plot where points are not joined and indicated by the small circles
The greed of the plot starts from the zero point.
Only positive values of X and Y arrays are allowed
```go
package main

import (
	"github.com/nikk-gr/svgPlot"
	"os"
)

func main() {
	TotalHeight := uint(200) // px
	TotalWidth := uint(400)  // px
	xDivisionsQty := uint(10)
	yDivisionsQty := uint(10)
	NameOfX := string("x")
	NameOfY := string("y")
	xArray := []float64{1, 2, 3, 4, 5, 6, 7}
	yArray := []float64{9, 4, 1, 0, 1, 4, 9}
	plot, _ := svgPlot.DrawPointFromZero(TotalHeight, TotalWidth, xDivisionsQty, yDivisionsQty, xArray, yArray, NameOfX, NameOfY)
	_ = os.WriteFile("point_from_zero.svg", []byte(plot), 777)
}
```
![DrawPointFromZero](./.github/img/point_from_zero.svg "DrawPointFromZero")
### DrawAngular
**DrawAngular** makes a plot where points are joined with strait lines
The greed of the plot starts from the smallest values of xArray and yArray.
```go
package main

import (
	"github.com/nikk-gr/svgPlot"
	"os"
)

func main() {
	TotalHeight := uint(200) // px
	TotalWidth := uint(400)  // px
	xDivisionsQty := uint(10)
	yDivisionsQty := uint(10)
	NameOfX := string("x")
	NameOfY := string("y")
	xArray := []float64{-3, -2, -1, 0, 1, 2, 3}
	yArray := []float64{9, 4, 1, 0, 1, 4, 9}
	plot, _ := svgPlot.DrawAngular(TotalHeight, TotalWidth, xDivisionsQty, yDivisionsQty, xArray, yArray, NameOfX, NameOfY)
	_ = os.WriteFile("angular.svg", []byte(plot), 777)
}
```
![DrawAngular](./.github/img/angular.svg "DrawAngular")
### DrawAngularFromZero
**DrawAngularFromZero** makes a plot where points are joined with strait lines
The greed of the plot starts from the zero point.
Only positive values of X and Y arrays are allowed
```go
package main

import (
	"github.com/nikk-gr/svgPlot"
	"os"
)

func main() {
	TotalHeight := uint(200) // px
	TotalWidth := uint(400)  // px
	xDivisionsQty := uint(10)
	yDivisionsQty := uint(10)
	NameOfX := string("x")
	NameOfY := string("y")
	xArray := []float64{1, 2, 3, 4, 5, 6, 7}
	yArray := []float64{9, 4, 1, 0, 1, 4, 9}
	plot, _ := svgPlot.DrawAngularFromZero(TotalHeight, TotalWidth, xDivisionsQty, yDivisionsQty, xArray, yArray, NameOfX, NameOfY)
	_ = os.WriteFile("angular_from_zero.svg", []byte(plot), 777)
}
```
![DrawAngularFromZero](./.github/img/angular_from_zero.svg "DrawAngularFromZero")
### Combined Plot
The **Combined plot** is a plot with a combination of angular lines, smooth curves and point plots

```go
package main

import (
	"github.com/nikk-gr/svgPlot"
	"os"
)

func main() {
	// Creating the new plot 200x400 px 10 divisions on each side
	plot, _ := svgPlot.NewPlot(200, 400, 10, 10, "x", "y")
	// Add smooth parabola plot
	_ = plot.AddSmooth([]float64{1, 2, 3, 4, 5, 6, 7}, []float64{10, 5, 2, 1, 2, 5, 10})
	// Add a point to the bottom of the parabola
	_ = plot.AddPoint([]float64{4}, []float64{1})
	// Add some line
	_ = plot.AddAngular([]float64{1, 1, 7, 7}, []float64{10, 1, 1, 10})
	// generate svg
	svg, _ := plot.Draw()
	svgFromZero, _ := plot.DrawFromZero()
	// write it to the disk
	_ = os.WriteFile("combined.svg", []byte(svg), 777)
	_ = os.WriteFile("combined_from_zero.svg", []byte(svgFromZero), 777)
}
```
![Draw](./.github/img/combined.svg "Draw")
![DrawFromZero](./.github/img/combined_from_zero.svg "DrawFromZero")
