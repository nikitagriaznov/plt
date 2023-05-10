// Package svgPlot Copyright 2023 Gryaznov Nikita
// Licensed under the Apache License, Version 2.0
package svgPlot

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

type (
	point struct {
		x float64
		y float64
	}
	translation struct {
		dx float64
		dy float64
	}
	bezier2 [3]point
	bezier3 [4]point
)

// makeCurve returns svg code for a plot line (strait lines between points)
func makeAngularCurve(x0, y0, gradX, gradY float64, x, y []float64) (curve string, err error) {
	// check input
	{
		if len(x) != len(y) {
			err = errors.New("len(x)!=len(y)")
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

// makeCurve returns svg code for a plot line (strait lines between points)
func makePointCurve(x0, y0, gradX, gradY float64, x, y []float64) (curve string, err error) {
	// check input
	{
		if len(x) != len(y) {
			err = errors.New("len(x)!=len(y)")
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
	var tmpX, tmpY string
	for i := 0; i < len(x); i++ {
		tmpX = strconv.FormatFloat(x0+x[i]/gradX, 'f', 2, 64)
		tmpY = strconv.FormatFloat(y0-y[i]/gradY, 'f', 2, 64)
		curve += "<circle cx=\"" + tmpX + "\" cy=\"" + tmpY + "\" r=\"5\"/>\n"
	}

	return
}

// makeCurve returns svg code for a plot line (strait lines between points)
func makeSmoothCurve(x0, y0, gradX, gradY float64, x, y []float64) (curve string, err error) {
	// check input
	{
		if len(x) != len(y) {
			err = errors.New("len(x)!=len(y)")
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

	// Remove repetitions & go to image coordinates
	points := make([]point, len(x))
	{
		var (
			j   int
			tmp point
		)
		for i := 0; i < len(x); i++ {
			tmp = point{x0 + x[i]/gradX, y0 - y[i]/gradY}
			if i == 0 {
				points[j] = tmp
				j++
			} else if tmp != points[j-1] {
				points[j] = tmp
				j++
			}
		}
		if j != len(x) {
			points = points[0:j]
		}
	}

	// Check if it more than 3 points
	if len(points) < 3 {
		return makeAngularCurve(x0, y0, gradX, gradY, x, y)
	}

	// get bezier ref points
	leftPoint := make([]point, len(points)-2)
	rightPoint := make([]point, len(points)-2)
	for i := 1; i < len(points)-1; i++ {
		rightPoint[i-1], leftPoint[i-1], err = getAsymmetricPointCoefficients(points[i-1], points[i], points[i+1])
	}

	// make the extreme right and left curves and elevate it's degrees
	curves := make([]bezier3, len(points)-1)
	curves[0] = bezier2{
		points[0],
		leftPoint[0],
		points[1],
	}.elevate()
	curves[len(curves)-1] = bezier2{
		points[len(points)-2],
		rightPoint[len(rightPoint)-1],
		points[len(points)-1],
	}.elevate()

	// Check if points arrays are longer then 3. Make other segments
	if len(points) > 3 {
		for i := 1; i < len(curves)-1; i++ {
			curves[i] = bezier3{
				points[i],
				rightPoint[i-1],
				leftPoint[i],
				points[i+1],
			}
		}
	}

	// make curve
	curve += fmt.Sprintf("<path d=\"M %.2f %.2f ", points[0].x, points[0].y)
	for i := 0; i < len(curves); i++ {
		curve += curves[i].Print()
	}
	curve += "\" fill=\"none\" stroke=\"black\" stroke-width=\"3\" stroke-linecap=\"round\"/>"
	return
}

// getAsymmetricPointCoefficients returns additional bezier points to smooth p2 point
func getAsymmetricPointCoefficients(p0, p1, p2 point) (ref0, ref1 point, err error) {
	// check if the points have the same coordinates

	if p0 == p1 {
		err = errors.New("points 1 and 2 have the same coordinates")
		return
	}
	if p0 == p2 {
		err = errors.New("points 1 and 3 have the same coordinates")
		return
	}
	if p1 == p2 {
		err = errors.New("points 2 and 3 have the same coordinates")
		return
	}

	// 3 points describe the triangle
	// Find sides length p0 to p1 = l1, p0 to p2 = l0, p1 to p2 = l0
	l0 := getDistance(p1, p2)
	l1 := getDistance(p0, p2)
	l2 := getDistance(p0, p1)

	// Find half perimeter - p

	p := (l0 + l1 + l2) / 2

	//  Find l1 to p1 height - h1

	h1 := (2 * math.Sqrt(p*(p-l0)*(p-l1)*(p-l2))) / l1

	// Height h1 divides side l1 in to parts: l10 (on the p0 side) and l12 (on the p2 side).
	// (h1, l10, l0) and (h1, l12, l2) is two right angle triangles. Find it lengths of l10 & l12

	l10 := math.Sqrt(l0*l0 - h1*h1)
	l12 := math.Sqrt(l2*l2 - h1*h1)

	// Point coefficients is the coordinates of the points that is in the half of l10 & l12 lengths from p2
	// Find required displacement

	d10 := translation{
		dx: (p2.x - p0.x) * (l10 / l1) * 0.3,
		dy: (p2.y - p0.y) * (l10 / l1) * 0.3,
	}
	d12 := translation{
		dx: (p0.x - p2.x) * (l12 / l1) * 0.3,
		dy: (p0.y - p2.y) * (l12 / l1) * 0.3,
	}

	// Find coefficients

	ref0 = p1.translate(d10)
	ref1 = p1.translate(d12)
	return
}

// elevate increases the degree of the Bézier curve by 1
func (in bezier2) elevate() (out bezier3) {
	return bezier3{
		in[0],
		point{
			x: in[0].x + ((in[1].x - in[0].x) * (2 / 3)),
			y: in[0].y + ((in[1].y - in[0].y) * (2 / 3)),
		},
		point{
			x: in[1].x + ((in[2].x - in[1].x) / 3),
			y: in[1].y + ((in[2].y - in[1].y) / 3),
		},
		in[2],
	}
}

func (in point) translate(t translation) (out point) {
	return point{
		x: in.x + t.dx,
		y: in.y + t.dy,
	}
}

func (in bezier3) Print() string {
	return fmt.Sprintf("C %.2f,%.2f %.3f,%.2f %.2f,%.2f ", in[1].x, in[1].y, in[2].x, in[2].y, in[3].x, in[3].y)
}
