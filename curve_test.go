// Package svgPlot Copyright 2023 Gryaznov Nikita
// Licensed under the Apache License, Version 2.0
package svgPlot

import (
	"errors"
	"fmt"
	"testing"
)

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

		if fmt.Sprint(err) != fmt.Sprint(t1.err) {
			t.Errorf("%sTest %d failed%s\tgot: %s, want: %s\n", red, testId, normal, t1.err, err)
		} else if curve != t1.curve {
			t.Errorf("%sTest %d failed%s\tgot: %s, want: %s\n", red, testId, normal, t1.curve, curve)
		} else {
			t.Logf("%sTest %d  success%s\t%s\n", green, testId, normal, "")
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
		{
			x0:    194.36363636363637,
			y0:    187,
			gradX: 0.026525198938992044,
			gradY: 0.053475935828877004,
			x:     []float64{-3, -2, 0, 1, 2, 3},
			y:     []float64{9, 4, 1, 0, 1, 4, 9},
			err:   errors.New("len(x)!=len(y)"),
		},
		{
			x0:    194.36363636363637,
			y0:    187,
			gradX: 0,
			gradY: 0.053475935828877004,
			x:     []float64{-3, -2, -1, 0, 1, 2, 3},
			y:     []float64{9, 4, 1, 0, 1, 4, 9},
			err:   errors.New("gradX<=0"),
		},
		{
			x0:    194.36363636363637,
			y0:    187,
			gradX: 0.026525198938992044,
			gradY: 0.0,
			x:     []float64{-3, -2, -1, 0, 1, 2, 3},
			y:     []float64{9, 4, 1, 0, 1, 4, 9},
			err:   errors.New("gradY<=0"),
		},
	}
	for _, testCase := range testArr {
		test(testCase)
	}
}
func TestMakeSmoothCurve(t *testing.T) {
	type testType struct {
		x0, y0, gradX, gradY float64
		x, y                 []float64
		curve                string
		err                  error
	}
	var testId uint64
	test := func(t1 testType) {
		t.Logf("Test %d start", testId)
		curve, err := makeSmoothCurve(t1.x0, t1.y0, t1.gradX, t1.gradY, t1.x, t1.y)

		if fmt.Sprint(err) != fmt.Sprint(t1.err) {
			t.Errorf("%sTest %d failed%s\tgot: %s, want: %s\n", red, testId, normal, t1.err, err)
		} else if curve != t1.curve {
			t.Errorf("%sTest %d failed%s\twant: %s, got: %s\n", red, testId, normal, t1.curve, curve)
			t.Errorf("%s\n", curve)
		} else {
			t.Logf("%sTest %d success%s\t%s\n", green, testId, normal, "")
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
			curve: "<path d=\"M 81.26 18.70 C 81.26,18.70 109.920,94.26 118.96,112.20 C 128.02,130.17 142.549,154.30 156.66,168.30 C 165.17,176.74 183.054,187.00 194.36,187.00 C 205.67,187.00 223.559,176.74 232.06,168.30 C 246.18,154.30 260.708,130.17 269.76,112.20 C 269.76,112.20 291.373,63.09 307.46,18.70 \" fill=\"none\" stroke=\"black\" stroke-width=\"3\" stroke-linecap=\"round\"/>",
		},
		{
			x0:    194.36363636363637,
			y0:    187,
			gradX: 0.026525198938992044,
			gradY: 0.053475935828877004,
			x:     []float64{1, 3},
			y:     []float64{1, 9},
			curve: "<path d=\"M 232.06363636363636 168.3 L 307.4636363636364 18.69999999999999\" fill=\"none\" stroke=\"black\" stroke-width=\"3\"/>",
		},
		{
			x0:    194.36363636363637,
			y0:    187,
			gradX: 0.026525198938992044,
			gradY: 0.053475935828877004,
			x:     []float64{1, 1, 3},
			y:     []float64{1, 1, 9},
			curve: "<path d=\"M 232.06363636363636 168.3 L 232.06363636363636 168.3 L 307.4636363636364 18.69999999999999\" fill=\"none\" stroke=\"black\" stroke-width=\"3\"/>",
		},
		{
			x0:    194.36363636363637,
			y0:    187,
			gradX: 0.026525198938992044,
			gradY: 0.053475935828877004,
			x:     []float64{-3, -2, 0, 1, 2, 3},
			y:     []float64{9, 4, 1, 0, 1, 4, 9},
			err:   errors.New("len(x)!=len(y)"),
		},
		{
			x0:    194.36363636363637,
			y0:    187,
			gradX: 0,
			gradY: 0.053475935828877004,
			x:     []float64{-3, -2, -1, 0, 1, 2, 3},
			y:     []float64{9, 4, 1, 0, 1, 4, 9},
			err:   errors.New("gradX<=0"),
		},
		{
			x0:    194.36363636363637,
			y0:    187,
			gradX: 0.026525198938992044,
			gradY: 0.0,
			x:     []float64{-3, -2, -1, 0, 1, 2, 3},
			y:     []float64{9, 4, 1, 0, 1, 4, 9},
			err:   errors.New("gradY<=0"),
		},
	}
	for _, testCase := range testArr {
		test(testCase)
	}
}
func TestMakePointCurve(t *testing.T) {
	type testType struct {
		x0, y0, gradX, gradY float64
		x, y                 []float64
		curve                string
		err                  error
	}
	var testId uint64
	test := func(t1 testType) {
		t.Logf("Test %d start", testId)
		curve, err := makePointCurve(t1.x0, t1.y0, t1.gradX, t1.gradY, t1.x, t1.y)

		if fmt.Sprint(err) != fmt.Sprint(t1.err) {
			t.Errorf("%sTest %d failed%s\tgot: %s, want: %s\n", red, testId, normal, t1.err, err)
		} else if curve != t1.curve {
			t.Errorf("%sTest %d failed%s\twant: %s, got: %s\n", red, testId, normal, t1.curve, curve)
			t.Errorf("%s\n", curve)
		} else {
			t.Logf("%sTest %d success%s\t%s\n", green, testId, normal, "")
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
			curve: `<circle cx="81" cy="19" r="5"/>
<circle cx="1.2E+02" cy="1.1E+02" r="5"/>
<circle cx="1.6E+02" cy="1.7E+02" r="5"/>
<circle cx="1.9E+02" cy="1.9E+02" r="5"/>
<circle cx="2.3E+02" cy="1.7E+02" r="5"/>
<circle cx="2.7E+02" cy="1.1E+02" r="5"/>
<circle cx="3.1E+02" cy="19" r="5"/>
`,
		},
		{
			x0:    194.36363636363637,
			y0:    187,
			gradX: 0.026525198938992044,
			gradY: 0.053475935828877004,
			x:     []float64{-3, -2, 0, 1, 2, 3},
			y:     []float64{9, 4, 1, 0, 1, 4, 9},
			err:   errors.New("len(x)!=len(y)"),
		},
		{
			x0:    194.36363636363637,
			y0:    187,
			gradX: 0,
			gradY: 0.053475935828877004,
			x:     []float64{-3, -2, -1, 0, 1, 2, 3},
			y:     []float64{9, 4, 1, 0, 1, 4, 9},
			err:   errors.New("gradX<=0"),
		},
		{
			x0:    194.36363636363637,
			y0:    187,
			gradX: 0.026525198938992044,
			gradY: 0.0,
			x:     []float64{-3, -2, -1, 0, 1, 2, 3},
			y:     []float64{9, 4, 1, 0, 1, 4, 9},
			err:   errors.New("gradY<=0"),
		},
	}
	for _, testCase := range testArr {
		test(testCase)
	}
}
func TestGetAsymmetricPointCoefficients(t *testing.T) {
	type testType struct {
		p0, p1, p2 point
		r0, r1     point
		err        error
	}
	var testId uint64
	test := func(t1 testType) {
		t.Logf("Test %d start", testId)
		r0, r1, err := getAsymmetricPointCoefficients(t1.p0, t1.p1, t1.p2)

		if fmt.Sprint(err) != fmt.Sprint(t1.err) {
			t.Errorf("%sTest %d failed%s\tgot: %s, want: %s\n", red, testId, normal, t1.err, err)
		} else if r0 != t1.r0 {
			t.Errorf("%sTest %d failed%s\twant: %f, got: %f\n", red, testId, normal, t1.r0, r0)
		} else if r1 != t1.r1 {
			t.Errorf("%sTest %d failed%s\twant: %f, got: %f\n", red, testId, normal, t1.r1, r1)
		} else {
			t.Logf("%sTest %d success%s\t%s\n", green, testId, normal, "")
		}
		testId++
	}
	testArr := []testType{
		{
			p0: point{0, 0},
			p1: point{50, 50},
			p2: point{100, 0},
			r0: point{65, 50},
			r1: point{35, 50},
		},
		{
			p0: point{0, 0},
			p1: point{10, 50},
			p2: point{100, 0},
			r0: point{37, 50},
			r1: point{7, 50},
		},
		{
			p0:  point{0, 0},
			p1:  point{0, 0},
			p2:  point{100, 0},
			err: errors.New("points 1 and 2 have the same coordinates"),
		},
		{
			p0:  point{0, 0},
			p1:  point{100, 0},
			p2:  point{100, 0},
			err: errors.New("points 2 and 3 have the same coordinates"),
		},
		{
			p0:  point{100, 0},
			p1:  point{0, 0},
			p2:  point{100, 0},
			err: errors.New("points 1 and 3 have the same coordinates"),
		},
	}
	for _, testCase := range testArr {
		test(testCase)
	}
}
