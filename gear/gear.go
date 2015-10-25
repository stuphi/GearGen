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

// A package to define our gear
package gear

import (
	"fmt"
	"math"
)

const (
	RadToDeg  = 180.0 / math.Pi
	DegToRad  = math.Pi / 180.0
	RadToGrad = 200.0 / math.Pi
	GradToDeg = math.Pi / 200.0
)

// Structure to hold the supplied parameters for our gear
type Gear struct {
	Pd float64 // Pitch Diameter
	N  int     // Number of teeth
	A  float64 // pressure angle
	B  float64 // backlash angle
}

// Calculate and return the diametric pitch
func (g Gear) GetDiametricPitch() float64 {
	return float64(g.N) / g.Pd
}

// Calculate and return clearence. This needs to be checked.
func (g Gear) GetClearence() float64 {
	return g.A / 100.0
}

// Calculate and return the gear addendum
func (g Gear) GetAddendum() float64 {
	return 1.0 / g.GetDiametricPitch()
}

// Calculate and return the gear dedendum
func (g Gear) GetDedendum() float64 {
	return (1.0 + g.GetClearence()) / g.GetDiametricPitch()
}

// Calculate and return the outside diameter
func (g Gear) GetOutsideDia() float64 {
	return (float64(g.N) + 2.0) / g.GetDiametricPitch()
}

// Calculate and return the base diameter
func (g Gear) GetBaseCircleDia() float64 {
	return g.Pd * math.Cos(g.A*DegToRad)
}

// Calculate and return the tooth chordal thickness
func (g Gear) GetChordalToothThickness() float64 {
	return g.Pd * math.Sin((90*DegToRad)/float64(g.N))
}

// Calculate and return the tooth angular thickness
func (g Gear) GetAngularToothThickness() float64 {
	return 360 / float64(g.N) / 2
}

// Calculate and return the gear root circle diameter
func (g Gear) GetRootCircleDia() float64 {
	return g.Pd - (2 * g.GetDedendum())
}

// Return the alpha angle from the root to the point the involute crosses the
// pitch circle. This needs to be checked.
func (g Gear) GetAlphaAngle() float64 {
	return math.Sqrt(math.Pow(g.Pd, 2)-math.Pow(g.GetBaseCircleDia(), 2))/
		g.GetBaseCircleDia()*RadToDeg - g.A
}

// Spit out a load of text that describes this gear.
func (g Gear) String() string {
	var retval string
	retval += fmt.Sprintf("Pitch Diameter:          %.3f\n", g.Pd)
	retval += fmt.Sprintf("Outside Diameter:        %.3f\n", g.GetOutsideDia())
	retval += fmt.Sprintf("Diametric Pitch:         %.3f\n",
		g.GetDiametricPitch())
	retval += fmt.Sprintf("Clearance:               %.3f\n", g.GetClearence())
	retval += fmt.Sprintf("Addendum:                %.3f\n", g.GetAddendum())
	retval += fmt.Sprintf("Dedendum:                %.3f\n", g.GetDedendum())
	retval += fmt.Sprintf("Base Circle Diameter:    %.3f\n", g.GetBaseCircleDia())
	retval += fmt.Sprintf("Root Circle Diameter:    %.3f\n", g.GetRootCircleDia())
	retval += fmt.Sprintf("Chordal Tooth Thickness: %.3f\n",
		g.GetChordalToothThickness())
	retval += fmt.Sprintf("Angular Tooth Thickness: %.3f\n",
		g.GetAngularToothThickness())
	retval += fmt.Sprintf("Alpha Angle:             %.3f\n", g.GetAlphaAngle())
	return retval
}
