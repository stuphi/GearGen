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

// GearGen will produce an SVG image of a pair of meshing gears. Run the command
// with the -h option to get usage details.
package main

import (
	"flag"
	"strconv"
	"github.com/stuphi/GearGen/gear"
	"github.com/stuphi/GearGen/plot"
)

func main() {

	var Centres float64 // Distance between Centres
	var Ratio float64   // Required Ratio
	var DriveTeeth int  // Number of teeth on drive gear
	var DrivenTeeth int // Number of teeth on drive gear
	var PressureAngle float64
	var Backlash float64
	var Rotation int    // Percent of rotation
	var FileName string // File name for output.

	var pCentres = flag.Int("c", 100, "Distance between centres. (Whole mm only)")
	var pDriveTeeth = flag.Int("n1", 7, "Number of teeth on the first gear")
	var pDrivenTeeth = flag.Int("n2", 23, "Number of teeth on the second gear")
	var pPressureAngle = flag.Int("p", 25, "Pressure angle")
	var pBacklash = flag.String("b", "0.5", "Backlash angle (degrees)")
	var pFileName = flag.String("o", "", "Output file name, .svg will be appended. stdout if not given")
	var pRotation = flag.Int("r", 0, "Rotation as percentage of one tooth")
	flag.Parse()
	Centres = float64(*pCentres)
	DriveTeeth = *pDriveTeeth
	DrivenTeeth = *pDrivenTeeth
	PressureAngle = float64(*pPressureAngle)
	var err error
	Backlash, err = strconv.ParseFloat(*pBacklash, 64)
	if err != nil {
		Backlash = 0.0
	}
	Rotation = *pRotation
	FileName = *pFileName

	Ratio = float64(DrivenTeeth) / float64(DriveTeeth)

	var Gear1 gear.Gear
	Gear1.Pd = (1 / (Ratio + 1)) * Centres * 2
	Gear1.N = DriveTeeth
	Gear1.A = PressureAngle
	Gear1.B = Backlash

	var Gear2 gear.Gear
	Gear2.Pd = (Ratio / (Ratio + 1)) * Centres * 2
	Gear2.N = DrivenTeeth
	Gear2.A = PressureAngle
	Gear2.B = Backlash

	plot.Plot(Gear1, Gear2, Rotation, FileName)
}
