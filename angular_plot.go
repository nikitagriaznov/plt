// Package svgPlot Copyright 2023 Gryaznov Nikita
// Licensed under the Apache License, Version 2.0
package svgPlot

import (
	"errors"
	"fmt"
)

type (
	numeric interface {
		int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float64 | float32
	}
	Style struct {
		TotalHeight, TotalWidth      uint   // Size of resulting picture
		XDivisionsQty, YDivisionsQty uint   // Required number of divisions on X and Y scale
		NameOfX, NameOfY             string // Names of X and Y axis. Hidden if empty
	}
)

// DrawAngular make the most compact plot with strait lines between points
// TotalHeight and TotalWidth defines the size of resulting picture
// X and Y is a parallel arrays of point coordinates
// NameOfX and NameOfY is axis labels. Max allowed length is 6 chars
// if NameOfX or NameOfX != "" it will replace the last number on the axis
func DrawAngular(TotalHeight, TotalWidth, xDivisionsQty, yDivisionsQty uint, X, Y []float64, NameOfX, NameOfY string) (plot string, err error) {
	var (
		xMin, xMax, yMin, yMax float64
	)
	// convert X, Y slices to float64

	// get min & max values
	yMin, yMax, err = getMinMax(Y)
	if err != nil {
		return
	}
	xMin, xMax, err = getMinMax(X)
	if err != nil {
		return
	}
	plot, err = drawAng(TotalHeight, TotalWidth, xDivisionsQty, yDivisionsQty, X, Y, xMin, xMax, yMin, yMax, NameOfX, NameOfY)
	return
}

// DrawAngularFromZero make plot with strait lines between points
// Coordinate plane starts from (0, 0) point
// TotalHeight and TotalWidth defines the size of resulting picture
// X and Y is a parallel arrays of point coordinates, only positive x and y is allowed
// NameOfX and NameOfY is axis labels. Max allowed length is 6 chars
// if NameOfX or NameOfX != "" it will replace the last number on the axis
func DrawAngularFromZero(TotalHeight, TotalWidth, xDivisionsQty, yDivisionsQty uint, x, y []float64, NameOfX, NameOfY string) (plot string, err error) {
	var (
		xMax, yMax float64
	)
	// convert x, y slices to float64
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
	plot, err = drawAng(TotalHeight, TotalWidth, xDivisionsQty, yDivisionsQty, x, y, 0, xMax, 0, yMax, NameOfX, NameOfY)
	return
}

// checkPositive returns nil if every number in arr is >=0, error if one or more is less, then 0
func checkPositive(arr []float64) (err error) {
	for key, val := range arr {
		if val < 0 {
			return fmt.Errorf("%d-th element is les, then 0", key)
		}
	}
	return nil
}

// ConvertSliceToFloat64 convert slices of any numeric type to []float64
func ConvertSliceToFloat64[T numeric](in []T) (out []float64) {
	out = make([]float64, len(in))
	for i := 0; i < len(in); i++ {
		out[i] = float64(in[i])
	}
	return
}

// draw returns a complete plot picture with points joined by straight lines
func drawAng(height, width, xDivisionsQty, yDivisionsQty uint, x, y []float64, xMin, xMax, yMin, yMax float64, xName, yName string) (result string, err error) {
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
	plot, err = makeAngularCurve(x0, y0, gradX, gradY, x, y)
	if err != nil {
		return
	}
	result = "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n"
	result += fmt.Sprintf("<svg width=\"%d\" height=\"%d\" viewBox=\"0 0 %d %d\" version=\"1.1\" xmlns=\"http://www.w3.org/2000/svg\">\n", width, height, width, height)
	result += greed + "\n"
	result += plot + "\n"
	result += "</svg>"
	return
}
