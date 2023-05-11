// Package svgPlot Copyright 2023 Gryaznov Nikita
// Licensed under the Apache License, Version 2.0
package svgPlot

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

func TestConvertSliceToFloat64(t *testing.T) {
	equal := func(a, b []float64) bool {
		if len(a) != len(b) {
			return false
		}
		for i, v := range a {
			if v != b[i] {
				return false
			}
		}
		return true
	}
	var testId uint64
	{
		testId = 0
		t.Logf("Test %d\tstart", testId)
		check1 := []float64{0, 1, 2, 3, 4}
		res := ConvertSliceToFloat64([]int{0, 1, 2, 3, 4})
		ok := equal(check1, res)
		if !ok {
			t.Errorf("%sTest %d failed%s\tgot: %g, want: %g\n", red, testId, normal, check1, res)
		} else {
			t.Logf("%sTest %d  success%s\t%s\n", green, testId, normal, "")
		}
	}
	{
		testId = 1
		t.Logf("Test %d\tstart", testId)
		check1 := []float64{-2, -1, 0, 1, 2, 0.5, 4}
		res := ConvertSliceToFloat64([]float64{-2, -1, 0, 1, 2, 0.5, 4})
		ok := equal(check1, res)
		if !ok {
			t.Errorf("%sTest %d failed%s\tgot: %g, want: %g\n", red, testId, normal, check1, res)
		} else {
			t.Logf("%sTest %d  success%s\t%s\n", green, testId, normal, "")
		}
	}
	{
		testId = 1
		t.Logf("Test %d\tstart", testId)
		check1 := []float64{0, 1, 2, 3, 4}
		res := ConvertSliceToFloat64([]uint{0, 1, 2, 3, 4})
		ok := equal(check1, res)
		if !ok {
			t.Errorf("%sTest %d failed%s\tgot: %g, want: %g\n", red, testId, normal, check1, res)
		} else {
			t.Logf("%sTest %d  success%s\t%s\n", green, testId, normal, "")
		}
	}
}
func TestCheckPositive(t *testing.T) {
	var testId uint
	test := func(in []float64, err1 error) {
		t.Logf("Test %d start", testId)
		err2 := checkPositive(in)

		if fmt.Sprint(err1) == fmt.Sprint(err2) {
			t.Logf("%sTest %d  success%s\t%s\n", green, testId, normal, "")
		} else {
			t.Errorf("%sTest %d failed%s\tgot: %s, want: %s\n", red, testId, normal, err2, err1)
		}

		testId++
	}
	test([]float64{1, 3, 5, 2, 2, 66, 34, 0, -2.4}, fmt.Errorf("%d-th element is les, then 0", 8))
	test([]float64{}, nil)
	test([]float64{1, 3, 5, 2, -2, 66, 34, 0, -0.00004}, fmt.Errorf("%d-th element is les, then 0", 4))
	test([]float64{1, 3, 5, 2, 2, 66, 34, 0, 0.00004}, nil)

}
func TestDrawAng(t *testing.T) {
	type testType struct {
		height, width, xDivisionsQty, yDivisionsQty uint
		x, y                                        []float64
		xMin, xMax, yMin, yMax                      float64
		xName, yName                                string
		plot                                        string
		err                                         error
	}
	var testId uint64
	test := func(t1 testType) {
		t.Logf("Test %d start", testId)
		plot, err := drawAng(t1.height, t1.width, t1.xDivisionsQty, t1.yDivisionsQty, t1.x, t1.y, t1.xMin, t1.xMax, t1.yMin, t1.yMax, t1.xName, t1.yName)

		if fmt.Sprint(err) != fmt.Sprint(t1.err) {
			t.Errorf("%sTest %d failed%s\tgot: %s, want: %s\n", red, testId, normal, t1.err, err)
		} else if plot != t1.plot {
			t.Errorf("%sTest %d failed%s\tgot: %s, want: %s\n", red, testId, normal, t1.plot, plot)
			_ = os.WriteFile(fmt.Sprintf("TestDrawAng%d.svg", testId), []byte(plot), 777)
		} else {
			t.Logf("%sTest %d  success%s\t%s\n", green, testId, normal, "")
		}
		testId++
	}
	testArr := []testType{
		{
			height:        200,
			width:         400,
			xDivisionsQty: 10,
			yDivisionsQty: 10,
			x:             []float64{-3, -2, -1, 0, 1, 2, 3},
			y:             []float64{9, 4, 1, 0, 1, 4, 9},
			xMin:          -5,
			xMax:          5,
			yMin:          0,
			yMax:          10,
			xName:         "x",
			yName:         "y",
			plot:          TestDrawAng0,
		},
		{
			height:        200,
			width:         400,
			xDivisionsQty: 10,
			yDivisionsQty: 10,
			x:             []float64{-3, -2, -1, 0, 1, 2, 3},
			y:             []float64{9, 4, 1, 0, 1, 4, 9},
			xMin:          -5,
			xMax:          5,
			yMin:          0,
			yMax:          10,
			plot:          TestDrawAng1,
		},
		{
			height:        200,
			width:         400,
			xDivisionsQty: 10,
			yDivisionsQty: 10,
			x:             []float64{-3, -2, -1, 0, 1, 2, 3},
			y:             []float64{9, 4, 1, 0, 1, 4, 9},
			xMin:          5,
			xMax:          5,
			yMin:          0,
			yMax:          10,
			xName:         "x",
			yName:         "y",
			err:           errors.New("min>=max"),
		},
		{
			height:        200,
			width:         400,
			xDivisionsQty: 10,
			yDivisionsQty: 10,
			x:             []float64{-3, -2, -1, 0, 1, 2, 3},
			y:             []float64{9, 4, 1, 0, 1, 4, 9},
			xMin:          -5,
			xMax:          5,
			yMin:          10,
			yMax:          10,
			xName:         "x",
			yName:         "y",
			err:           errors.New("min>=max"),
		},
		{
			height:        200,
			width:         400,
			xDivisionsQty: 10,
			yDivisionsQty: 10,
			x:             []float64{-3, -2, -1, 0, 1, 2, 3},
			y:             []float64{9, 4, 1, 0, 1, 4, 9},
			xMin:          -5,
			xMax:          5,
			yMin:          0,
			yMax:          10,
			xName:         "lorem ipsum",
			yName:         "y",
			err:           errors.New("xName max len is 6"),
		},
		{
			height:        200,
			width:         400,
			xDivisionsQty: 10,
			yDivisionsQty: 10,
			x:             []float64{-3, -2, -1, 0, 1, 2, 3},
			y:             []float64{9, 4, 1, 0, 1, 4, 9},
			xMin:          -5,
			xMax:          5,
			yMin:          0,
			yMax:          10,
			xName:         "x",
			yName:         "lorem ipsum",
			err:           errors.New("yName max len is 6"),
		},
		{
			height:        0,
			width:         0,
			xDivisionsQty: 10,
			yDivisionsQty: 10,
			x:             []float64{-3, -2, -1, 0, 1, 2, 3},
			y:             []float64{9, 4, 1, 0, 1, 4, 9},
			xMin:          -5,
			xMax:          5,
			yMin:          0,
			yMax:          10,
			xName:         "x",
			yName:         "y",
			err:           errors.New("height is too small"),
		},
		{
			height:        400,
			width:         400,
			xDivisionsQty: 10,
			yDivisionsQty: 10,
			x:             []float64{-3, -2, 0, 1, 2, 3},
			y:             []float64{9, 4, 1, 0, 1, 4, 9},
			xMin:          -5,
			xMax:          5,
			yMin:          0,
			yMax:          10,
			xName:         "x",
			yName:         "y",
			err:           errors.New("len(x)!=len(y)"),
		},
	}
	for _, testCase := range testArr {
		test(testCase)
	}
}
func TestDrawAngular(t *testing.T) {
	type testCase struct {
		TotalHeight, TotalWidth      uint
		xDivisionsQty, yDivisionsQty uint
		X, Y                         []float64
		NameOfX, NameOfY             string
		plot                         string
		err                          error
	}
	doTest := func(id int, c testCase) {
		t.Logf("Test %d\tstart", id)
		plot, err := DrawAngular(c.TotalHeight, c.TotalWidth, c.xDivisionsQty, c.yDivisionsQty, c.X, c.Y, c.NameOfX, c.NameOfY)
		if fmt.Sprint(err) != fmt.Sprint(c.err) {
			t.Errorf("%sTest %d failed%s\twant: %s, got: %s\n", red, id, normal, c.err, err)
		} else if plot != c.plot {
			t.Errorf("%sTest %d failed%s\twant: %s, got: %s\n", red, id, normal, c.plot, plot)
			_ = os.WriteFile(fmt.Sprintf("TestDrawAngular%d.svg", id), []byte(plot), 777)
		} else {
			t.Logf("%sTest %d  success%s\t%s\n", green, id, normal, "")
		}

	}
	testArr := []testCase{
		{
			TotalHeight:   200,
			TotalWidth:    400,
			xDivisionsQty: 10,
			yDivisionsQty: 10,
			NameOfX:       "x",
			NameOfY:       "y",
			X:             []float64{-3, -2, -1, 0, 1, 2, 3},
			Y:             []float64{9, 4, 1, 0, 1, 4, 9},
			plot:          TestDrawAngular0,
		},
		{
			TotalHeight:   200,
			TotalWidth:    400,
			xDivisionsQty: 10,
			yDivisionsQty: 10,
			NameOfX:       "x",
			NameOfY:       "y",
			X:             []float64{-3, 0, 1, 2, 3},
			Y:             []float64{9, 4, 1, 0, 1, 4, 9},
			err:           errors.New("len(x)!=len(y)"),
		},
		{
			TotalHeight:   200,
			TotalWidth:    400,
			xDivisionsQty: 10,
			yDivisionsQty: 10,
			NameOfX:       "x",
			NameOfY:       "y",
			X:             []float64{},
			Y:             []float64{9, 4, 1, 0, 1, 4, 9},
			err:           errors.New("empty input"),
		},
		{
			TotalHeight:   200,
			TotalWidth:    400,
			xDivisionsQty: 10,
			yDivisionsQty: 10,
			NameOfX:       "x",
			NameOfY:       "y",
			X:             []float64{9, 4, 1, 0, 1, 4, 9},
			Y:             []float64{},
			err:           errors.New("empty input"),
		},
	}
	for key, val := range testArr {
		doTest(key, val)
	}
}
func TestDrawAngularFromZero(t *testing.T) {

	type testCase struct {
		TotalHeight, TotalWidth      uint
		xDivisionsQty, yDivisionsQty uint
		X, Y                         []float64
		NameOfX, NameOfY             string
		plot                         string
		err                          error
	}
	doTest := func(id int, c testCase) {
		t.Logf("Test %d\tstart", id)
		plot, err := DrawAngularFromZero(c.TotalHeight, c.TotalWidth, c.xDivisionsQty, c.yDivisionsQty, c.X, c.Y, c.NameOfX, c.NameOfY)
		if fmt.Sprint(err) != fmt.Sprint(c.err) {
			t.Errorf("%sTest %d failed%s\twant: %s, got: %s\n", red, id, normal, c.err, err)
		} else if plot != c.plot {
			t.Errorf("%sTest %d failed%s\twant: %s, got: %s\n", red, id, normal, c.plot, plot)
			_ = os.WriteFile(fmt.Sprintf("TestDrawAngularFromZero%d.svg", id), []byte(plot), 777)
		} else {
			t.Logf("%sTest %d  success%s\t%s\n", green, id, normal, "")
		}

	}
	testArr := []testCase{
		{
			TotalHeight:   200,
			TotalWidth:    400,
			xDivisionsQty: 10,
			yDivisionsQty: 10,
			NameOfX:       "x",
			NameOfY:       "y",
			X:             []float64{0, 1, 2, 3},
			Y:             []float64{0, 1, 4, 9},
			plot:          TestDrawAngularFromZero0,
		},
		{
			TotalHeight:   200,
			TotalWidth:    400,
			xDivisionsQty: 10,
			yDivisionsQty: 10,
			NameOfX:       "x",
			NameOfY:       "y",
			X:             []float64{3, 0, 1, 2, 3},
			Y:             []float64{9, 4, 1, 0, 1, 4, 9},
			err:           errors.New("len(x)!=len(y)"),
		},
		{
			TotalHeight:   200,
			TotalWidth:    400,
			xDivisionsQty: 10,
			yDivisionsQty: 10,
			NameOfX:       "x",
			NameOfY:       "y",
			X:             []float64{},
			Y:             []float64{9, 4, 1, 0, 1, 4, 9},
			err:           errors.New("empty input"),
		},
		{
			TotalHeight:   200,
			TotalWidth:    400,
			xDivisionsQty: 10,
			yDivisionsQty: 10,
			NameOfX:       "x",
			NameOfY:       "y",
			X:             []float64{9, 4, 1, 0, 1, 4, 9},
			Y:             []float64{},
			err:           errors.New("empty input"),
		},
		{
			TotalHeight:   200,
			TotalWidth:    400,
			xDivisionsQty: 10,
			yDivisionsQty: 10,
			NameOfX:       "x",
			NameOfY:       "y",
			X:             []float64{9, 4, 1, 0, 1, 4, 9},
			Y:             []float64{9, 4, 1, 0, 1, 4, -9},
			err:           errors.New("6-th element is les, then 0"),
		},

		{
			TotalHeight:   200,
			TotalWidth:    400,
			xDivisionsQty: 10,
			yDivisionsQty: 10,
			NameOfX:       "x",
			NameOfY:       "y",
			X:             []float64{9, 4, 1, 0, 1, 4, -9},
			Y:             []float64{9, 4, 1, 0, 1, 4, 9},
			err:           errors.New("6-th element is les, then 0"),
		},
	}
	for key, val := range testArr {
		doTest(key, val)
	}
}
