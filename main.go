package main

import (
	"fmt"
	"math"
	"strconv"
	"github.com/stuphi/GearGen/gear"
	"github.com/stuphi/GearGen/plot"
)

func main() {
	fmt.Println("This is my gear program")
	fmt.Println("#######################")

	var Centers float64 // Distance between Centers
	var Ratio float64 // Required Ratio
	var DriveTeeth int // Number of teeth on drive gear
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
	if err != nil {
		fmt.Println("Didn't understand that. Using 150.")
		Centers = 150.0
	}

	fmt.Print("Enter Ratio: ")
	fmt.Scanln(&input)
	Ratio, err = strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("Didn't understand that. Using 3.0.")
		Ratio = 3.0
	}

	fmt.Print("Enter Number of Drive Teeth: ")
	fmt.Scanln(&input)
	var tmpTeeth int64
	tmpTeeth, err = strconv.ParseInt(input, 10, 32)
	if err != nil {
		fmt.Println("Didn't understand that. Using 18.")
		DriveTeeth = 18
	}else{
		DriveTeeth = int(tmpTeeth)
	}

	fmt.Print("Enter Pressure Angle: ")
	fmt.Scanln(&input)
	PressureAngle, err = strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("Didn't understand that. Using 20.0.")
		PressureAngle = 20.0
	}

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

	fmt.Printf("Actual Ratio: %0.4f\n", Gear2.Pd / Gear1.Pd)

	fmt.Println("################################")
	plot.Plot(Gear1)
}
