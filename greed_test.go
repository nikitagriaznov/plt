// Package svgPlot Copyright 2023 Gryaznov Nikita Licensed under the Apache
// License, Version 2.0 (the «License»);
package svgPlot

import "testing"

func TestGetMinMax(t *testing.T) {
	var testId uint
	test := func(in []float64, rightMin, rightMax float64, isErr bool) {
		t.Logf("Test %d start", testId)
		min, max, err := getMinMax(in)
		if isErr {
			if err != nil {
				t.Logf("%sTest %d seccess%s\t%s\n", green, testId, normal, "")
			} else {
				t.Errorf("%sTest %d failed%s\tgot: nil, want: error\n", red, testId, normal)
			}
		} else {
			if err != nil {
				t.Errorf("%sTest %d failed%s\tgot: %s, want: nil\n", red, testId, normal, err)
			} else if min != rightMin {
				t.Errorf("%sTest %d failed%s\tgot: %G, want: %G\n", red, testId, normal, rightMin, min)
			} else if max != rightMax {
				t.Errorf("%sTest %d failed%s\tgot: %G, want: %G\n", red, testId, normal, rightMax, max)
			} else {
				t.Logf("%sTest %d seccess%s\t%s\n", green, testId, normal, "")
			}
		}
		testId++
	}
	test([]float64{1, 3, 5, 2, 2, 66, 34, 0, -2.4}, -2.4, 66, false)
	test([]float64{}, 0, 0, true)
	test([]float64{1, 3, 5, 2, 2, 66, 34, 0, -0.00004}, -0.00004, 66, false)
}
func TestGetExp(t *testing.T) {
	var testId uint
	test := func(in, rightOut float64) {
		t.Logf("Test %d start", testId)
		res := getExp(in)
		ok := rightOut == res
		if ok {
			t.Logf("%sTest %d seccess%s\t%s\n", green, testId, normal, "")
		} else {
			t.Errorf("%sTest %d failed%s\tgot: %G, want: %G\n", red, testId, normal, rightOut, res)
		}
		testId++
	}
	test(7, 1)
	test(30, 10)
	test(700, 100)
	test(-237, 100)
	test(-0.234, 0.1)
	test(0.00323, 0.001)
}
func TestMakeArr(t *testing.T) {
	equal := func(a, b []string) bool {
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
	test := func(min, max float64, divQty uint, rightArr []string, rightLen float64, rightZero int) {
		t.Logf("Test %d\tstart", testId)
		res, l, z, err := makeArr(min, max, divQty)
		ok := equal(rightArr, res)
		if err != nil {
			t.Errorf("%sTest %d failed%s\t%s", red, testId, normal, err)
		} else if !ok {
			t.Errorf("%sTest %d failed%s\tgot: %s, want: %s\n", red, testId, normal, rightArr, res)
		} else if l != rightLen {
			t.Errorf("%sTest %d failed%s\tgot: %f, want: %f\n", red, testId, normal, rightLen, l)
		} else if z != rightZero {
			t.Errorf("%sTest %d failed%s\tgot: %d, want: %d\n", red, testId, normal, rightZero, z)
		} else {
			t.Logf("%sTest %d seccess%s\t%s\n", green, testId, normal, "")
		}
		testId++
	}
	test(0, 7, 10, []string{"0", "1", "2", "3", "4", "5", "6", "7"}, 7, 0)
	test(-5, 9, 10, []string{"-6", "-4", "-2", "0", "2", "4", "6", "8", "10"}, 16, 3)
	test(-5, 9, 5, []string{"-5", "-2.5", "0", "2.5", "5", "7.5", "10"}, 15, 2)
}
func TestGetWordLen(t *testing.T) {
	var testId uint64
	test := func(word string, rightLen int) {
		t.Logf("Test %d\tstart", testId)
		res := getWordLen(word)
		ok := rightLen == res
		if ok {
			t.Logf("%sTest %d seccess%s\t%s\n", green, testId, normal, "")
		} else {
			t.Errorf("%sTest %d failed%s\tgot: %d, want: %d\n", red, testId, normal, rightLen, res)
		}
		testId++
	}
	test("7", 7)
	test("93", 14)
	test("0.83", 24)
	test("0.123", 31)
	test("qty", 17)
}
func TestMakeGreed(t *testing.T) {
	type testType struct {
		height, width        uint
		xNumArray            []string
		yNumArray            []string
		xLen, yLen           float64
		xZeroPos, yZeroPos   int
		greed                string
		x0, y0, gradX, gradY float64
		err                  error
	}
	var testId uint64
	test := func(t1 testType) {
		t.Logf("Test %d start", testId)
		greed, x0, y0, gradX, gradY, err := makeGreed(t1.height, t1.width, t1.xNumArray, t1.yNumArray, t1.xLen, t1.yLen, t1.xZeroPos, t1.yZeroPos)

		if err != t1.err {
			t.Errorf("%sTest %d failed%s\tgot: %s, want: %s\n", red, testId, normal, t1.err, err)
		} else if greed != t1.greed {
			t.Errorf("%sTest %d failed%s\tgot: %s, want: %s\n", red, testId, normal, t1.greed, greed)
		} else if x0 != t1.x0 {
			t.Errorf("%sTest %d failed%s\tgot: %g, want: %g\n", red, testId, normal, t1.x0, x0)
		} else if y0 != t1.y0 {
			t.Errorf("%sTest %d failed%s\tgot: %g, want: %g\n", red, testId, normal, t1.y0, y0)
		} else if gradX != t1.gradX {
			t.Errorf("%sTest %d failed%s\tgot: %g, want: %g\n", red, testId, normal, t1.gradX, gradX)
		} else if gradY != t1.gradY {
			t.Errorf("%sTest %d failed%s\tgot: %g, want: %g\n", red, testId, normal, t1.gradY, gradY)
		} else {
			t.Logf("%sTest %d seccess%s\t%s\n", green, testId, normal, "")
		}
		testId++
	}
	testArr := []testType{
		{
			height:    200,
			width:     400,
			xNumArray: []string{"-5", "-4", "-3", "-2", "-1", "0", "1", "2", "3", "4", "4", "5"},
			yNumArray: []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10"},
			xLen:      10,
			yLen:      10,
			xZeroPos:  5,
			yZeroPos:  0,
			greed: `<style>
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
<line x1="57.27" y1="0" x2="57.27" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="57.27" y="200">-4</text>
<line x1="91.55" y1="0" x2="91.55" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="91.55" y="200">-3</text>
<line x1="125.82" y1="0" x2="125.82" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="125.82" y="200">-2</text>
<line x1="160.09" y1="0" x2="160.09" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="160.09" y="200">-1</text>
<line x1="194.36" y1="0" x2="194.36" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="194.36" y="200">0</text>
<line x1="228.64" y1="0" x2="228.64" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="228.64" y="200">1</text>
<line x1="262.91" y1="0" x2="262.91" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="262.91" y="200">2</text>
<line x1="297.18" y1="0" x2="297.18" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="297.18" y="200">3</text>
<line x1="331.45" y1="0" x2="331.45" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="331.45" y="200">4</text>
<line x1="365.73" y1="0" x2="365.73" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="365.73" y="200">4</text>
<line x1="400" y1="0" x2="400" y2="187" stroke="black" />
<text text-anchor="end" class="axis" x="400" y="200">5</text>
`,
			x0:    194.36363636363637,
			y0:    187,
			gradX: 0.026525198938992044,
			gradY: 0.053475935828877004,
		},
	}
	for _, testCase := range testArr {
		test(testCase)
	}
}
