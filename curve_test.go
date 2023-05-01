// Copyright 2023 Gryaznov Nikita Licensed under the Apache License, Version 2.0 (the «License»);
package svgPlot

import "testing"

func TestMakeAngularCurve(t *testing.T) {
	type testType struct {
		x0, y0, gradX, gradY float64
		x, y                 []float64
		curve                string
		err                  error
	}
	var testId uint64
	test := func(t1 testType) {
		t.Logf("Test %d start", testId)
		curve, err := makeAngularCurve(t1.x0, t1.y0, t1.gradX, t1.gradY, t1.x, t1.y)

		if err != t1.err {
			t.Errorf("%sTest %d failed%s\tgot: %s, want: %s\n", red, testId, normal, t1.err, err)
		} else if curve != t1.curve {
			t.Errorf("%sTest %d failed%s\tgot: %s, want: %s\n", red, testId, normal, t1.curve, curve)
		} else {
			t.Logf("%sTest %d seccess%s\t%s\n", green, testId, normal, "")
		}
		testId++
	}
	testArr := []testType{
		{
			x0:    194.36363636363637,
			y0:    187,
			gradX: 0.026525198938992044,
			gradY: 0.053475935828877004,
			x:     []float64{-3, -2, -1, 0, 1, 2, 3},
			y:     []float64{9, 4, 1, 0, 1, 4, 9},
			curve: "<path d=\"M 81.26363636363638 18.69999999999999 L 118.96363636363638 112.2 L 156.66363636363639 168.3 L 194.36363636363637 187 L 232.06363636363636 168.3 L 269.76363636363635 112.2 L 307.4636363636364 18.69999999999999\" fill=\"none\" stroke=\"black\" stroke-width=\"3\"/>",
		},
	}
	for _, testCase := range testArr {
		test(testCase)
	}
}
