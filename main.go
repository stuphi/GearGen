package main

import (
	"fmt"
	"math"
	"github.com/stuphi/GearGen/gear"
)

func main() {
	fmt.Println("This is my gear program")
	fmt.Println("#######################")

	var Centers float64 // Distance between Centers
	var Ratio float64 // Required Ratio
	var DriveTeeth int // Number of teeth on drive gear
	var PressureAngle float64

	Centers = 150.0
	Ratio = 3.0
	DriveTeeth = 18
	PressureAngle = 20

	var Gear1 gear.Gear
	Gear1.Pd = (1 / (Ratio + 1)) * Centers * 2
	Gear1.N = DriveTeeth
	Gear1.A = PressureAngle

	var Gear2 gear.Gear
	Gear2.Pd = (Ratio / (Ratio + 1)) * Centers * 2
	Gear2.N = int(math.Floor((float64(DriveTeeth) * Ratio) + 0.5))
	Gear2.A = PressureAngle

	fmt.Print(Gear1.ToString())
	fmt.Print(Gear2.ToString())

	fmt.Printf("Actual Ratio: %0.4f\n", Gear2.Pd / Gear1.Pd)
}
