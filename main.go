package main

import (
	"fmt"
	"github.com/stuphi/GearGen/gear"
)

func main() {
	fmt.Println("This is my gear program")
	fmt.Println("#######################")

	var myGear gear.Gear
	myGear.Pd = 100
	myGear.N = 20
	myGear.A = 25

	fmt.Print(myGear.ToString())
}
