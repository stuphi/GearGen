// GearGen -- Simple utility to generate ger profiles in SVG format
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
	"fmt"
	"github.com/stuphi/GearGen/gear"
	"github.com/stuphi/GearGen/plot"
	"math"
	"strconv"
)

func main() {
	fmt.Println("This is my gear program")
	fmt.Println("#######################")

	var Centers float64 // Distance between Centers
	var Ratio float64   // Required Ratio
	var DriveTeeth int  // Number of teeth on drive gear
	var DrivenTeeth int  // Number of teeth on drive gear
	var PressureAngle float64

	/*Centers = 150.0
	Ratio = 3.0
	DriveTeeth = 18
	PressureAngle = 20 */

	var input string
	var err error

	fmt.Print("Enter Centre Distance: ")
	fmt.Scanln(&input)
	Centers, err = strconv.ParseFloat(input, 64)
	if err != nil || Centers == 0 {
		fmt.Println("Didn't understand that. Using 100.")
		Centers = 100.0
	}

	fmt.Print("Enter Number of Drive Teeth: ")
	fmt.Scanln(&input)
	var tmpTeeth int64
	tmpTeeth, err = strconv.ParseInt(input, 10, 32)
	if err != nil || tmpTeeth < 2 {
		fmt.Println("Didn't understand that. Using 14.")
		DriveTeeth = 14
	} else {
		DriveTeeth = int(tmpTeeth)
	}

	fmt.Print("Enter Number of Driven Teeth: ")
	fmt.Scanln(&input)
	tmpTeeth, err = strconv.ParseInt(input, 10, 32)
	if err != nil || tmpTeeth < 2 {
		fmt.Println("Didn't understand that. Using 23.")
		DrivenTeeth = 23
	} else {
		DrivenTeeth = int(tmpTeeth)
	}


	fmt.Print("Enter Pressure Angle: ")
	fmt.Scanln(&input)
	PressureAngle, err = strconv.ParseFloat(input, 64)
	if err != nil || PressureAngle == 0 {
		fmt.Println("Didn't understand that. Using 25.0.")
		PressureAngle = 25.0
	}

	Ratio = float64(DrivenTeeth) / float64(DriveTeeth)

	var Gear1 gear.Gear
	Gear1.Pd = (1 / (Ratio + 1)) * Centers * 2
	Gear1.N = DriveTeeth
	Gear1.A = PressureAngle

	var Gear2 gear.Gear
	Gear2.Pd = (Ratio / (Ratio + 1)) * Centers * 2
	Gear2.N = int(math.Floor((float64(DriveTeeth) * Ratio) + 0.5))
	Gear2.A = PressureAngle

	fmt.Println("################################")
	fmt.Println("#        First Gear            #")
	fmt.Println("################################")
	fmt.Print(Gear1)
	fmt.Println("################################")
	fmt.Println("#       Second Gear            #")
	fmt.Println("################################")
	fmt.Print(Gear2)

	fmt.Printf("Actual Ratio: %0.4f\n", Gear2.Pd/Gear1.Pd)

	fmt.Println("################################")
	plot.Plot(Gear1, Gear2)
}
