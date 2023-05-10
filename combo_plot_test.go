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
			_ = os.WriteFile(fmt.Sprintf("test%d.svg", testId), []byte(pic), 777)
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
			picOut: `<?xml version="1.0" encoding="UTF-8"?>
<svg width="200" height="200" viewBox="0 0 200 200" version="1.1" xmlns="http://www.w3.org/2000/svg">
"<style>
.axis {
font-family="Arial, Helvetica, sans-serif">;
font-size: 12pt;
}
</style>
<line x1="16" y1="0" x2="200" y2="0" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="9">y</text>
<line x1="16" y1="18.7" x2="200" y2="18.7" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="22.7">9</text>
<line x1="16" y1="37.4" x2="200" y2="37.4" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="41.4">8</text>
<line x1="16" y1="56.1" x2="200" y2="56.1" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="60.1">7</text>
<line x1="16" y1="74.8" x2="200" y2="74.8" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="78.8">6</text>
<line x1="16" y1="93.5" x2="200" y2="93.5" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="97.5">5</text>
<line x1="16" y1="112.2" x2="200" y2="112.2" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="116.2">4</text>
<line x1="16" y1="130.9" x2="200" y2="130.9" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="134.9">3</text>
<line x1="16" y1="149.6" x2="200" y2="149.6" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="153.6">2</text>
<line x1="16" y1="168.3" x2="200" y2="168.3" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="172.3">1</text>
<line x1="16" y1="187" x2="200" y2="187" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="191">0</text>
<line x1="16" y1="0" x2="16" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="16" y="200">-3</text>
<line x1="31.33" y1="0" x2="31.33" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="31.33" y="200">-2.5</text>
<line x1="46.67" y1="0" x2="46.67" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="46.67" y="200">-2</text>
<line x1="62" y1="0" x2="62" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="62" y="200">-1.5</text>
<line x1="77.33" y1="0" x2="77.33" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="77.33" y="200">-1</text>
<line x1="92.67" y1="0" x2="92.67" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="92.67" y="200">-0.5</text>
<line x1="108" y1="0" x2="108" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="108" y="200">0</text>
<line x1="123.33" y1="0" x2="123.33" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="123.33" y="200">0.5</text>
<line x1="138.67" y1="0" x2="138.67" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="138.67" y="200">1</text>
<line x1="154" y1="0" x2="154" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="154" y="200">1.5</text>
<line x1="169.33" y1="0" x2="169.33" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="169.33" y="200">2</text>
<line x1="184.67" y1="0" x2="184.67" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="184.67" y="200">2.5</text>
<line x1="200" y1="0" x2="200" y2="187" stroke="black" />
<text text-anchor="end" class="axis" x="200" y="200">x</text>

<path d="M 108 112.2 L 169.33333333333334 112.2 L 169.33333333333334 187" fill="none" stroke="black" stroke-width="3"/>
<circle cx="169.33" cy="112.20" r="5"/>

<circle cx="16.00" cy="0.00" r="5"/>
<circle cx="46.67" cy="93.50" r="5"/>
<circle cx="77.33" cy="149.60" r="5"/>
<circle cx="108.00" cy="168.30" r="5"/>
<circle cx="138.67" cy="149.60" r="5"/>
<circle cx="169.33" cy="93.50" r="5"/>
<circle cx="200.00" cy="0.00" r="5"/>

<path d="M 16.00 18.70 C 16.00,18.70 39.221,94.04 46.67,112.20 C 53.90,129.84 65.383,153.73 77.33,168.30 C 83.78,176.17 98.800,187.00 108.00,187.00 C 117.20,187.00 132.217,176.17 138.67,168.30 C 150.62,153.73 162.102,129.84 169.33,112.20 C 169.33,112.20 187.002,62.87 200.00,18.70 " fill="none" stroke="black" stroke-width="3" stroke-linecap="round"/>
<path d="M 16.00 0.00 C 16.00,0.00 39.221,75.34 46.67,93.50 C 53.90,111.14 65.383,135.03 77.33,149.60 C 83.78,157.47 98.800,168.30 108.00,168.30 C 117.20,168.30 132.217,157.47 138.67,149.60 C 150.62,135.03 162.102,111.14 169.33,93.50 C 169.33,93.50 187.002,44.17 200.00,0.00 " fill="none" stroke="black" stroke-width="3" stroke-linecap="round"/>

</svg>`,
			err: nil,
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
			picOut: `<?xml version="1.0" encoding="UTF-8"?>
<svg width="200" height="200" viewBox="0 0 200 200" version="1.1" xmlns="http://www.w3.org/2000/svg">
"<style>
.axis {
font-family="Arial, Helvetica, sans-serif">;
font-size: 12pt;
}
</style>
<line x1="23" y1="0" x2="200" y2="0" stroke="black" />
<text text-anchor="end" class="axis" x="19" y="9">10</text>
<line x1="23" y1="18.7" x2="200" y2="18.7" stroke="black" />
<text text-anchor="end" class="axis" x="19" y="22.7">9</text>
<line x1="23" y1="37.4" x2="200" y2="37.4" stroke="black" />
<text text-anchor="end" class="axis" x="19" y="41.4">8</text>
<line x1="23" y1="56.1" x2="200" y2="56.1" stroke="black" />
<text text-anchor="end" class="axis" x="19" y="60.1">7</text>
<line x1="23" y1="74.8" x2="200" y2="74.8" stroke="black" />
<text text-anchor="end" class="axis" x="19" y="78.8">6</text>
<line x1="23" y1="93.5" x2="200" y2="93.5" stroke="black" />
<text text-anchor="end" class="axis" x="19" y="97.5">5</text>
<line x1="23" y1="112.2" x2="200" y2="112.2" stroke="black" />
<text text-anchor="end" class="axis" x="19" y="116.2">4</text>
<line x1="23" y1="130.9" x2="200" y2="130.9" stroke="black" />
<text text-anchor="end" class="axis" x="19" y="134.9">3</text>
<line x1="23" y1="149.6" x2="200" y2="149.6" stroke="black" />
<text text-anchor="end" class="axis" x="19" y="153.6">2</text>
<line x1="23" y1="168.3" x2="200" y2="168.3" stroke="black" />
<text text-anchor="end" class="axis" x="19" y="172.3">1</text>
<line x1="23" y1="187" x2="200" y2="187" stroke="black" />
<text text-anchor="end" class="axis" x="19" y="191">0</text>
<line x1="23" y1="0" x2="23" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="23" y="200">-3</text>
<line x1="37.75" y1="0" x2="37.75" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="37.75" y="200">-2.5</text>
<line x1="52.5" y1="0" x2="52.5" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="52.5" y="200">-2</text>
<line x1="67.25" y1="0" x2="67.25" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="67.25" y="200">-1.5</text>
<line x1="82" y1="0" x2="82" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="82" y="200">-1</text>
<line x1="96.75" y1="0" x2="96.75" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="96.75" y="200">-0.5</text>
<line x1="111.5" y1="0" x2="111.5" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="111.5" y="200">0</text>
<line x1="126.25" y1="0" x2="126.25" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="126.25" y="200">0.5</text>
<line x1="141" y1="0" x2="141" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="141" y="200">1</text>
<line x1="155.75" y1="0" x2="155.75" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="155.75" y="200">1.5</text>
<line x1="170.5" y1="0" x2="170.5" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="170.5" y="200">2</text>
<line x1="185.25" y1="0" x2="185.25" y2="187" stroke="black" />
<text text-anchor="middle" class="axis" x="185.25" y="200">2.5</text>
<line x1="200" y1="0" x2="200" y2="187" stroke="black" />
<text text-anchor="end" class="axis" x="200" y="200">3</text>

<path d="M 111.5 112.2 L 170.5 112.2 L 170.5 187" fill="none" stroke="black" stroke-width="3"/>
<circle cx="170.50" cy="112.20" r="5"/>

<circle cx="23.00" cy="0.00" r="5"/>
<circle cx="52.50" cy="93.50" r="5"/>
<circle cx="82.00" cy="149.60" r="5"/>
<circle cx="111.50" cy="168.30" r="5"/>
<circle cx="141.00" cy="149.60" r="5"/>
<circle cx="170.50" cy="93.50" r="5"/>
<circle cx="200.00" cy="0.00" r="5"/>

<path d="M 23.00 18.70 C 23.00,18.70 45.324,94.00 52.50,112.20 C 59.44,129.79 70.422,153.62 82.00,168.30 C 88.12,176.06 102.650,187.00 111.50,187.00 C 120.35,187.00 134.878,176.06 141.00,168.30 C 152.58,153.62 163.565,129.79 170.50,112.20 C 170.50,112.20 187.510,62.84 200.00,18.70 " fill="none" stroke="black" stroke-width="3" stroke-linecap="round"/>
<path d="M 23.00 0.00 C 23.00,0.00 45.324,75.30 52.50,93.50 C 59.44,111.09 70.422,134.92 82.00,149.60 C 88.12,157.36 102.650,168.30 111.50,168.30 C 120.35,168.30 134.878,157.36 141.00,149.60 C 152.58,134.92 163.565,111.09 170.50,93.50 C 170.50,93.50 187.510,44.14 200.00,0.00 " fill="none" stroke="black" stroke-width="3" stroke-linecap="round"/>

</svg>`,
			err: nil,
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
			_ = os.WriteFile(fmt.Sprintf("test%d.svg", testId), []byte(pic), 777)
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
			picOut: `<?xml version="1.0" encoding="UTF-8"?>
<svg width="300" height="300" viewBox="0 0 300 300" version="1.1" xmlns="http://www.w3.org/2000/svg">
"<style>
.axis {
font-family="Arial, Helvetica, sans-serif">;
font-size: 12pt;
}
</style>
<line x1="16" y1="0" x2="300" y2="0" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="9">y</text>
<line x1="16" y1="28.7" x2="300" y2="28.7" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="32.7">9</text>
<line x1="16" y1="57.4" x2="300" y2="57.4" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="61.4">8</text>
<line x1="16" y1="86.1" x2="300" y2="86.1" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="90.1">7</text>
<line x1="16" y1="114.8" x2="300" y2="114.8" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="118.8">6</text>
<line x1="16" y1="143.5" x2="300" y2="143.5" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="147.5">5</text>
<line x1="16" y1="172.2" x2="300" y2="172.2" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="176.2">4</text>
<line x1="16" y1="200.9" x2="300" y2="200.9" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="204.9">3</text>
<line x1="16" y1="229.6" x2="300" y2="229.6" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="233.6">2</text>
<line x1="16" y1="258.3" x2="300" y2="258.3" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="262.3">1</text>
<line x1="16" y1="287" x2="300" y2="287" stroke="black" />
<text text-anchor="end" class="axis" x="12" y="291">0</text>
<line x1="16" y1="0" x2="16" y2="287" stroke="black" />
<text text-anchor="middle" class="axis" x="16" y="300">0</text>
<line x1="63.33" y1="0" x2="63.33" y2="287" stroke="black" />
<text text-anchor="middle" class="axis" x="63.33" y="300">0.5</text>
<line x1="110.67" y1="0" x2="110.67" y2="287" stroke="black" />
<text text-anchor="middle" class="axis" x="110.67" y="300">1</text>
<line x1="158" y1="0" x2="158" y2="287" stroke="black" />
<text text-anchor="middle" class="axis" x="158" y="300">1.5</text>
<line x1="205.33" y1="0" x2="205.33" y2="287" stroke="black" />
<text text-anchor="middle" class="axis" x="205.33" y="300">2</text>
<line x1="252.67" y1="0" x2="252.67" y2="287" stroke="black" />
<text text-anchor="middle" class="axis" x="252.67" y="300">2.5</text>
<line x1="300" y1="0" x2="300" y2="287" stroke="black" />
<text text-anchor="end" class="axis" x="300" y="300">x</text>

<path d="M 16 172.2 L 205.33333333333334 172.2 L 205.33333333333334 287" fill="none" stroke="black" stroke-width="3"/>
<circle cx="205.33" cy="172.20" r="5"/>

<circle cx="16.00" cy="258.30" r="5"/>
<circle cx="110.67" cy="229.60" r="5"/>
<circle cx="205.33" cy="143.50" r="5"/>
<circle cx="300.00" cy="0.00" r="5"/>

<path d="M 16.00 287.00 C 16.00,287.00 94.278,268.24 110.67,258.30 C 142.88,238.77 181.160,201.52 205.33,172.20 C 205.33,172.20 258.640,97.99 300.00,28.70 " fill="none" stroke="black" stroke-width="3" stroke-linecap="round"/>
<path d="M 16.00 258.30 C 16.00,258.30 94.278,239.54 110.67,229.60 C 142.88,210.07 181.160,172.82 205.33,143.50 C 205.33,143.50 258.640,69.29 300.00,0.00 " fill="none" stroke="black" stroke-width="3" stroke-linecap="round"/>

</svg>`,
			err: nil,
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
			picOut: `<?xml version="1.0" encoding="UTF-8"?>
<svg width="300" height="300" viewBox="0 0 300 300" version="1.1" xmlns="http://www.w3.org/2000/svg">
"<style>
.axis {
font-family="Arial, Helvetica, sans-serif">;
font-size: 12pt;
}
</style>
<line x1="23" y1="0" x2="300" y2="0" stroke="black" />
<text text-anchor="end" class="axis" x="19" y="9">10</text>
<line x1="23" y1="28.7" x2="300" y2="28.7" stroke="black" />
<text text-anchor="end" class="axis" x="19" y="32.7">9</text>
<line x1="23" y1="57.4" x2="300" y2="57.4" stroke="black" />
<text text-anchor="end" class="axis" x="19" y="61.4">8</text>
<line x1="23" y1="86.1" x2="300" y2="86.1" stroke="black" />
<text text-anchor="end" class="axis" x="19" y="90.1">7</text>
<line x1="23" y1="114.8" x2="300" y2="114.8" stroke="black" />
<text text-anchor="end" class="axis" x="19" y="118.8">6</text>
<line x1="23" y1="143.5" x2="300" y2="143.5" stroke="black" />
<text text-anchor="end" class="axis" x="19" y="147.5">5</text>
<line x1="23" y1="172.2" x2="300" y2="172.2" stroke="black" />
<text text-anchor="end" class="axis" x="19" y="176.2">4</text>
<line x1="23" y1="200.9" x2="300" y2="200.9" stroke="black" />
<text text-anchor="end" class="axis" x="19" y="204.9">3</text>
<line x1="23" y1="229.6" x2="300" y2="229.6" stroke="black" />
<text text-anchor="end" class="axis" x="19" y="233.6">2</text>
<line x1="23" y1="258.3" x2="300" y2="258.3" stroke="black" />
<text text-anchor="end" class="axis" x="19" y="262.3">1</text>
<line x1="23" y1="287" x2="300" y2="287" stroke="black" />
<text text-anchor="end" class="axis" x="19" y="291">0</text>
<line x1="23" y1="0" x2="23" y2="287" stroke="black" />
<text text-anchor="middle" class="axis" x="23" y="300">0</text>
<line x1="69.17" y1="0" x2="69.17" y2="287" stroke="black" />
<text text-anchor="middle" class="axis" x="69.17" y="300">0.5</text>
<line x1="115.33" y1="0" x2="115.33" y2="287" stroke="black" />
<text text-anchor="middle" class="axis" x="115.33" y="300">1</text>
<line x1="161.5" y1="0" x2="161.5" y2="287" stroke="black" />
<text text-anchor="middle" class="axis" x="161.5" y="300">1.5</text>
<line x1="207.67" y1="0" x2="207.67" y2="287" stroke="black" />
<text text-anchor="middle" class="axis" x="207.67" y="300">2</text>
<line x1="253.83" y1="0" x2="253.83" y2="287" stroke="black" />
<text text-anchor="middle" class="axis" x="253.83" y="300">2.5</text>
<line x1="300" y1="0" x2="300" y2="287" stroke="black" />
<text text-anchor="end" class="axis" x="300" y="300">3</text>

<path d="M 23 172.2 L 207.66666666666669 172.2 L 207.66666666666669 287" fill="none" stroke="black" stroke-width="3"/>
<circle cx="207.67" cy="172.20" r="5"/>

<circle cx="23.00" cy="258.30" r="5"/>
<circle cx="115.33" cy="229.60" r="5"/>
<circle cx="207.67" cy="143.50" r="5"/>
<circle cx="300.00" cy="0.00" r="5"/>

<path d="M 23.00 287.00 C 23.00,287.00 99.440,268.18 115.33,258.30 C 146.89,238.68 184.172,201.41 207.67,172.20 C 207.67,172.20 259.714,97.92 300.00,28.70 " fill="none" stroke="black" stroke-width="3" stroke-linecap="round"/>
<path d="M 23.00 258.30 C 23.00,258.30 99.440,239.48 115.33,229.60 C 146.89,209.98 184.172,172.71 207.67,143.50 C 207.67,143.50 259.714,69.22 300.00,0.00 " fill="none" stroke="black" stroke-width="3" stroke-linecap="round"/>

</svg>`,
			err: nil,
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
