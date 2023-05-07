// Package svgPlot Copyright 2023 Gryaznov Nikita
// Licensed under the Apache License, Version 2.0
package svgPlot

import "math"

// linearInterpolation returns result of linear interpolation
// f( x ) = a * x + b, f( x1 ) = val1, f( x2 ) = val2, f( target_x ) = result
func linearInterpolation(targetX, x1, x2, val1, val2 float64) float64 {
	if x1 == x2 {
		return (val1 + val2) / 2
	} else {

		return -((x2*val1 - x1*val2) / (x1 - x2)) - ((-val1+val2)*targetX)/(x1-x2)
	}
}

func getDistance(p1, p2 point) float64 {
	return math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2))
}
