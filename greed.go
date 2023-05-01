// Package svgPlot Copyright 2023 Gryaznov Nikita Licensed under the Apache
// License, Version 2.0 (the «License»);
package svgPlot

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

// getMinMax returns minimum and maximum values of array
func getMinMax(arr []float64) (min, max float64, err error) {
	if len(arr) == 0 {
		err = errors.New("empty input")
		return
	}
	min = arr[0]
	max = arr[0]
	if len(arr) > 1 {
		for i := 1; i < len(arr); i++ {
			if arr[i] < min {
				min = arr[i]
			} else if arr[i] > max {
				max = arr[i]
			}
		}
	}
	return
}

// getExp returns order and first digit of a number
func getExp(num float64) (order float64) {
	exp := math.Log10(math.Abs(num))
	exp2 := int(math.Floor(exp))
	return math.Pow10(exp2)
}

// makeArr returns the array of
func makeArr(min float64, max float64, divisionsQty uint) (result []string, l float64, zeroPosition int, err error) {
	stepArr := [8]float64{1, 2, 2.5, 5, 0.1, 0.2, 0.25, 0.5}
	if divisionsQty == 0 {
		divisionsQty = 10
	}
	// check input
	if min >= max {
		err = errors.New("min>=max")
		return
	}
	// count scale
	l = max - min
	order := getExp(l)
	// search for right step
	stp := stepArr[0] * order
	qty := l / stp
	for i := 1; i < len(stepArr); i++ {
		if math.Abs(qty-float64(divisionsQty)) > math.Abs(l/(stepArr[i]*order)-float64(divisionsQty)) {
			stp = stepArr[i] * order
			qty = l / stp
		}
	}

	// make & fill arr
	var tmp float64
	if max > 0 && min <= 0 {
		minQty := math.Ceil(math.Abs(min) / stp)
		result = make([]string, int(math.Ceil(math.Abs(max)/stp)+minQty)+1)
		tmp = minQty * stp * -1
	} else if max > 0 && min >= 0 {
		minQty := math.Floor(math.Abs(min) / stp)
		result = make([]string, int(math.Ceil(math.Abs(max)/stp)+minQty)+1)
		tmp = minQty * stp
	} else /*if max <= 0 && min < 0*/ {
		minQty := math.Ceil(math.Abs(min) / stp)
		result = make([]string, int(math.Floor(math.Abs(max)/stp)+minQty)+1)
		tmp = minQty * stp
	}
	min = tmp
	convert := func(num float64) string {
		if num == 0 {
			return "0"
		} else {
			return strconv.FormatFloat(tmp, 'g', 2, 64)
		}
	}
	result[0] = convert(tmp)
	j := 1
	for {
		tmp += stp
		result[j] = convert(tmp)
		if tmp >= max {
			max = tmp
			break
		}
		j++
	}
	l = max - min
	zeroPosition = -int(min / stp)
	return
}

// getWordLen returns width of number as a text with Arial font
func getWordLen(num string) (width int) {
	var tmp float64
	getRuneW := func(r uint8) (w float64) {
		switch r {
		case 'c', 's', 'v', 'x', 'y', 'z', 'k', 'J':
			w = 6
		case 'i', 'l', 'j':
			w = 2.67
		case '.', '/', ',', 'f', 't', 'I':
			w = 3.34
		case '-', 'r':
			w = 4
		case 'm', 'M':
			w = 10
		case 'w', 'C', 'D', 'H', 'N', 'R':
			w = 8.67
		case 'W':
			w = 11.33
		case 'G', 'O', 'Q':
			w = 9.34
		case 'A', 'B', 'E', 'P', 'S', 'V', 'X', 'Y', 'K':
			w = 8.01
		case 'Z', 'T', 'F':
			w = 7.34
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'd', 'e', 'g', 'h', 'n', 'o', 'p', 'q', 'L':
			w = 6.68
		default:
			w = 10
		}
		return
	}
	for i := 0; i < len(num); i++ {
		tmp += getRuneW(num[i])
	}
	return int(math.Ceil(tmp))
}

// makeGreed returns svg code for a greed
func makeGreed(height, width uint, xNumArray []string, yNumArray []string, xLen, yLen float64, xZeroPos, yZeroPos int) (greed string, x0, y0, gradX, gradY float64, err error) {
	// check input values

	if len(yNumArray) == 0 {
		err = errors.New("empty yNumArray")
		return
	}
	if len(xNumArray) == 0 {
		err = errors.New("empty xNumArray")
		return
	}

	// count paddings
	const (
		textHeight    = 9
		paddingTop    = 0 // height af arial 12 digits
		paddingBottom = textHeight + textHeight/2
		paddingRight  = 0
	)
	var (
		paddingLeft int
	)
	for i := 0; i < len(yNumArray); i++ {
		tmp := getWordLen(yNumArray[i])
		if tmp > paddingLeft {
			paddingLeft = tmp
		}
	}
	paddingLeft += textHeight

	verticalStep := float64(int(height)-paddingTop-paddingBottom) / float64(len(yNumArray)-1)
	horizontalStep := float64(int(width)-paddingRight-paddingLeft) / float64(len(xNumArray)-1)

	// check input values
	if height < 100+paddingTop+paddingBottom {
		err = errors.New("height is too small")
		return
	}
	if int(width) < 150+paddingLeft+paddingRight {
		err = errors.New("width is too small")
		return
	}

	// Add text styles
	greed += "<style>\n.axis {\nfont-family=\"Arial, Helvetica, sans-serif\">;\nfont-size: 12pt;\n}\n</style>\n"
	// make horizontal greed
	{
		x1 := paddingLeft
		x2 := width - paddingRight
		xt := paddingLeft - textHeight/2
		y := float64(paddingTop)
		for i := 0; i < len(yNumArray); i++ {
			y1 := fmt.Sprint(math.Round(y*100) / 100)
			greed += fmt.Sprintf("<line x1=\"%d\" y1=\"%s\" x2=\"%d\" y2=\"%s\" stroke=\"black\" />\n", x1, y1, x2, y1)
			if i == 0 {
				greed += fmt.Sprintf("<text text-anchor=\"end\" class=\"axis\" x=\"%d\" y=\"%s\">%s</text>\n", xt, fmt.Sprint(math.Round((y+textHeight)*100)/100), yNumArray[len(yNumArray)-1-i])
			} else {
				greed += fmt.Sprintf("<text text-anchor=\"end\" class=\"axis\" x=\"%d\" y=\"%s\">%s</text>\n", xt, fmt.Sprint(math.Round((y+textHeight/2)*100)/100), yNumArray[len(yNumArray)-1-i])
			}
			y += verticalStep
		}
	}
	// make vertical greed
	{
		y1 := paddingTop
		y2 := height - paddingBottom
		yt := height - paddingBottom + textHeight + textHeight/2
		x := float64(paddingLeft)
		for i := 0; i < len(xNumArray); i++ {
			x1 := fmt.Sprint(math.Round(x*100) / 100)
			greed += fmt.Sprintf("<line x1=\"%s\" y1=\"%d\" x2=\"%s\" y2=\"%d\" stroke=\"black\" />\n", x1, y1, x1, y2)
			if i != len(xNumArray)-1 {
				greed += fmt.Sprintf("<text text-anchor=\"middle\" class=\"axis\" x=\"%s\" y=\"%d\">%s</text>\n", x1, yt, fmt.Sprint(xNumArray[i]))
			} else {
				greed += fmt.Sprintf("<text text-anchor=\"end\" class=\"axis\" x=\"%s\" y=\"%d\">%s</text>\n", x1, yt, fmt.Sprint(xNumArray[i]))
			}
			x += horizontalStep
		}
	}
	x0 = float64(paddingLeft) + float64(xZeroPos)*horizontalStep
	y0 = float64(height-paddingBottom) + float64(yZeroPos)*verticalStep
	gradX = xLen / float64(int(width)-paddingRight-paddingLeft)
	gradY = yLen / float64(int(height)-paddingTop-paddingBottom)
	return
}
