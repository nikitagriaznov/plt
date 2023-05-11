// Package svgPlot Copyright 2023 Gryaznov Nikita
// Licensed under the Apache License, Version 2.0
package svgPlot

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestNewPlot(t *testing.T) {
	type TestCase struct {
		TotalHeight, TotalWidth      uint
		XDivisionsQty, YDivisionsQty uint
		NameOfX, NameOfY             string
		plot                         Plot
		err                          error
	}

	doTest := func(c TestCase, testId int) {
		t.Logf("Test %d start", testId)
		plot, err := NewPlot(c.TotalHeight, c.TotalWidth, c.XDivisionsQty, c.YDivisionsQty, c.NameOfX, c.NameOfY)

		if fmt.Sprint(err) != fmt.Sprint(c.err) {
			t.Errorf("%sTest %d failed%s\tgot: %s, want: %s\n", red, testId, normal, c.err, err)
		} else if !reflect.DeepEqual(plot, c.plot) {
			t.Errorf("%sTest %d failed%s\twant: %s, got: %s\n", red, testId, normal, fmt.Sprint(c.plot), fmt.Sprint(plot))
		} else {
			t.Logf("%sTest %d success%s\t%s\n", green, testId, normal, "")
		}
	}
	testArray := []TestCase{
		{
			TotalHeight:   200,
			TotalWidth:    200,
			XDivisionsQty: 12,
			YDivisionsQty: 12,
			NameOfX:       "",
			NameOfY:       "",
			plot: Plot{
				style: Style{
					TotalHeight:   200,
					TotalWidth:    200,
					XDivisionsQty: 12,
					YDivisionsQty: 12,
					NameOfX:       "",
					NameOfY:       "",
				},
			},
			err: nil,
		},
		{
			TotalHeight:   100,
			TotalWidth:    200,
			XDivisionsQty: 12,
			YDivisionsQty: 12,
			NameOfX:       "",
			NameOfY:       "",
			plot:          Plot{},
			err:           errors.New("TotalHeight is too low"),
		},
		{
			TotalHeight:   200,
			TotalWidth:    100,
			XDivisionsQty: 12,
			YDivisionsQty: 12,
			NameOfX:       "",
			NameOfY:       "",
			plot:          Plot{},
			err:           errors.New("TotalWidth is too low"),
		},
		{
			TotalHeight:   200,
			TotalWidth:    200,
			XDivisionsQty: 52,
			YDivisionsQty: 12,
			NameOfX:       "",
			NameOfY:       "",
			plot:          Plot{},
			err:           errors.New("XDivisionsQty is too high"),
		},
		{
			TotalHeight:   200,
			TotalWidth:    200,
			XDivisionsQty: 12,
			YDivisionsQty: 52,
			NameOfX:       "",
			NameOfY:       "",
			plot:          Plot{},
			err:           errors.New("YDivisionsQty is too high"),
		},
		{
			TotalHeight:   200,
			TotalWidth:    200,
			XDivisionsQty: 12,
			YDivisionsQty: 12,
			NameOfX:       "qwerty123",
			NameOfY:       "",
			plot:          Plot{},
			err:           errors.New("NameOfX is too long"),
		},
		{
			TotalHeight:   200,
			TotalWidth:    200,
			XDivisionsQty: 12,
			YDivisionsQty: 12,
			NameOfX:       "",
			NameOfY:       "qwerty123",
			plot:          Plot{},
			err:           errors.New("NameOfY is too long"),
		},
	}
	for key, val := range testArray {
		doTest(val, key)
	}
}
func TestPlot_AddAngular(t *testing.T) {
	type TestCase struct {
		plotIn, plotOut Plot
		X, Y            []float64
		err             error
	}

	doTest := func(c TestCase, testId int) {
		t.Logf("Test %d start", testId)
		err := c.plotIn.AddAngular(c.X, c.Y)

		if fmt.Sprint(err) != fmt.Sprint(c.err) {
			t.Errorf("%sTest %d failed%s\tgot: %s, want: %s\n", red, testId, normal, c.err, err)
		} else if !reflect.DeepEqual(c.plotOut, c.plotIn) {
			t.Errorf("%sTest %d failed%s\twant: %s, got: %s\n", red, testId, normal, fmt.Sprint(c.plotOut), fmt.Sprint(c.plotIn))
		} else {
			t.Logf("%sTest %d success%s\t%s\n", green, testId, normal, "")
		}
	}
	testArray := []TestCase{
		{
			plotIn: Plot{
				style: Style{
					TotalHeight:   200,
					TotalWidth:    200,
					XDivisionsQty: 12,
					YDivisionsQty: 12,
					NameOfX:       "",
					NameOfY:       "",
				},
			},
			X: []float64{1, 2, 3, 4, 5},
			Y: []float64{6, 7, 8, 9, 0},
			plotOut: Plot{
				style: Style{
					TotalHeight:   200,
					TotalWidth:    200,
					XDivisionsQty: 12,
					YDivisionsQty: 12,
					NameOfX:       "",
					NameOfY:       "",
				},
				angular: []pointArr{
					{
						X: []float64{1, 2, 3, 4, 5},
						Y: []float64{6, 7, 8, 9, 0},
					},
				},
			},
			err: nil,
		},
		{
			plotIn: Plot{
				style: Style{
					TotalHeight:   200,
					TotalWidth:    200,
					XDivisionsQty: 12,
					YDivisionsQty: 12,
					NameOfX:       "",
					NameOfY:       "",
				},
				angular: []pointArr{
					{
						X: []float64{1, 2, 3, 4, 5},
						Y: []float64{6, 7, 8, 9, 0},
					},
				},
			},
			X: []float64{1, 2, 3, 4, 5},
			Y: []float64{6, 7, 8, 9, 4},
			plotOut: Plot{
				style: Style{
					TotalHeight:   200,
					TotalWidth:    200,
					XDivisionsQty: 12,
					YDivisionsQty: 12,
					NameOfX:       "",
					NameOfY:       "",
				},
				angular: []pointArr{
					{
						X: []float64{1, 2, 3, 4, 5},
						Y: []float64{6, 7, 8, 9, 0},
					},
					{
						X: []float64{1, 2, 3, 4, 5},
						Y: []float64{6, 7, 8, 9, 4},
					},
				},
			},
			err: nil,
		},
		{
			plotIn: Plot{
				style: Style{
					TotalHeight:   200,
					TotalWidth:    200,
					XDivisionsQty: 12,
					YDivisionsQty: 12,
					NameOfX:       "",
					NameOfY:       "",
				},
			},
			X: []float64{1, 3, 4, 5},
			Y: []float64{6, 7, 8, 9, 0},
			plotOut: Plot{
				style: Style{
					TotalHeight:   200,
					TotalWidth:    200,
					XDivisionsQty: 12,
					YDivisionsQty: 12,
					NameOfX:       "",
					NameOfY:       "",
				},
			},
			err: errors.New("len(X) ≠ len(Y)"),
		},
		{
			plotIn: Plot{
				style: Style{
					TotalHeight:   200,
					TotalWidth:    200,
					XDivisionsQty: 12,
					YDivisionsQty: 12,
					NameOfX:       "",
					NameOfY:       "",
				},
			},
			X: []float64{},
			Y: []float64{},
			plotOut: Plot{
				style: Style{
					TotalHeight:   200,
					TotalWidth:    200,
					XDivisionsQty: 12,
					YDivisionsQty: 12,
					NameOfX:       "",
					NameOfY:       "",
				},
			},
			err: nil,
		},
	}
	for key, val := range testArray {
		doTest(val, key)
	}
}
func TestPlot_AddPoint(t *testing.T) {
	type TestCase struct {
		plotIn, plotOut Plot
		X, Y            []float64
		err             error
	}

	doTest := func(c TestCase, testId int) {
		t.Logf("Test %d start", testId)
		err := c.plotIn.AddPoint(c.X, c.Y)

		if fmt.Sprint(err) != fmt.Sprint(c.err) {
			t.Errorf("%sTest %d failed%s\tgot: %s, want: %s\n", red, testId, normal, c.err, err)
		} else if !reflect.DeepEqual(c.plotOut, c.plotIn) {
			t.Errorf("%sTest %d failed%s\twant: %s, got: %s\n", red, testId, normal, fmt.Sprint(c.plotOut), fmt.Sprint(c.plotIn))
		} else {
			t.Logf("%sTest %d success%s\t%s\n", green, testId, normal, "")
		}
	}
	testArray := []TestCase{
		{
			plotIn: Plot{
				style: Style{
					TotalHeight:   200,
					TotalWidth:    200,
					XDivisionsQty: 12,
					YDivisionsQty: 12,
					NameOfX:       "",
					NameOfY:       "",
				},
			},
			X: []float64{1, 2, 3, 4, 5},
			Y: []float64{6, 7, 8, 9, 0},
			plotOut: Plot{
				style: Style{
					TotalHeight:   200,
					TotalWidth:    200,
					XDivisionsQty: 12,
					YDivisionsQty: 12,
					NameOfX:       "",
					NameOfY:       "",
				},
				point: []pointArr{
					{
						X: []float64{1, 2, 3, 4, 5},
						Y: []float64{6, 7, 8, 9, 0},
					},
				},
			},
			err: nil,
		},
		{
			plotIn: Plot{
				style: Style{
					TotalHeight:   200,
					TotalWidth:    200,
					XDivisionsQty: 12,
					YDivisionsQty: 12,
					NameOfX:       "",
					NameOfY:       "",
				},
				point: []pointArr{
					{
						X: []float64{1, 2, 3, 4, 5},
						Y: []float64{6, 7, 8, 9, 0},
					},
				},
			},
			X: []float64{1, 2, 3, 4, 5},
			Y: []float64{6, 7, 8, 9, 4},
			plotOut: Plot{
				style: Style{
					TotalHeight:   200,
					TotalWidth:    200,
					XDivisionsQty: 12,
					YDivisionsQty: 12,
					NameOfX:       "",
					NameOfY:       "",
				},
				point: []pointArr{
					{
						X: []float64{1, 2, 3, 4, 5},
						Y: []float64{6, 7, 8, 9, 0},
					},
					{
						X: []float64{1, 2, 3, 4, 5},
						Y: []float64{6, 7, 8, 9, 4},
					},
				},
			},
			err: nil,
		},
		{
			plotIn: Plot{
				style: Style{
					TotalHeight:   200,
					TotalWidth:    200,
					XDivisionsQty: 12,
					YDivisionsQty: 12,
					NameOfX:       "",
					NameOfY:       "",
				},
			},
			X: []float64{1, 3, 4, 5},
			Y: []float64{6, 7, 8, 9, 0},
			plotOut: Plot{
				style: Style{
					TotalHeight:   200,
					TotalWidth:    200,
					XDivisionsQty: 12,
					YDivisionsQty: 12,
					NameOfX:       "",
					NameOfY:       "",
				},
			},
			err: errors.New("len(X) ≠ len(Y)"),
		},
		{
			plotIn: Plot{
				style: Style{
					TotalHeight:   200,
					TotalWidth:    200,
					XDivisionsQty: 12,
					YDivisionsQty: 12,
					NameOfX:       "",
					NameOfY:       "",
				},
			},
			X: []float64{},
			Y: []float64{},
			plotOut: Plot{
				style: Style{
					TotalHeight:   200,
					TotalWidth:    200,
					XDivisionsQty: 12,
					YDivisionsQty: 12,
					NameOfX:       "",
					NameOfY:       "",
				},
			},
			err: nil,
		},
	}
	for key, val := range testArray {
		doTest(val, key)
	}
}
func TestPlot_AddSmooth(t *testing.T) {
	type TestCase struct {
		plotIn, plotOut Plot
		X, Y            []float64
		err             error
	}

	doTest := func(c TestCase, testId int) {
		t.Logf("Test %d start", testId)
		err := c.plotIn.AddSmooth(c.X, c.Y)

		if fmt.Sprint(err) != fmt.Sprint(c.err) {
			t.Errorf("%sTest %d failed%s\tgot: %s, want: %s\n", red, testId, normal, c.err, err)
		} else if !reflect.DeepEqual(c.plotOut, c.plotIn) {
			t.Errorf("%sTest %d failed%s\twant: %s, got: %s\n", red, testId, normal, fmt.Sprint(c.plotOut), fmt.Sprint(c.plotIn))
		} else {
			t.Logf("%sTest %d success%s\t%s\n", green, testId, normal, "")
		}
	}
	testArray := []TestCase{
		{
			plotIn: Plot{
				style: Style{
					TotalHeight:   200,
					TotalWidth:    200,
					XDivisionsQty: 12,
					YDivisionsQty: 12,
					NameOfX:       "",
					NameOfY:       "",
				},
			},
			X: []float64{1, 2, 3, 4, 5},
			Y: []float64{6, 7, 8, 9, 0},
			plotOut: Plot{
				style: Style{
					TotalHeight:   200,
					TotalWidth:    200,
					XDivisionsQty: 12,
					YDivisionsQty: 12,
					NameOfX:       "",
					NameOfY:       "",
				},
				smooth: []pointArr{
					{
						X: []float64{1, 2, 3, 4, 5},
						Y: []float64{6, 7, 8, 9, 0},
					},
				},
			},
			err: nil,
		},
		{
			plotIn: Plot{
				style: Style{
					TotalHeight:   200,
					TotalWidth:    200,
					XDivisionsQty: 12,
					YDivisionsQty: 12,
					NameOfX:       "",
					NameOfY:       "",
				},
				smooth: []pointArr{
					{
						X: []float64{1, 2, 3, 4, 5},
						Y: []float64{6, 7, 8, 9, 0},
					},
				},
			},
			X: []float64{1, 2, 3, 4, 5},
			Y: []float64{6, 7, 8, 9, 4},
			plotOut: Plot{
				style: Style{
					TotalHeight:   200,
					TotalWidth:    200,
					XDivisionsQty: 12,
					YDivisionsQty: 12,
					NameOfX:       "",
					NameOfY:       "",
				},
				smooth: []pointArr{
					{
						X: []float64{1, 2, 3, 4, 5},
						Y: []float64{6, 7, 8, 9, 0},
					},
					{
						X: []float64{1, 2, 3, 4, 5},
						Y: []float64{6, 7, 8, 9, 4},
					},
				},
			},
			err: nil,
		},
		{
			plotIn: Plot{
				style: Style{
					TotalHeight:   200,
					TotalWidth:    200,
					XDivisionsQty: 12,
					YDivisionsQty: 12,
					NameOfX:       "",
					NameOfY:       "",
				},
			},
			X: []float64{1, 3, 4, 5},
			Y: []float64{6, 7, 8, 9, 0},
			plotOut: Plot{
				style: Style{
					TotalHeight:   200,
					TotalWidth:    200,
					XDivisionsQty: 12,
					YDivisionsQty: 12,
					NameOfX:       "",
					NameOfY:       "",
				},
			},
			err: errors.New("len(X) ≠ len(Y)"),
		},
		{
			plotIn: Plot{
				style: Style{
					TotalHeight:   200,
					TotalWidth:    200,
					XDivisionsQty: 12,
					YDivisionsQty: 12,
					NameOfX:       "",
					NameOfY:       "",
				},
			},
			X: []float64{},
			Y: []float64{},
			plotOut: Plot{
				style: Style{
					TotalHeight:   200,
					TotalWidth:    200,
					XDivisionsQty: 12,
					YDivisionsQty: 12,
					NameOfX:       "",
					NameOfY:       "",
				},
			},
			err: nil,
		},
	}
	for key, val := range testArray {
		doTest(val, key)
	}
}
func TestPlot_Draw(t *testing.T) {
	type TestCase struct {
		dataIn Plot
		picOut string
		err    error
	}

	doTest := func(c TestCase, testId int) {
		t.Logf("Test %d start", testId)
		pic, err := c.dataIn.Draw()

		if fmt.Sprint(err) != fmt.Sprint(c.err) {
			t.Errorf("%sTest %d failed%s\tgot: %s, want: %s\n", red, testId, normal, c.err, err)
		} else if pic != c.picOut {
			t.Errorf("%sTest %d failed%s\twant: %s, got: %s\n", red, testId, normal, c.picOut, pic)
			_ = os.WriteFile(fmt.Sprintf("TestPlot_Draw%d.svg", testId), []byte(pic), 777)
		} else {
			t.Logf("%sTest %d success%s\t%s\n", green, testId, normal, "")
		}
	}
	testArray := []TestCase{
		{
			dataIn: Plot{
				style: Style{
					TotalHeight:   200,
					TotalWidth:    200,
					XDivisionsQty: 12,
					YDivisionsQty: 12,
					NameOfX:       "x",
					NameOfY:       "y",
				},
				smooth: []pointArr{
					{
						X: []float64{-3, -2, -1, 0, 1, 2, 3},
						Y: []float64{9, 4, 1, 0, 1, 4, 9},
					},
					{
						X: []float64{-3, -2, -1, 0, 1, 2, 3},
						Y: []float64{10, 5, 2, 1, 2, 5, 10},
					},
				},
				angular: []pointArr{
					{
						X: []float64{0, 2, 2},
						Y: []float64{4, 4, 0},
					},
				},
				point: []pointArr{
					{
						X: []float64{2},
						Y: []float64{4},
					},
					{
						X: []float64{-3, -2, -1, 0, 1, 2, 3},
						Y: []float64{10, 5, 2, 1, 2, 5, 10},
					},
				},
			},
			picOut: TestPlot_Draw0,
			err:    nil,
		},
		{
			dataIn: Plot{
				style: Style{
					TotalHeight:   200,
					TotalWidth:    200,
					XDivisionsQty: 12,
					YDivisionsQty: 12,
					NameOfX:       "",
					NameOfY:       "",
				},
				smooth: []pointArr{
					{
						X: []float64{-3, -2, -1, 0, 1, 2, 3},
						Y: []float64{9, 4, 1, 0, 1, 4, 9},
					},
					{
						X: []float64{-3, -2, -1, 0, 1, 2, 3},
						Y: []float64{10, 5, 2, 1, 2, 5, 10},
					},
				},
				angular: []pointArr{
					{
						X: []float64{0, 2, 2},
						Y: []float64{4, 4, 0},
					},
				},
				point: []pointArr{
					{
						X: []float64{2},
						Y: []float64{4},
					},
					{
						X: []float64{-3, -2, -1, 0, 1, 2, 3},
						Y: []float64{10, 5, 2, 1, 2, 5, 10},
					},
				},
			},
			picOut: TestPlot_Draw1,
			err:    nil,
		},
		{
			dataIn: Plot{
				style: Style{
					TotalHeight:   200,
					TotalWidth:    200,
					XDivisionsQty: 12,
					YDivisionsQty: 12,
					NameOfX:       "qwerty123",
					NameOfY:       "",
				},
				smooth: []pointArr{
					{
						X: []float64{-3, -2, -1, 0, 1, 2, 3},
						Y: []float64{9, 4, 1, 0, 1, 4, 9},
					},
				},
				angular: []pointArr{
					{
						X: []float64{0, 2, 2},
						Y: []float64{4, 4, 0},
					},
				},
				point: []pointArr{
					{
						X: []float64{2},
						Y: []float64{4},
					},
				},
			},
			err: errors.New("x name max len is 6"),
		},
		{
			dataIn: Plot{
				style: Style{
					TotalHeight:   200,
					TotalWidth:    200,
					XDivisionsQty: 12,
					YDivisionsQty: 12,
					NameOfX:       "",
					NameOfY:       "qwerty123",
				},
				smooth: []pointArr{
					{
						X: []float64{-3, -2, -1, 0, 1, 2, 3},
						Y: []float64{9, 4, 1, 0, 1, 4, 9},
					},
				},
				angular: []pointArr{
					{
						X: []float64{0, 2, 2},
						Y: []float64{4, 4, 0},
					},
				},
				point: []pointArr{
					{
						X: []float64{2},
						Y: []float64{4},
					},
				},
			},
			err: errors.New("y name max len is 6"),
		},
		{
			dataIn: Plot{
				style: Style{
					TotalHeight:   100,
					TotalWidth:    200,
					XDivisionsQty: 12,
					YDivisionsQty: 12,
					NameOfX:       "",
					NameOfY:       "",
				},
				smooth: []pointArr{
					{
						X: []float64{-3, -2, -1, 0, 1, 2, 3},
						Y: []float64{9, 4, 1, 0, 1, 4, 9},
					},
				},
				angular: []pointArr{
					{
						X: []float64{0, 2, 2},
						Y: []float64{4, 4, 0},
					},
				},
				point: []pointArr{
					{
						X: []float64{2},
						Y: []float64{4},
					},
				},
			},
			err: errors.New("height is too small"),
		},
		{
			dataIn: Plot{
				style: Style{
					TotalHeight:   200,
					TotalWidth:    100,
					XDivisionsQty: 12,
					YDivisionsQty: 12,
					NameOfX:       "",
					NameOfY:       "",
				},
				smooth: []pointArr{
					{
						X: []float64{-3, -2, -1, 0, 1, 2, 3},
						Y: []float64{9, 4, 1, 0, 1, 4, 9},
					},
				},
				angular: []pointArr{
					{
						X: []float64{0, 2, 2},
						Y: []float64{4, 4, 0},
					},
				},
				point: []pointArr{
					{
						X: []float64{2},
						Y: []float64{4},
					},
				},
			},
			err: errors.New("width is too small"),
		},
		{
			dataIn: Plot{
				style: Style{
					TotalHeight:   200,
					TotalWidth:    200,
					XDivisionsQty: 12,
					YDivisionsQty: 12,
					NameOfX:       "",
					NameOfY:       "",
				},
				angular: []pointArr{
					{
						X: []float64{1},
						Y: []float64{},
					},
				},
			},
			err: errors.New("empty input"),
		},
		{
			dataIn: Plot{
				style: Style{
					TotalHeight:   200,
					TotalWidth:    200,
					XDivisionsQty: 12,
					YDivisionsQty: 12,
					NameOfX:       "",
					NameOfY:       "",
				},
				angular: []pointArr{
					{
						X: []float64{},
						Y: []float64{1},
					},
				},
			},
			err: errors.New("empty input"),
		},
		{
			dataIn: Plot{
				style: Style{
					TotalHeight:   200,
					TotalWidth:    200,
					XDivisionsQty: 12,
					YDivisionsQty: 12,
					NameOfX:       "",
					NameOfY:       "",
				},
				smooth: []pointArr{
					{
						X: []float64{1},
						Y: []float64{},
					},
				},
			},
			err: errors.New("empty input"),
		},
		{
			dataIn: Plot{
				style: Style{
					TotalHeight:   200,
					TotalWidth:    200,
					XDivisionsQty: 12,
					YDivisionsQty: 12,
					NameOfX:       "",
					NameOfY:       "",
				},
				smooth: []pointArr{
					{
						X: []float64{},
						Y: []float64{1},
					},
				},
			},
			err: errors.New("empty input"),
		},
		{
			dataIn: Plot{
				style: Style{
					TotalHeight:   200,
					TotalWidth:    200,
					XDivisionsQty: 12,
					YDivisionsQty: 12,
					NameOfX:       "",
					NameOfY:       "",
				},
				point: []pointArr{
					{
						X: []float64{1},
						Y: []float64{},
					},
				},
			},
			err: errors.New("empty input"),
		},
		{
			dataIn: Plot{
				style: Style{
					TotalHeight:   200,
					TotalWidth:    200,
					XDivisionsQty: 12,
					YDivisionsQty: 12,
					NameOfX:       "",
					NameOfY:       "",
				},
				point: []pointArr{
					{
						X: []float64{},
						Y: []float64{1},
					},
				},
			},
			err: errors.New("empty input"),
		},
		{
			dataIn: Plot{
				style: Style{
					TotalHeight:   200,
					TotalWidth:    200,
					XDivisionsQty: 12,
					YDivisionsQty: 12,
					NameOfX:       "",
					NameOfY:       "",
				},
				point: []pointArr{
					{
						X: []float64{2, 2, 5, 8},
						Y: []float64{1, 2, 2},
					},
				},
			},
			err: errors.New("len(x)!=len(y)"),
		},
		{
			dataIn: Plot{
				style: Style{
					TotalHeight:   200,
					TotalWidth:    200,
					XDivisionsQty: 12,
					YDivisionsQty: 12,
					NameOfX:       "",
					NameOfY:       "",
				},
				smooth: []pointArr{
					{
						X: []float64{2, 2, 5, 8},
						Y: []float64{1, 2, 2},
					},
				},
			},
			err: errors.New("len(x)!=len(y)"),
		},
		{
			dataIn: Plot{
				style: Style{
					TotalHeight:   200,
					TotalWidth:    200,
					XDivisionsQty: 12,
					YDivisionsQty: 12,
					NameOfX:       "",
					NameOfY:       "",
				},
				angular: []pointArr{
					{
						X: []float64{2, 2, 5, 8},
						Y: []float64{1, 2, 2},
					},
				},
			},
			err: errors.New("len(x)!=len(y)"),
		},
		{
			dataIn: Plot{
				style: Style{
					TotalHeight:   200,
					TotalWidth:    200,
					XDivisionsQty: 12,
					YDivisionsQty: 12,
					NameOfX:       "",
					NameOfY:       "",
				},
				angular: []pointArr{
					{
						X: []float64{2, 2, 5, 8},
						Y: []float64{0, 0, 0},
					},
				},
			},
			err: errors.New("min>=max"),
		},
		{
			dataIn: Plot{
				style: Style{
					TotalHeight:   200,
					TotalWidth:    200,
					XDivisionsQty: 12,
					YDivisionsQty: 12,
					NameOfX:       "",
					NameOfY:       "",
				},
				angular: []pointArr{
					{
						X: []float64{0, 0, 0},
						Y: []float64{12, 3, 4},
					},
				},
			},
			err: errors.New("min>=max"),
		},
	}
	for key, val := range testArray {
		doTest(val, key)
	}
}
func TestPlot_DrawFromZero(t *testing.T) {
	type TestCase struct {
		dataIn Plot
		picOut string
		err    error
	}

	doTest := func(c TestCase, testId int) {
		t.Logf("Test %d start", testId)
		pic, err := c.dataIn.DrawFromZero()

		if fmt.Sprint(err) != fmt.Sprint(c.err) {
			t.Errorf("%sTest %d failed%s\tgot: %s, want: %s\n", red, testId, normal, c.err, err)
		} else if pic != c.picOut {
			t.Errorf("%sTest %d failed%s\twant: %s, got: %s\n", red, testId, normal, c.picOut, pic)
			_ = os.WriteFile(fmt.Sprintf("TestPlot_DrawFromZero%d.svg", testId), []byte(pic), 777)
		} else {
			t.Logf("%sTest %d success%s\t%s\n", green, testId, normal, "")
		}
	}
	testArray := []TestCase{
		{
			dataIn: Plot{
				style: Style{
					TotalHeight:   300,
					TotalWidth:    300,
					XDivisionsQty: 5,
					YDivisionsQty: 12,
					NameOfX:       "x",
					NameOfY:       "y",
				},
				smooth: []pointArr{
					{
						X: []float64{0, 1, 2, 3},
						Y: []float64{0, 1, 4, 9},
					},
					{
						X: []float64{0, 1, 2, 3},
						Y: []float64{1, 2, 5, 10},
					},
				},
				angular: []pointArr{
					{
						X: []float64{0, 2, 2},
						Y: []float64{4, 4, 0},
					},
				},
				point: []pointArr{
					{
						X: []float64{2},
						Y: []float64{4},
					},
					{
						X: []float64{0, 1, 2, 3},
						Y: []float64{1, 2, 5, 10},
					},
				},
			},
			picOut: TestPlot_DrawFromZero0,
			err:    nil,
		},
		{
			dataIn: Plot{
				style: Style{
					TotalHeight:   300,
					TotalWidth:    300,
					XDivisionsQty: 5,
					YDivisionsQty: 12,
					NameOfX:       "",
					NameOfY:       "",
				},
				smooth: []pointArr{
					{
						X: []float64{0, 1, 2, 3},
						Y: []float64{0, 1, 4, 9},
					},
					{
						X: []float64{0, 1, 2, 3},
						Y: []float64{1, 2, 5, 10},
					},
				},
				angular: []pointArr{
					{
						X: []float64{0, 2, 2},
						Y: []float64{4, 4, 0},
					},
				},
				point: []pointArr{
					{
						X: []float64{2},
						Y: []float64{4},
					},
					{
						X: []float64{0, 1, 2, 3},
						Y: []float64{1, 2, 5, 10},
					},
				},
			},
			picOut: TestPlot_DrawFromZero1,
			err:    nil,
		},
		{
			dataIn: Plot{
				style: Style{
					TotalHeight:   200,
					TotalWidth:    200,
					XDivisionsQty: 12,
					YDivisionsQty: 12,
					NameOfX:       "qwerty123",
					NameOfY:       "",
				},
				smooth: []pointArr{
					{
						X: []float64{0, 1, 2, 3},
						Y: []float64{0, 1, 4, 9},
					},
				},
				angular: []pointArr{
					{
						X: []float64{0, 2, 2},
						Y: []float64{4, 4, 0},
					},
				},
				point: []pointArr{
					{
						X: []float64{2},
						Y: []float64{4},
					},
				},
			},
			err: errors.New("x name max len is 6"),
		},
		{
			dataIn: Plot{
				style: Style{
					TotalHeight:   200,
					TotalWidth:    200,
					XDivisionsQty: 12,
					YDivisionsQty: 12,
					NameOfX:       "",
					NameOfY:       "qwerty123",
				},
				smooth: []pointArr{
					{
						X: []float64{0, 1, 2, 3},
						Y: []float64{0, 1, 4, 9},
					},
				},
				angular: []pointArr{
					{
						X: []float64{0, 2, 2},
						Y: []float64{4, 4, 0},
					},
				},
				point: []pointArr{
					{
						X: []float64{2},
						Y: []float64{4},
					},
				},
			},
			err: errors.New("y name max len is 6"),
		},
		{
			dataIn: Plot{
				style: Style{
					TotalHeight:   100,
					TotalWidth:    200,
					XDivisionsQty: 12,
					YDivisionsQty: 12,
					NameOfX:       "",
					NameOfY:       "",
				},
				smooth: []pointArr{
					{
						X: []float64{0, 1, 2, 3},
						Y: []float64{0, 1, 4, 9},
					},
				},
				angular: []pointArr{
					{
						X: []float64{0, 2, 2},
						Y: []float64{4, 4, 0},
					},
				},
				point: []pointArr{
					{
						X: []float64{2},
						Y: []float64{4},
					},
				},
			},
			err: errors.New("height is too small"),
		},
		{
			dataIn: Plot{
				style: Style{
					TotalHeight:   200,
					TotalWidth:    100,
					XDivisionsQty: 12,
					YDivisionsQty: 12,
					NameOfX:       "",
					NameOfY:       "",
				},
				smooth: []pointArr{
					{
						X: []float64{0, 1, 2, 3},
						Y: []float64{0, 1, 4, 9},
					},
				},
				angular: []pointArr{
					{
						X: []float64{0, 2, 2},
						Y: []float64{4, 4, 0},
					},
				},
				point: []pointArr{
					{
						X: []float64{2},
						Y: []float64{4},
					},
				},
			},
			err: errors.New("width is too small"),
		},
		{
			dataIn: Plot{
				style: Style{
					TotalHeight:   200,
					TotalWidth:    200,
					XDivisionsQty: 12,
					YDivisionsQty: 12,
					NameOfX:       "",
					NameOfY:       "",
				},
				angular: []pointArr{
					{
						X: []float64{1},
						Y: []float64{},
					},
				},
			},
			err: errors.New("empty input"),
		},
		{
			dataIn: Plot{
				style: Style{
					TotalHeight:   200,
					TotalWidth:    200,
					XDivisionsQty: 12,
					YDivisionsQty: 12,
					NameOfX:       "",
					NameOfY:       "",
				},
				angular: []pointArr{
					{
						X: []float64{},
						Y: []float64{1},
					},
				},
			},
			err: errors.New("empty input"),
		},
		{
			dataIn: Plot{
				style: Style{
					TotalHeight:   200,
					TotalWidth:    200,
					XDivisionsQty: 12,
					YDivisionsQty: 12,
					NameOfX:       "",
					NameOfY:       "",
				},
				smooth: []pointArr{
					{
						X: []float64{1},
						Y: []float64{},
					},
				},
			},
			err: errors.New("empty input"),
		},
		{
			dataIn: Plot{
				style: Style{
					TotalHeight:   200,
					TotalWidth:    200,
					XDivisionsQty: 12,
					YDivisionsQty: 12,
					NameOfX:       "",
					NameOfY:       "",
				},
				smooth: []pointArr{
					{
						X: []float64{},
						Y: []float64{1},
					},
				},
			},
			err: errors.New("empty input"),
		},
		{
			dataIn: Plot{
				style: Style{
					TotalHeight:   200,
					TotalWidth:    200,
					XDivisionsQty: 12,
					YDivisionsQty: 12,
					NameOfX:       "",
					NameOfY:       "",
				},
				point: []pointArr{
					{
						X: []float64{1},
						Y: []float64{},
					},
				},
			},
			err: errors.New("empty input"),
		},
		{
			dataIn: Plot{
				style: Style{
					TotalHeight:   200,
					TotalWidth:    200,
					XDivisionsQty: 12,
					YDivisionsQty: 12,
					NameOfX:       "",
					NameOfY:       "",
				},
				point: []pointArr{
					{
						X: []float64{},
						Y: []float64{1},
					},
				},
			},
			err: errors.New("empty input"),
		},
		{
			dataIn: Plot{
				style: Style{
					TotalHeight:   200,
					TotalWidth:    200,
					XDivisionsQty: 12,
					YDivisionsQty: 12,
					NameOfX:       "",
					NameOfY:       "",
				},
				point: []pointArr{
					{
						X: []float64{2, 2, 5, 8},
						Y: []float64{1, 2, 2},
					},
				},
			},
			err: errors.New("len(x)!=len(y)"),
		},
		{
			dataIn: Plot{
				style: Style{
					TotalHeight:   200,
					TotalWidth:    200,
					XDivisionsQty: 12,
					YDivisionsQty: 12,
					NameOfX:       "",
					NameOfY:       "",
				},
				smooth: []pointArr{
					{
						X: []float64{2, 2, 5, 8},
						Y: []float64{1, 2, 2},
					},
				},
			},
			err: errors.New("len(x)!=len(y)"),
		},
		{
			dataIn: Plot{
				style: Style{
					TotalHeight:   200,
					TotalWidth:    200,
					XDivisionsQty: 12,
					YDivisionsQty: 12,
					NameOfX:       "",
					NameOfY:       "",
				},
				angular: []pointArr{
					{
						X: []float64{2, 2, 5, 8},
						Y: []float64{1, 2, 2},
					},
				},
			},
			err: errors.New("len(x)!=len(y)"),
		},
		{
			dataIn: Plot{
				style: Style{
					TotalHeight:   200,
					TotalWidth:    200,
					XDivisionsQty: 12,
					YDivisionsQty: 12,
					NameOfX:       "",
					NameOfY:       "",
				},
				angular: []pointArr{
					{
						X: []float64{2, 2, 5, 8},
						Y: []float64{0, 0, 0},
					},
				},
			},
			err: errors.New("min>=max"),
		},
		{
			dataIn: Plot{
				style: Style{
					TotalHeight:   200,
					TotalWidth:    200,
					XDivisionsQty: 12,
					YDivisionsQty: 12,
					NameOfX:       "",
					NameOfY:       "",
				},
				angular: []pointArr{
					{
						X: []float64{0, 0, 0},
						Y: []float64{12, 3, 4},
					},
				},
			},
			err: errors.New("min>=max"),
		},
		{
			dataIn: Plot{
				style: Style{
					TotalHeight:   200,
					TotalWidth:    200,
					XDivisionsQty: 12,
					YDivisionsQty: 12,
					NameOfX:       "",
					NameOfY:       "",
				},
				angular: []pointArr{
					{
						X: []float64{2, 2, 5, 8},
						Y: []float64{0, -5, 0},
					},
				},
			},
			err: errors.New("y values should be higher then zero"),
		},
		{
			dataIn: Plot{
				style: Style{
					TotalHeight:   200,
					TotalWidth:    200,
					XDivisionsQty: 12,
					YDivisionsQty: 12,
					NameOfX:       "",
					NameOfY:       "",
				},
				angular: []pointArr{
					{
						X: []float64{0, -1, 0},
						Y: []float64{12, 3, 4},
					},
				},
			},
			err: errors.New("x values should be higher then zero"),
		},
	}
	for key, val := range testArray {
		doTest(val, key)
	}
}

//	func TestPlot_Draw2(t *testing.T) {
//		ex1()
//		ex2()
//		ex3()
//		ex4()
//		ex5()
//		ex6()
//		ex7()
//	}
func ex1() {

	TotalHeight := uint(200) // px
	TotalWidth := uint(400)  // px
	xDivisionsQty := uint(10)
	yDivisionsQty := uint(10)
	NameOfX := string("x")
	NameOfY := string("y")
	xArray := []float64{-3, -2, -1, 0, 1, 2, 3}
	yArray := []float64{9, 4, 1, 0, 1, 4, 9}
	plot, _ := DrawSmooth(TotalHeight, TotalWidth, xDivisionsQty, yDivisionsQty, xArray, yArray, NameOfX, NameOfY)
	_ = os.WriteFile("smooth.svg", []byte(plot), 777)

}
func ex2() {
	TotalHeight := uint(200) // px
	TotalWidth := uint(400)  // px
	xDivisionsQty := uint(10)
	yDivisionsQty := uint(10)
	NameOfX := string("x")
	NameOfY := string("y")
	xArray := []float64{1, 2, 3, 4, 5, 6, 7}
	yArray := []float64{9, 4, 1, 0, 1, 4, 9}
	plot, _ := DrawSmoothFromZero(TotalHeight, TotalWidth, xDivisionsQty, yDivisionsQty, xArray, yArray, NameOfX, NameOfY)
	_ = os.WriteFile("smooth_from_zero.svg", []byte(plot), 777)

}
func ex3() {
	TotalHeight := uint(200) // px
	TotalWidth := uint(400)  // px
	xDivisionsQty := uint(10)
	yDivisionsQty := uint(10)
	NameOfX := string("x")
	NameOfY := string("y")
	xArray := []float64{-3, -2, -1, 0, 1, 2, 3}
	yArray := []float64{9, 4, 1, 0, 1, 4, 9}
	plot, _ := DrawPoint(TotalHeight, TotalWidth, xDivisionsQty, yDivisionsQty, xArray, yArray, NameOfX, NameOfY)
	_ = os.WriteFile("point.svg", []byte(plot), 777)
}
func ex4() {
	TotalHeight := uint(200) // px
	TotalWidth := uint(400)  // px
	xDivisionsQty := uint(10)
	yDivisionsQty := uint(10)
	NameOfX := string("x")
	NameOfY := string("y")
	xArray := []float64{1, 2, 3, 4, 5, 6, 7}
	yArray := []float64{9, 4, 1, 0, 1, 4, 9}
	plot, _ := DrawPointFromZero(TotalHeight, TotalWidth, xDivisionsQty, yDivisionsQty, xArray, yArray, NameOfX, NameOfY)
	_ = os.WriteFile("point_from_zero.svg", []byte(plot), 777)
}
func ex5() {
	TotalHeight := uint(200) // px
	TotalWidth := uint(400)  // px
	xDivisionsQty := uint(10)
	yDivisionsQty := uint(10)
	NameOfX := string("x")
	NameOfY := string("y")
	xArray := []float64{-3, -2, -1, 0, 1, 2, 3}
	yArray := []float64{9, 4, 1, 0, 1, 4, 9}
	plot, _ := DrawAngular(TotalHeight, TotalWidth, xDivisionsQty, yDivisionsQty, xArray, yArray, NameOfX, NameOfY)
	_ = os.WriteFile("angular.svg", []byte(plot), 777)
}
func ex6() {
	TotalHeight := uint(200) // px
	TotalWidth := uint(400)  // px
	xDivisionsQty := uint(10)
	yDivisionsQty := uint(10)
	NameOfX := string("x")
	NameOfY := string("y")
	xArray := []float64{1, 2, 3, 4, 5, 6, 7}
	yArray := []float64{9, 4, 1, 0, 1, 4, 9}
	plot, _ := DrawAngularFromZero(TotalHeight, TotalWidth, xDivisionsQty, yDivisionsQty, xArray, yArray, NameOfX, NameOfY)
	_ = os.WriteFile("angular_from_zero.svg", []byte(plot), 777)
}
func ex7() {
	// Creating the new plot 200x400 px 10 divisions on each side
	plot, _ := NewPlot(200, 400, 10, 10, "x", "y")
	// Add smooth parabola plot
	_ = plot.AddSmooth([]float64{1, 2, 3, 4, 5, 6, 7}, []float64{10, 5, 2, 1, 2, 5, 10})
	// Add a point to the bottom of the parabola
	_ = plot.AddPoint([]float64{4}, []float64{1})
	// Add some line
	_ = plot.AddAngular([]float64{1, 1, 7, 7}, []float64{10, 1, 1, 10})
	// generate svg
	svg, _ := plot.Draw()
	svgFromZero, _ := plot.DrawFromZero()
	// write it to the disk
	_ = os.WriteFile("combined.svg", []byte(svg), 777)
	_ = os.WriteFile("combined_from_zero.svg", []byte(svgFromZero), 777)
}
