package svgPlot

import (
	"errors"
	"fmt"
)

type (
	Plot struct {
		style   Style
		angular []pointArr
		point   []pointArr
		smooth  []pointArr
	}
	pointArr struct {
		X, Y []float64
	}
)

func NewPlot(TotalHeight, TotalWidth, XDivisionsQty, YDivisionsQty uint, NameOfX, NameOfY string) (Plot Plot, err error) {
	if TotalHeight < 150 {
		err = errors.New("TotalHeight is too low")
		return
	}
	if TotalWidth < 150 {
		err = errors.New("TotalWidth is too low")
		return
	}
	if XDivisionsQty > 30 {
		err = errors.New("XDivisionsQty is too high")
		return
	}
	if YDivisionsQty > 30 {
		err = errors.New("YDivisionsQty is too high")
		return
	}
	if len(NameOfX) > 6 {
		err = errors.New("NameOfX is too long")
		return
	}
	if len(NameOfY) > 6 {
		err = errors.New("NameOfY is too long")
		return
	}
	Plot.style = Style{
		TotalHeight:   TotalHeight,
		TotalWidth:    TotalWidth,
		XDivisionsQty: XDivisionsQty,
		YDivisionsQty: YDivisionsQty,
		NameOfX:       NameOfX,
		NameOfY:       NameOfY,
	}
	return
}

func (s *Plot) AddSmooth(X, Y []float64) error {
	if len(X) != len(Y) {
		return errors.New("len(X) ≠ len(Y)")
	}
	if len(X) == 0 {
		return nil
	}
	s.smooth = append(s.smooth, pointArr{X, Y})
	return nil
}
func (s *Plot) AddAngular(X, Y []float64) error {
	if len(X) != len(Y) {
		return errors.New("len(X) ≠ len(Y)")
	}
	if len(X) == 0 {
		return nil
	}
	s.angular = append(s.angular, pointArr{X, Y})
	return nil
}
func (s *Plot) AddPoint(X, Y []float64) error {
	if len(X) != len(Y) {
		return errors.New("len(X) ≠ len(Y)")
	}
	if len(X) == 0 {
		return nil
	}
	s.point = append(s.point, pointArr{X, Y})
	return nil
}
func (s *Plot) Draw() (result string, err error) {
	var (
		yMin, yMax, xMin, xMax           float64
		tmpMin, tmpMax                   float64
		greed, plot, tmpStr              string
		x0, y0, gradX, gradY, xLen, yLen float64
		xZeroPos, yZeroPos               int
		xNums, yNums                     []string
	)
	var ()

	if len(s.style.NameOfX) > 6 {
		err = errors.New("x name max len is 6")
		return
	}
	if len(s.style.NameOfY) > 6 {
		err = errors.New("y name max len is 6")
		return
	}

	// get min & max values
	yMin, xMin = 1.8e300, 1.8e300
	yMax, xMax = -1.8e300, -1.8e300
	for _, p := range s.angular {
		tmpMin, tmpMax, err = getMinMax(p.Y)
		if err != nil {
			return
		}
		if tmpMin < yMin {
			yMin = tmpMin
		}
		if tmpMax > yMax {
			yMax = tmpMax
		}

		tmpMin, tmpMax, err = getMinMax(p.X)
		if err != nil {
			return
		}
		if tmpMin < xMin {
			xMin = tmpMin
		}
		if tmpMax > xMax {
			xMax = tmpMax
		}
	}
	for _, p := range s.point {
		tmpMin, tmpMax, err = getMinMax(p.Y)
		if err != nil {
			return
		}
		if tmpMin < yMin {
			yMin = tmpMin
		}
		if tmpMax > yMax {
			yMax = tmpMax
		}

		tmpMin, tmpMax, err = getMinMax(p.X)
		if err != nil {
			return
		}
		if tmpMin < xMin {
			xMin = tmpMin
		}
		if tmpMax > xMax {
			xMax = tmpMax
		}
	}
	for _, p := range s.smooth {
		tmpMin, tmpMax, err = getMinMax(p.Y)
		if err != nil {
			return
		}
		if tmpMin < yMin {
			yMin = tmpMin
		}
		if tmpMax > yMax {
			yMax = tmpMax
		}

		tmpMin, tmpMax, err = getMinMax(p.X)
		if err != nil {
			return
		}
		if tmpMin < xMin {
			xMin = tmpMin
		}
		if tmpMax > xMax {
			xMax = tmpMax
		}
	}

	xNums, xLen, xZeroPos, err = makeArr(xMin, xMax, s.style.XDivisionsQty)
	if err != nil {
		return
	}
	yNums, yLen, yZeroPos, err = makeArr(yMin, yMax, s.style.YDivisionsQty)
	if err != nil {
		return
	}
	if s.style.NameOfX != "" {
		xNums[len(xNums)-1] = s.style.NameOfX
	}
	if s.style.NameOfY != "" {
		yNums[len(yNums)-1] = s.style.NameOfY
	}

	greed, x0, y0, gradX, gradY, err = makeGreed(s.style.TotalHeight, s.style.TotalWidth, xNums, yNums, xLen, yLen, xZeroPos, yZeroPos)
	if err != nil {
		return
	}
	for _, p := range s.angular {
		tmpStr, err = makeAngularCurve(x0, y0, gradX, gradY, p.X, p.Y)
		if err != nil {
			return
		}
		plot += tmpStr + "\n"
	}
	for _, p := range s.point {
		tmpStr, err = makePointCurve(x0, y0, gradX, gradY, p.X, p.Y)
		if err != nil {
			return
		}
		plot += tmpStr + "\n"
	}
	for _, p := range s.smooth {
		tmpStr, err = makeSmoothCurve(x0, y0, gradX, gradY, p.X, p.Y)
		if err != nil {
			return
		}
		plot += tmpStr + "\n"
	}
	result = "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n"
	result += fmt.Sprintf("<svg width=\"%d\" height=\"%d\" viewBox=\"0 0 %d %d\" version=\"1.1\" xmlns=\"http://www.w3.org/2000/svg\">\n\"", s.style.TotalWidth, s.style.TotalHeight, s.style.TotalWidth, s.style.TotalHeight)
	result += greed + "\n"
	result += plot + "\n"
	result += "</svg>"
	return
}
