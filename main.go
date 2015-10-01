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


package main

import (
	"flag"
	"github.com/stuphi/GearGen/gear"
	"github.com/stuphi/GearGen/plot"
)

func main() {

	var Centers float64 // Distance between Centers
	var Ratio float64   // Required Ratio
	var DriveTeeth int  // Number of teeth on drive gear
	var DrivenTeeth int  // Number of teeth on drive gear
	var PressureAngle float64
	var Rotation int	// Percent of rotation
	var FileName string  // File name for output.

	var pCenters = flag.Int("c", 100, "Distance between centers. (Whole mm only)")
	var pDriveTeeth = flag.Int("n1", 7, "Number of teeth on the first gear")
	var pDrivenTeeth = flag.Int("n2", 23, "Number of teeth on the second gear")
	var pPressureAngle = flag.Int("p", 25, "Pressure angle")
	var pFileName = flag.String("o", "", "Output file name, .svg will be appended. stdout if not given")
	var pRotation = flag.Int("r", 0, "Rotation as percentage of one tooth")
	flag.Parse()
	Centers = float64(*pCenters)
	DriveTeeth = *pDriveTeeth
	DrivenTeeth = *pDrivenTeeth
	PressureAngle = float64(*pPressureAngle)
	Rotation = *pRotation
	FileName = *pFileName

	Ratio = float64(DrivenTeeth) / float64(DriveTeeth)

	var Gear1 gear.Gear
	Gear1.Pd = (1 / (Ratio + 1)) * Centers * 2
	Gear1.N = DriveTeeth
	Gear1.A = PressureAngle

	var Gear2 gear.Gear
	Gear2.Pd = (Ratio / (Ratio + 1)) * Centers * 2
	Gear2.N = DrivenTeeth
	Gear2.A = PressureAngle

	plot.Plot(Gear1, Gear2, Rotation, FileName)
}
