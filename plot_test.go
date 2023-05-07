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
		res := convertSliceToFloat64([]int{0, 1, 2, 3, 4})
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
		res := convertSliceToFloat64([]float32{-2, -1, 0, 1, 2, 0.5, 4})
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
		res := convertSliceToFloat64([]uint{0, 1, 2, 3, 4})
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
			_ = os.WriteFile(fmt.Sprintf("test%d.svg", testId), []byte(plot), 777)
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
			plot: `<?xml version="1.0" encoding="UTF-8"?>
<svg width="400" height="200" viewBox="0 0 400 200" version="1.1" xmlns="http://www.w3.org/2000/svg">
"<style>
.axis {
font-family="Arial, Helvetica, sans-serif">;
font-size: 12pt;
}
</style>
<line x1="16" y1="0" x2="400" y2="0" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="9">y</text>
<line x1="16" y1="18.7" x2="400" y2="18.7" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="22.7">9</text>
<line x1="16" y1="37.4" x2="400" y2="37.4" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="41.4">8</text>
<line x1="16" y1="56.1" x2="400" y2="56.1" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="60.1">7</text>
<line x1="16" y1="74.8" x2="400" y2="74.8" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="78.8">6</text>
<line x1="16" y1="93.5" x2="400" y2="93.5" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="97.5">5</text>
<line x1="16" y1="112.2" x2="400" y2="112.2" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="116.2">4</text>
<line x1="16" y1="130.9" x2="400" y2="130.9" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="134.9">3</text>
<line x1="16" y1="149.6" x2="400" y2="149.6" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="153.6">2</text>
<line x1="16" y1="168.3" x2="400" y2="168.3" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="172.3">1</text>
<line x1="16" y1="187" x2="400" y2="187" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="191">0</text>
<line x1="16" y1="0" x2="16" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="16" y="200">-5</text>
<line x1="54.4" y1="0" x2="54.4" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="54.4" y="200">-4</text>
<line x1="92.8" y1="0" x2="92.8" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="92.8" y="200">-3</text>
<line x1="131.2" y1="0" x2="131.2" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="131.2" y="200">-2</text>
<line x1="169.6" y1="0" x2="169.6" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="169.6" y="200">-1</text>
<line x1="208" y1="0" x2="208" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="208" y="200">0</text>
<line x1="246.4" y1="0" x2="246.4" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="246.4" y="200">1</text>
<line x1="284.8" y1="0" x2="284.8" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="284.8" y="200">2</text>
<line x1="323.2" y1="0" x2="323.2" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="323.2" y="200">3</text>
<line x1="361.6" y1="0" x2="361.6" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="361.6" y="200">4</text>
<line x1="400" y1="0" x2="400" y2="187" stroke="black" />
<text text-anchor="end" class="axis" x="400" y="200">x</text>

<path d="M 92.80000000000001 18.69999999999999 L 131.2 112.2 L 169.6 168.3 L 208 187 L 246.4 168.3 L 284.8 112.2 L 323.2 18.69999999999999" fill="none" stroke="black" stroke-width="3"/>
</svg>`,
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
			plot: `<?xml version="1.0" encoding="UTF-8"?>
<svg width="400" height="200" viewBox="0 0 400 200" version="1.1" xmlns="http://www.w3.org/2000/svg">
"<style>
.axis {
font-family="Arial, Helvetica, sans-serif">;
font-size: 12pt;
}
</style>
<line x1="23" y1="0" x2="400" y2="0" stroke="black" />
<text text-anchor="end" class="axis" x="19" y="9">10</text>
<line x1="23" y1="18.7" x2="400" y2="18.7" stroke="black" />
<text text-anchor="end" class="axis" x="19" y="22.7">9</text>
<line x1="23" y1="37.4" x2="400" y2="37.4" stroke="black" />
<text text-anchor="end" class="axis" x="19" y="41.4">8</text>
<line x1="23" y1="56.1" x2="400" y2="56.1" stroke="black" />
<text text-anchor="end" class="axis" x="19" y="60.1">7</text>
<line x1="23" y1="74.8" x2="400" y2="74.8" stroke="black" />
<text text-anchor="end" class="axis" x="19" y="78.8">6</text>
<line x1="23" y1="93.5" x2="400" y2="93.5" stroke="black" />
<text text-anchor="end" class="axis" x="19" y="97.5">5</text>
<line x1="23" y1="112.2" x2="400" y2="112.2" stroke="black" />
<text text-anchor="end" class="axis" x="19" y="116.2">4</text>
<line x1="23" y1="130.9" x2="400" y2="130.9" stroke="black" />
<text text-anchor="end" class="axis" x="19" y="134.9">3</text>
<line x1="23" y1="149.6" x2="400" y2="149.6" stroke="black" />
<text text-anchor="end" class="axis" x="19" y="153.6">2</text>
<line x1="23" y1="168.3" x2="400" y2="168.3" stroke="black" />
<text text-anchor="end" class="axis" x="19" y="172.3">1</text>
<line x1="23" y1="187" x2="400" y2="187" stroke="black" />
<text text-anchor="end" class="axis" x="19" y="191">0</text>
<line x1="23" y1="0" x2="23" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="23" y="200">-5</text>
<line x1="60.7" y1="0" x2="60.7" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="60.7" y="200">-4</text>
<line x1="98.4" y1="0" x2="98.4" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="98.4" y="200">-3</text>
<line x1="136.1" y1="0" x2="136.1" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="136.1" y="200">-2</text>
<line x1="173.8" y1="0" x2="173.8" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="173.8" y="200">-1</text>
<line x1="211.5" y1="0" x2="211.5" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="211.5" y="200">0</text>
<line x1="249.2" y1="0" x2="249.2" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="249.2" y="200">1</text>
<line x1="286.9" y1="0" x2="286.9" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="286.9" y="200">2</text>
<line x1="324.6" y1="0" x2="324.6" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="324.6" y="200">3</text>
<line x1="362.3" y1="0" x2="362.3" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="362.3" y="200">4</text>
<line x1="400" y1="0" x2="400" y2="187" stroke="black" />
<text text-anchor="end" class="axis" x="400" y="200">5</text>

<path d="M 98.4 18.69999999999999 L 136.10000000000002 112.2 L 173.8 168.3 L 211.5 187 L 249.2 168.3 L 286.9 112.2 L 324.6 18.69999999999999" fill="none" stroke="black" stroke-width="3"/>
</svg>`,
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
	var testId uint64
	{
		testId = 0
		t.Logf("Test %d\tstart", testId)
		check := `<?xml version="1.0" encoding="UTF-8"?>
<svg width="400" height="200" viewBox="0 0 400 200" version="1.1" xmlns="http://www.w3.org/2000/svg">
"<style>
.axis {
font-family="Arial, Helvetica, sans-serif">;
font-size: 12pt;
}
</style>
<line x1="16" y1="0" x2="400" y2="0" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="9">y</text>
<line x1="16" y1="20.78" x2="400" y2="20.78" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="24.78">8</text>
<line x1="16" y1="41.56" x2="400" y2="41.56" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="45.56">7</text>
<line x1="16" y1="62.33" x2="400" y2="62.33" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="66.33">6</text>
<line x1="16" y1="83.11" x2="400" y2="83.11" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="87.11">5</text>
<line x1="16" y1="103.89" x2="400" y2="103.89" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="107.89">4</text>
<line x1="16" y1="124.67" x2="400" y2="124.67" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="128.67">3</text>
<line x1="16" y1="145.44" x2="400" y2="145.44" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="149.44">2</text>
<line x1="16" y1="166.22" x2="400" y2="166.22" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="170.22">1</text>
<line x1="16" y1="187" x2="400" y2="187" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="191">0</text>
<line x1="16" y1="0" x2="16" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="16" y="200">-3</text>
<line x1="48" y1="0" x2="48" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="48" y="200">-2.5</text>
<line x1="80" y1="0" x2="80" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="80" y="200">-2</text>
<line x1="112" y1="0" x2="112" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="112" y="200">-1.5</text>
<line x1="144" y1="0" x2="144" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="144" y="200">-1</text>
<line x1="176" y1="0" x2="176" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="176" y="200">-0.5</text>
<line x1="208" y1="0" x2="208" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="208" y="200">0</text>
<line x1="240" y1="0" x2="240" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="240" y="200">0.5</text>
<line x1="272" y1="0" x2="272" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="272" y="200">1</text>
<line x1="304" y1="0" x2="304" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="304" y="200">1.5</text>
<line x1="336" y1="0" x2="336" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="336" y="200">2</text>
<line x1="368" y1="0" x2="368" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="368" y="200">2.5</text>
<line x1="400" y1="0" x2="400" y2="187" stroke="black" />
<text text-anchor="end" class="axis" x="400" y="200">x</text>

<path d="M 16 0 L 80 103.88888888888889 L 144 166.22222222222223 L 208 187 L 272 166.22222222222223 L 336 103.88888888888889 L 400 0" fill="none" stroke="black" stroke-width="3"/>
</svg>`
		plot, err := DrawAngular(200, 400, 10, 10, []float64{-3, -2, -1, 0, 1, 2, 3}, []float64{9, 4, 1, 0, 1, 4, 9}, "x", "y")
		if err != nil {
			t.Errorf("%sTest %d failed%s\tgot: %s, want: nil\n", red, testId, normal, err)
		} else if plot != check {
			t.Errorf("%sTest %d failed%s\tgot: %s, want: %s\n", red, testId, normal, plot, check)
			_ = os.WriteFile(fmt.Sprintf("test%d.svg", testId), []byte(plot), 777)
		} else {
			t.Logf("%sTest %d  success%s\t%s\n", green, testId, normal, "")
		}
	}
	{
		testId = 1
		t.Logf("Test %d\tstart", testId)
		check := `<?xml version="1.0" encoding="UTF-8"?>
<svg width="400" height="200" viewBox="0 0 400 200" version="1.1" xmlns="http://www.w3.org/2000/svg">
"<style>
.axis {
font-family="Arial, Helvetica, sans-serif">;
font-size: 12pt;
}
</style>
<line x1="16" y1="0" x2="400" y2="0" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="9">y</text>
<line x1="16" y1="20.78" x2="400" y2="20.78" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="24.78">8</text>
<line x1="16" y1="41.56" x2="400" y2="41.56" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="45.56">7</text>
<line x1="16" y1="62.33" x2="400" y2="62.33" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="66.33">6</text>
<line x1="16" y1="83.11" x2="400" y2="83.11" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="87.11">5</text>
<line x1="16" y1="103.89" x2="400" y2="103.89" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="107.89">4</text>
<line x1="16" y1="124.67" x2="400" y2="124.67" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="128.67">3</text>
<line x1="16" y1="145.44" x2="400" y2="145.44" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="149.44">2</text>
<line x1="16" y1="166.22" x2="400" y2="166.22" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="170.22">1</text>
<line x1="16" y1="187" x2="400" y2="187" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="191">0</text>
<line x1="16" y1="0" x2="16" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="16" y="200">-3</text>
<line x1="48" y1="0" x2="48" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="48" y="200">-2.5</text>
<line x1="80" y1="0" x2="80" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="80" y="200">-2</text>
<line x1="112" y1="0" x2="112" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="112" y="200">-1.5</text>
<line x1="144" y1="0" x2="144" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="144" y="200">-1</text>
<line x1="176" y1="0" x2="176" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="176" y="200">-0.5</text>
<line x1="208" y1="0" x2="208" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="208" y="200">0</text>
<line x1="240" y1="0" x2="240" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="240" y="200">0.5</text>
<line x1="272" y1="0" x2="272" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="272" y="200">1</text>
<line x1="304" y1="0" x2="304" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="304" y="200">1.5</text>
<line x1="336" y1="0" x2="336" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="336" y="200">2</text>
<line x1="368" y1="0" x2="368" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="368" y="200">2.5</text>
<line x1="400" y1="0" x2="400" y2="187" stroke="black" />
<text text-anchor="end" class="axis" x="400" y="200">x</text>

<path d="M 16 0 L 80 103.88888888888889 L 144 166.22222222222223 L 208 187 L 272 166.22222222222223 L 336 103.88888888888889 L 400 0" fill="none" stroke="black" stroke-width="3"/>
</svg>`
		plot, err := DrawAngular(200, 400, 10, 10, []int{-3, -2, -1, 0, 1, 2, 3}, []int{9, 4, 1, 0, 1, 4, 9}, "x", "y")
		if err != nil {
			t.Errorf("%sTest %d failed%s\tgot: %s, want: nil\n", red, testId, normal, err)
		} else if plot != check {
			t.Errorf("%sTest %d failed%s\tgot: %s, want: %s\n", red, testId, normal, plot, check)
			_ = os.WriteFile(fmt.Sprintf("test%d.svg", testId), []byte(plot), 777)
		} else {
			t.Logf("%sTest %d  success%s\t%s\n", green, testId, normal, "")
		}
	}
	{
		testId = 2
		t.Logf("Test %d\tstart", testId)
		check := `<?xml version="1.0" encoding="UTF-8"?>
<svg width="400" height="200" viewBox="0 0 400 200" version="1.1" xmlns="http://www.w3.org/2000/svg">
"<style>
.axis {
font-family="Arial, Helvetica, sans-serif">;
font-size: 12pt;
}
</style>
<line x1="16" y1="0" x2="400" y2="0" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="9">y</text>
<line x1="16" y1="20.78" x2="400" y2="20.78" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="24.78">8</text>
<line x1="16" y1="41.56" x2="400" y2="41.56" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="45.56">7</text>
<line x1="16" y1="62.33" x2="400" y2="62.33" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="66.33">6</text>
<line x1="16" y1="83.11" x2="400" y2="83.11" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="87.11">5</text>
<line x1="16" y1="103.89" x2="400" y2="103.89" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="107.89">4</text>
<line x1="16" y1="124.67" x2="400" y2="124.67" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="128.67">3</text>
<line x1="16" y1="145.44" x2="400" y2="145.44" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="149.44">2</text>
<line x1="16" y1="166.22" x2="400" y2="166.22" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="170.22">1</text>
<line x1="16" y1="187" x2="400" y2="187" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="191">0</text>
<line x1="16" y1="0" x2="16" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="16" y="200">-3</text>
<line x1="48" y1="0" x2="48" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="48" y="200">-2.5</text>
<line x1="80" y1="0" x2="80" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="80" y="200">-2</text>
<line x1="112" y1="0" x2="112" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="112" y="200">-1.5</text>
<line x1="144" y1="0" x2="144" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="144" y="200">-1</text>
<line x1="176" y1="0" x2="176" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="176" y="200">-0.5</text>
<line x1="208" y1="0" x2="208" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="208" y="200">0</text>
<line x1="240" y1="0" x2="240" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="240" y="200">0.5</text>
<line x1="272" y1="0" x2="272" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="272" y="200">1</text>
<line x1="304" y1="0" x2="304" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="304" y="200">1.5</text>
<line x1="336" y1="0" x2="336" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="336" y="200">2</text>
<line x1="368" y1="0" x2="368" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="368" y="200">2.5</text>
<line x1="400" y1="0" x2="400" y2="187" stroke="black" />
<text text-anchor="end" class="axis" x="400" y="200">x</text>

<path d="M 16 0 L 80 103.88888888888889 L 144 166.22222222222223 L 208 187 L 272 166.22222222222223 L 336 103.88888888888889 L 400 0" fill="none" stroke="black" stroke-width="3"/>
</svg>`
		plot, err := DrawAngular(200, 400, 10, 10, []float32{-3, -2, -1, 0, 1, 2, 3}, []float32{9, 4, 1, 0, 1, 4, 9}, "x", "y")
		if err != nil {
			t.Errorf("%sTest %d failed%s\tgot: %s, want: nil\n", red, testId, normal, err)
		} else if plot != check {
			t.Errorf("%sTest %d failed%s\tgot: %s, want: %s\n", red, testId, normal, plot, check)
			_ = os.WriteFile(fmt.Sprintf("test%d.svg", testId), []byte(plot), 777)
		} else {
			t.Logf("%sTest %d  success%s\t%s\n", green, testId, normal, "")
		}
	}
	{
		testId = 3
		t.Logf("Test %d\tstart", testId)
		_, err := DrawAngular(200, 400, 10, 10, []float32{-3, 0, 1, 2, 3}, []float32{9, 4, 1, 0, 1, 4, 9}, "x", "y")
		if err == nil {
			t.Errorf("%sTest %d failed%s\tgot: nil, want: error\n", red, testId, normal)
		} else {
			t.Logf("%sTest %d  success%s\t%s\n", green, testId, normal, "")
		}
	}
	{
		testId = 4
		t.Logf("Test %d\tstart", testId)
		_, err := DrawAngular(200, 400, 10, 10, []float32{}, []float32{9, 4, 1, 0, 1, 4, 9}, "x", "y")
		if err == nil {
			t.Errorf("%sTest %d failed%s\tgot: nil, want: error\n", red, testId, normal)
		} else {
			t.Logf("%sTest %d  success%s\t%s\n", green, testId, normal, "")
		}
	}
	{
		testId = 5
		t.Logf("Test %d\tstart", testId)
		_, err := DrawAngular(200, 400, 10, 10, []float32{9, 4, 1, 0, 1, 4, 9}, []float32{}, "x", "y")
		if err == nil {
			t.Errorf("%sTest %d failed%s\tgot: nil, want: error\n", red, testId, normal)
		} else {
			t.Logf("%sTest %d  success%s\t%s\n", green, testId, normal, "")
		}
	}
}
func TestDrawAngularFrom0(t *testing.T) {
	var testId uint64
	{
		testId = 0
		t.Logf("Test %d\tstart", testId)
		check := `<?xml version="1.0" encoding="UTF-8"?>
<svg width="200" height="200" viewBox="0 0 200 200" version="1.1" xmlns="http://www.w3.org/2000/svg">
"<style>
.axis {
font-family="Arial, Helvetica, sans-serif">;
font-size: 12pt;
}
</style>
<line x1="16" y1="0" x2="200" y2="0" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="9">y</text>
<line x1="16" y1="37.4" x2="200" y2="37.4" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="41.4">8</text>
<line x1="16" y1="74.8" x2="200" y2="74.8" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="78.8">6</text>
<line x1="16" y1="112.2" x2="200" y2="112.2" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="116.2">4</text>
<line x1="16" y1="149.6" x2="200" y2="149.6" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="153.6">2</text>
<line x1="16" y1="187" x2="200" y2="187" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="191">0</text>
<line x1="16" y1="0" x2="16" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="16" y="200">0</text>
<line x1="46.67" y1="0" x2="46.67" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="46.67" y="200">0.5</text>
<line x1="77.33" y1="0" x2="77.33" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="77.33" y="200">1</text>
<line x1="108" y1="0" x2="108" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="108" y="200">1.5</text>
<line x1="138.67" y1="0" x2="138.67" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="138.67" y="200">2</text>
<line x1="169.33" y1="0" x2="169.33" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="169.33" y="200">2.5</text>
<line x1="200" y1="0" x2="200" y2="187" stroke="black" />
<text text-anchor="end" class="axis" x="200" y="200">x</text>

<path d="M 16 187 L 77.33333333333334 168.3 L 138.66666666666669 112.2 L 200 18.69999999999999" fill="none" stroke="black" stroke-width="3"/>
</svg>`
		plot, err := DrawAngularFrom0(200, 200, 5, 5, []float64{0, 1, 2, 3}, []float64{0, 1, 4, 9}, "x", "y")
		if err != nil {
			t.Errorf("%sTest %d failed%s\tgot: %s, want: nil\n", red, testId, normal, err)
		} else if plot != check {
			t.Errorf("%sTest %d failed%s\tgot: %s, want: %s\n", red, testId, normal, plot, check)
			_ = os.WriteFile(fmt.Sprintf("TestDrawAngularFrom0-%d.svg", testId), []byte(plot), 777)
		} else {
			t.Logf("%sTest %d  success%s\t%s\n", green, testId, normal, "")
		}
	}
	{
		testId = 0
		t.Logf("Test %d\tstart", testId)
		check := `<?xml version="1.0" encoding="UTF-8"?>
<svg width="200" height="200" viewBox="0 0 200 200" version="1.1" xmlns="http://www.w3.org/2000/svg">
"<style>
.axis {
font-family="Arial, Helvetica, sans-serif">;
font-size: 12pt;
}
</style>
<line x1="16" y1="0" x2="200" y2="0" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="9">y</text>
<line x1="16" y1="37.4" x2="200" y2="37.4" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="41.4">8</text>
<line x1="16" y1="74.8" x2="200" y2="74.8" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="78.8">6</text>
<line x1="16" y1="112.2" x2="200" y2="112.2" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="116.2">4</text>
<line x1="16" y1="149.6" x2="200" y2="149.6" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="153.6">2</text>
<line x1="16" y1="187" x2="200" y2="187" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="191">0</text>
<line x1="16" y1="0" x2="16" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="16" y="200">0</text>
<line x1="46.67" y1="0" x2="46.67" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="46.67" y="200">0.5</text>
<line x1="77.33" y1="0" x2="77.33" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="77.33" y="200">1</text>
<line x1="108" y1="0" x2="108" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="108" y="200">1.5</text>
<line x1="138.67" y1="0" x2="138.67" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="138.67" y="200">2</text>
<line x1="169.33" y1="0" x2="169.33" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="169.33" y="200">2.5</text>
<line x1="200" y1="0" x2="200" y2="187" stroke="black" />
<text text-anchor="end" class="axis" x="200" y="200">x</text>

<path d="M 16 187 L 77.33333333333334 168.3 L 138.66666666666669 112.2 L 200 18.69999999999999" fill="none" stroke="black" stroke-width="3"/>
</svg>`
		plot, err := DrawAngularFrom0(200, 200, 5, 5, []uint8{0, 1, 2, 3}, []uint8{0, 1, 4, 9}, "x", "y")
		if err != nil {
			t.Errorf("%sTest %d failed%s\tgot: %s, want: nil\n", red, testId, normal, err)
		} else if plot != check {
			t.Errorf("%sTest %d failed%s\tgot: %s, want: %s\n", red, testId, normal, plot, check)
			_ = os.WriteFile(fmt.Sprintf("TestDrawAngularFrom0-%d.svg", testId), []byte(plot), 777)
		} else {
			t.Logf("%sTest %d  success%s\t%s\n", green, testId, normal, "")
		}
	}
	{
		testId = 0
		t.Logf("Test %d\tstart", testId)
		check := `<?xml version="1.0" encoding="UTF-8"?>
<svg width="200" height="200" viewBox="0 0 200 200" version="1.1" xmlns="http://www.w3.org/2000/svg">
"<style>
.axis {
font-family="Arial, Helvetica, sans-serif">;
font-size: 12pt;
}
</style>
<line x1="16" y1="0" x2="200" y2="0" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="9">y</text>
<line x1="16" y1="37.4" x2="200" y2="37.4" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="41.4">8</text>
<line x1="16" y1="74.8" x2="200" y2="74.8" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="78.8">6</text>
<line x1="16" y1="112.2" x2="200" y2="112.2" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="116.2">4</text>
<line x1="16" y1="149.6" x2="200" y2="149.6" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="153.6">2</text>
<line x1="16" y1="187" x2="200" y2="187" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="191">0</text>
<line x1="16" y1="0" x2="16" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="16" y="200">0</text>
<line x1="46.67" y1="0" x2="46.67" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="46.67" y="200">0.5</text>
<line x1="77.33" y1="0" x2="77.33" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="77.33" y="200">1</text>
<line x1="108" y1="0" x2="108" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="108" y="200">1.5</text>
<line x1="138.67" y1="0" x2="138.67" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="138.67" y="200">2</text>
<line x1="169.33" y1="0" x2="169.33" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="169.33" y="200">2.5</text>
<line x1="200" y1="0" x2="200" y2="187" stroke="black" />
<text text-anchor="end" class="axis" x="200" y="200">x</text>

<path d="M 16 187 L 77.33333333333334 168.3 L 138.66666666666669 112.2 L 200 18.69999999999999" fill="none" stroke="black" stroke-width="3"/>
</svg>`
		plot, err := DrawAngularFrom0(200, 200, 5, 5, []float32{0, 1, 2, 3}, []float32{0, 1, 4, 9}, "x", "y")
		if err != nil {
			t.Errorf("%sTest %d failed%s\tgot: %s, want: nil\n", red, testId, normal, err)
		} else if plot != check {
			t.Errorf("%sTest %d failed%s\tgot: %s, want: %s\n", red, testId, normal, plot, check)
			_ = os.WriteFile(fmt.Sprintf("TestDrawAngularFrom0-%d.svg", testId), []byte(plot), 777)
		} else {
			t.Logf("%sTest %d  success%s\t%s\n", green, testId, normal, "")
		}
	}
	{
		testId = 3
		t.Logf("Test %d\tstart", testId)
		_, err := DrawAngularFrom0(200, 400, 10, 10, []float32{3, 0, 1, 2, 3}, []float32{9, 4, 1, 0, 1, 4, 9}, "x", "y")
		if err == nil {
			t.Errorf("%sTest %d failed%s\tgot: nil, want: error\n", red, testId, normal)
		} else {
			t.Logf("%sTest %d  success%s\t%s\n", green, testId, normal, "")
		}
	}
	{
		testId = 4
		t.Logf("Test %d\tstart", testId)
		_, err := DrawAngularFrom0(200, 400, 10, 10, []float32{}, []float32{9, 4, 1, 0, 1, 4, 9}, "x", "y")
		if err == nil {
			t.Errorf("%sTest %d failed%s\tgot: nil, want: error\n", red, testId, normal)
		} else {
			t.Logf("%sTest %d  success%s\t%s\n", green, testId, normal, "")
		}
	}
	{
		testId = 5
		t.Logf("Test %d\tstart", testId)
		_, err := DrawAngularFrom0(200, 400, 10, 10, []float32{9, 4, 1, 0, 1, 4, 9}, []float32{}, "x", "y")
		if err == nil {
			t.Errorf("%sTest %d failed%s\tgot: nil, want: error\n", red, testId, normal)
		} else {
			t.Logf("%sTest %d  success%s\t%s\n", green, testId, normal, "")
		}
	}
	{
		testId = 5
		t.Logf("Test %d\tstart", testId)
		_, err := DrawAngularFrom0(200, 400, 10, 10, []float32{1, -2, 3}, []float32{1, 2, 3}, "x", "y")
		if err == nil {
			t.Errorf("%sTest %d failed%s\tgot: nil, want: error\n", red, testId, normal)
		} else {
			t.Logf("%sTest %d  success%s\t%s\n", green, testId, normal, "")
		}
	}
	{
		testId = 5
		t.Logf("Test %d\tstart", testId)
		_, err := DrawAngularFrom0(200, 400, 10, 10, []float32{1, 2, 3}, []float32{1, -2, 3}, "x", "y")
		if err == nil {
			t.Errorf("%sTest %d failed%s\tgot: nil, want: error\n", red, testId, normal)
		} else {
			t.Logf("%sTest %d  success%s\t%s\n", green, testId, normal, "")
		}
	}
}
