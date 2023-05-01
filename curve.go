// Package svgPlot Copyright 2023 Gryaznov Nikita Licensed under the Apache
// License, Version 2.0 (the «License»);
package svgPlot

import (
	"errors"
	"fmt"
)

// makeCurve returns svg code for a plot line (strait lines between points)
func makeAngularCurve(x0, y0, gradX, gradY float64, x, y []float64) (curve string, err error) {
	// check input
	{
		if len(x) != len(y) {
			err = errors.New("len(x)!=len(y)")
			return
		}
		if x0 < 0 {
			err = errors.New("x0<0")
			return
		}
		if y0 < 0 {
			err = errors.New("y0<0")
			return
		}
		if gradX <= 0 {
			err = errors.New("gradX<=0")
			return
		}
		if gradY <= 0 {
			err = errors.New("gradY<=0")
			return
		}
	}
	// make curve
	curve += "<path d=\""
	var tmpX, tmpY string
	for i := 0; i < len(x); i++ {
		if i == 0 {
			curve += "M "
		} else {
			curve += " L "
		}
		tmpX = fmt.Sprint(x0 + x[i]/gradX)
		tmpY = fmt.Sprint(y0 - y[i]/gradY)
		curve += tmpX + " " + tmpY
	}
	curve += "\" fill=\"none\" stroke=\"black\" stroke-width=\"3\"/>"
	return
}
