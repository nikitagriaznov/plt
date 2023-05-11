// Package svgPlot Copyright 2023 Gryaznov Nikita
// Licensed under the Apache License, Version 2.0
package svgPlot

import (
	"errors"
	"fmt"
)

// DrawPoint make the most compact plot with strait lines between points
// TotalHeight and TotalWidth defines the size of resulting picture
// X and Y is a parallel arrays of point coordinates
// NameOfX and NameOfY is axis labels. Max allowed length is 6 chars
// if NameOfX or NameOfX != "" it will replace the last number on the axis
func DrawPoint(TotalHeight, TotalWidth, xDivisionsQty, yDivisionsQty uint, X, Y []float64, NameOfX, NameOfY string) (plot string, err error) {
	var (
		xMin, xMax, yMin, yMax float64
	)

	// get min & max values
	yMin, yMax, err = getMinMax(Y)
	if err != nil {
		return
	}
	xMin, xMax, err = getMinMax(X)
	if err != nil {
		return
	}
	plot, err = drawPo(TotalHeight, TotalWidth, xDivisionsQty, yDivisionsQty, X, Y, xMin, xMax, yMin, yMax, NameOfX, NameOfY)
	return
}

// DrawPointFromZero make plot with strait lines between points
// Coordinate plane starts from (0, 0) point
// TotalHeight and TotalWidth defines the size of resulting picture
// X and Y is a parallel arrays of point coordinates, only positive x and y is allowed
// NameOfX and NameOfY is axis labels. Max allowed length is 6 chars
// if NameOfX or NameOfX != "" it will replace the last number on the axis
func DrawPointFromZero(TotalHeight, TotalWidth, xDivisionsQty, yDivisionsQty uint, x, y []float64, NameOfX, NameOfY string) (plot string, err error) {
	var (
		xMax, yMax float64
	)
	err = checkPositive(x)
	if err != nil {
		return
	}
	err = checkPositive(y)
	if err != nil {
		return
	}
	// get min & max values
	_, yMax, err = getMinMax(y)
	if err != nil {
		return
	}
	_, xMax, err = getMinMax(x)
	if err != nil {
		return
	}
	plot, err = drawPo(TotalHeight, TotalWidth, xDivisionsQty, yDivisionsQty, x, y, 0, xMax, 0, yMax, NameOfX, NameOfY)
	return
}

// drawPo returns a complete plot picture with points joined by straight lines
func drawPo(height, width, xDivisionsQty, yDivisionsQty uint, x, y []float64, xMin, xMax, yMin, yMax float64, xName, yName string) (result string, err error) {
	var (
		greed, plot                      string
		x0, y0, gradX, gradY, xLen, yLen float64
		xZeroPos, yZeroPos               int
		xNums, yNums                     []string
	)

	if len(xName) > 6 {
		err = errors.New("xName max len is 6")
		return
	}
	if len(yName) > 6 {
		err = errors.New("yName max len is 6")
		return
	}
	xNums, xLen, xZeroPos, err = makeArr(xMin, xMax, xDivisionsQty)
	if err != nil {
		return
	}
	yNums, yLen, yZeroPos, err = makeArr(yMin, yMax, yDivisionsQty)
	if err != nil {
		return
	}
	if xName != "" {
		xNums[len(xNums)-1] = xName
	}
	if yName != "" {
		yNums[len(yNums)-1] = yName
	}
	greed, x0, y0, gradX, gradY, err = makeGreed(height, width, xNums, yNums, xLen, yLen, xZeroPos, yZeroPos)
	if err != nil {
		return
	}
	plot, err = makePointCurve(x0, y0, gradX, gradY, x, y)
	if err != nil {
		return
	}
	result = "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n"
	result += fmt.Sprintf("<svg width=\"%d\" height=\"%d\" viewBox=\"0 0 %d %d\" version=\"1.1\" xmlns=\"http://www.w3.org/2000/svg\">\n\"", width, height, width, height)
	result += greed + "\n"
	result += plot + "\n"
	result += "</svg>"
	return
}
