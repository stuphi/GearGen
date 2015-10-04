// GearGen -- Simple utility to generate gear profiles in SVG format
// Copyright (C) 2015  Philip Stubbs
//
// GearGen is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// GearGen is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// A package to convert our gears to an SVG image.
package plot

import (
	"fmt"
	"github.com/ajstarks/svgo"
	"github.com/stuphi/GearGen/gear"
	"math"
	"os"
)

// Some useful conversions. Not all are used.
const (
	RadToDeg  = 180.0 / math.Pi
	DegToRad  = math.Pi / 180.0
	RadToGrad = 200.0 / math.Pi
	GradToDeg = math.Pi / 200.0
	// This factor is used as the svg library only accepts int. Therefore we need
	// to multiply all dimensions by this factor and use a viewport in the svg
	// document to be able to get sub mm acuracy. By using 1000 here, each unit
	// will be 0.001 mm in size.
	factor = 1000
)

// Return the apropriate style string for the requested line type.
func style(s string) string {
	switch s {
	case "dash":
		return fmt.Sprintf("fill:none; stroke-width:%d; stroke:black; stroke-dasharray:%d,%d,%d,%d",
			int(0.1*factor), 3*factor, 1*factor, 1*factor, 1*factor)
	case "solid":
		return fmt.Sprintf("fill:none; stroke-width:%d; stroke:black",
			int(0.25*factor))
	case "thin":
		return fmt.Sprintf("fill:none; stroke-width:%d; stroke:black",
			int(0.1*factor))
	case "grid":
		return fmt.Sprintf("fill:none; stroke-width:%d; stroke:lightgrey",
			int(0.1*factor))
	case "anott":
		return fmt.Sprintf("text-anchor:middle;font-size:%d;fill:#888888;fill-opacity:0.5",
			int(5*factor))

	}
	return "fill:none; stroke:none"
}

// Plot a grid on our canvas.
// cx and cy are the centre of the first gear in drawing units
// width and height are the size of the drawing in drawing units
// canvas is the canvas to draw to.
func plotGrid(cx int, cy int, width int, height int, canvas *svg.SVG) {
	// Set the spacing in drawing units
	spaceing := 5 * factor // 5mm

	// Calculate the start of the grid such that the centre of the first gear will
	// be on the grid
	gx := int(math.Mod(float64(cx), float64(spaceing)))
	gy := int(math.Mod(float64(cy), float64(spaceing)))

	// Set the width and height so that there is an even border all round.
	// could do with trying to sort out the width of the canvas so that the right
	// edge of the grid ends with a vertical line.
	gw := width - (2 * gx)
	gh := height - (2 * gy)

	canvas.Grid(gx, gy, gw, gh, spaceing, style("grid"))
}

// Calculate the involute angle at witch point the involute reaches the
// requested radius
func involuteIntersectAngle(br, r float64) float64 {
	return math.Sqrt(math.Pow(r/br, 2) - 1)
}

// Calculate the xy cordinates of the involute at a given angle
func xyLocation(br, ang float64) (float64, float64) {
	x := br * (math.Cos(ang) + ang*math.Sin(ang))
	y := br * (math.Sin(ang) - ang*math.Cos(ang))
	return x, y
}

// Given an xy point, calculate the new position if rotated by given angle
func rotXY(x, y int, ang float64) (int, int) {
	var xf, yf float64
	xf = float64(x)
	yf = float64(y)
	hf := math.Sqrt(math.Pow(xf, 2) + math.Pow(yf, 2))
	a := math.Atan(yf/xf) + ang
	xf = math.Cos(a) * hf
	yf = math.Sin(a) * hf
	return int(xf), int(yf)
}

// Plot the involute gear profile for a single tooth
func plotInvCurve(g gear.Gear, canvas *svg.SVG) {
	var px []int
	var py []int
	var pyi []int
	var x, y float64
	var r, sr, ang float64
	var offsetAng float64
	br := g.GetBaseCircleDia() * factor / 2 // Base Radius
	or := g.GetOutsideDia() * factor / 2    // Outside Radius
	rr := g.GetRootCircleDia() * factor / 2 // Root Radius
	pr := g.Pd * factor / 2                 // Pitch Circle Radius
	ang = involuteIntersectAngle(br, pr)
	x, y = xyLocation(br, ang)
	offsetAng = math.Atan(y/x) * -1
	offsetAng += (math.Pi / (float64(g.N) / 2.0)) / -4.0
	if rr > br {
		sr = rr
	} else {
		sr = br
	}
	rinc := (or - sr) / 100
	for r = sr; r <= or; r += rinc {
		ang = involuteIntersectAngle(br, r)
		x, y = xyLocation(br, ang)
		px = append(px, int(x))
		py = append(py, int(y))
		pyi = append(pyi, int(y)*-1)
	}
	if br > rr {
		px = append([]int{int(rr)}, px...)
		py = append([]int{0}, py...)
		pyi = append([]int{0}, pyi...)
	}
	canvas.Gtransform(fmt.Sprintf("rotate(%0.4f)", offsetAng*RadToDeg))
	canvas.Polyline(px, py, style("solid"))
	canvas.Gend()
	canvas.Gtransform(fmt.Sprintf("rotate(%0.4f)", offsetAng*RadToDeg*-1.0))
	canvas.Polyline(px, pyi, style("solid"))
	canvas.Gend()
	sx, sy := rotXY(px[len(px)-1], py[len(py)-1], offsetAng)
	ex, ey := rotXY(px[len(px)-1], pyi[len(pyi)-1], offsetAng*-1.0)
	canvas.Path(fmt.Sprintf("M%d,%d A%d,%d 0 0 1 %d,%d",
		sx, sy, int(or), int(or),
		ex, ey), style("solid"))

	sx, sy = rotXY(px[0], py[0], -offsetAng)
	ex, ey = rotXY(px[0], py[0], ((2*math.Pi)/float64(g.N))+offsetAng)
	canvas.Path(fmt.Sprintf("M%d,%d A%d,%d 0 0 1 %d,%d",
		sx, sy, int(rr), int(rr),
		ex, ey), style("solid"))
}

// Plot a complete gear at cx,cy rotated by angle rot.
func plotGear(cx int, cy int, rot float64, g gear.Gear, canvas *svg.SVG) {
	canvas.Gtransform(fmt.Sprintf("translate(%d, %d)", cx, cy))
	canvas.Circle(0, 0, int(g.Pd*factor/2), style("dash"))
	cntrLen := int(g.GetOutsideDia() * factor / 8)
	canvas.Line(-cntrLen, 0, cntrLen, 0, style("solid"))
	canvas.Line(0, -cntrLen, 0, cntrLen, style("solid"))
	canvas.Gtransform(fmt.Sprintf("rotate(%0.3f)", rot))
	for i := 0; i < g.N; i++ {
		canvas.Line(int((math.Cos((360/float64(g.N))*float64(i)*DegToRad) *
			g.GetRootCircleDia() * factor / 2)),
			int((math.Sin((360/float64(g.N))*float64(i)*DegToRad) *
				g.GetRootCircleDia() * factor / 2)),
			int((math.Cos((360/float64(g.N))*float64(i)*DegToRad) *
				g.GetOutsideDia() * factor / 2)),
			int((math.Sin((360/float64(g.N))*float64(i)*DegToRad) *
				g.GetOutsideDia() * factor / 2)),
			style("dash"))
		canvas.Gtransform(fmt.Sprintf("rotate(%f)", 360.0/float64(g.N)*float64(i)))
		plotInvCurve(g, canvas)
		canvas.Gend()
	}
	canvas.Gend()
	anottext := fmt.Sprintf("Pitch Dia: %0.1f", g.Pd)
	canvas.Text(0, -1 * factor, anottext, style("anott"))
	anottext = fmt.Sprintf("Teeth: %d", g.N)
	canvas.Text(0, 5 * factor, anottext, style("anott"))
	anottext = fmt.Sprintf("Pressure Angle: %0.1f", g.A)
	canvas.Text(0, 11 * factor, anottext, style("anott"))
	canvas.Gend()
}

// Plot the complete drawing of gears g1 and g2 to file fname or stdout if no
// file is given.
// rotfrac represents the percentage of one tooth to rotate both gears. Used
// to be able to draw the gears at different stages of engagment.
func Plot(g1, g2 gear.Gear, rotfrac int, fname string) {
	var width, height int

	border := 5.0

	// Determin the size of our canvas.
	if g1.Pd > g2.Pd {
		height = int(g1.GetOutsideDia() + (2 * border))
	} else {
		height = int(g2.GetOutsideDia() + (2 * border))
	}
	width = int((g1.Pd + g2.Pd + g1.GetAddendum() + g2.GetAddendum()) +
		(2 * border))

	centerDist := (g1.Pd + g2.Pd) / 2
	cx := int((border + (g1.GetOutsideDia() / 2.0)) * factor)
	cy := height * factor / 2
	var canvas *svg.SVG
	var f *os.File
	var err error
	// If we are writing to file, open the file or quit if there is an error.
	if fname != "" {
		f, err = os.Create(fmt.Sprintf("%s.svg", fname))
		if err != nil {
			fmt.Println("Something failed creating file!")
			os.Exit(1)
		}
		canvas = svg.New(f)
	} else {
		// If we are not using a file, send output to stdout.
		canvas = svg.New(os.Stdout)
	}

	// Setup canvas so that each drawing unit is 0.01mm.
	canvas.StartviewUnit(width, height, "mm", 0, 0, width*factor, height*
		factor)
	plotGrid(cx, cy, width*factor, height*factor, canvas)
	rot := 0.0 + ((float64(rotfrac) / 100) * (360 / float64(g1.N)))
	plotGear(cx, cy, rot, g1, canvas)
	cx = cx + int(centerDist*factor)
	if math.Mod(float64(g2.N), 2) == 0 {
		rot = 180.0 / float64(g2.N)
	} else {
		rot = 0.0
	}
	rot -= (float64(rotfrac) / 100) * (360 / float64(g2.N))
	plotGear(cx, cy, rot, g2, canvas)

	canvas.Text((width / 2) * factor, (height - 2) * factor,
		"Generated by GearGen. http://github/stuphi/GearGen", style("anott"))

	canvas.End()
	// Cleanly close file, if we are writing to file.
	if f != nil {
		f.Sync()
		f.Close()
	}
}
