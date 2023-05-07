// Package svgPlot Copyright 2023 Gryaznov Nikita
// Licensed under the Apache License, Version 2.0
package svgPlot

import "testing"

func TestLinearInterpolation(t *testing.T) {
	var testId uint
	test := func(x1, y1, x2, y2, x3, y3 float64) {
		t.Logf("Test %d start", testId)
		res := linearInterpolation(x3, x1, x2, y1, y2)

		if res != y3 {
			t.Errorf("%sTest %d failed%s\tgot: %G, want: %G\n", red, testId, normal, res, y3)
		} else {
			t.Logf("%sTest %d  success%s\t%s\n", green, testId, normal, "")
		}

		testId++
	}
	test(0, 0, 10, 10, 5, 5)
	test(-5, -5, 0, 10, 5, 25)
	test(-5, -5, -1, 10, 5, 32.5)
	test(-5, -5, 123, 10, 5, -3.828125)
	test(8, 10, 8, 10, 8, 10)
}
