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

package gear

import (
	"math"
	"testing"
)

func Round(f float64) float64 {
	return math.Floor(f + .5)
}

func RoundPlus(f float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	return Round(f*shift) / shift
}

// For each test we will provide three parameters. Pitch diameter, number of
// teeth, and pressure angle, along with the expected result
type testCase struct {
	inPd float64
	inN  int
	inA  float64
	want float64
}

func TestDiametricPitch(t *testing.T) {
	cases := []testCase{
		{100, 10, 30, 0.1},
		{200, 8, 25, 0.04},
	}
	for _, c := range cases {
		g := Gear{
			Pd: c.inPd,
			N:  c.inN,
			A:  c.inA,
		}
		got := g.GetDiametricPitch()
		if got != c.want {
			t.Errorf("GetDiametricPitch(Pd %f, N %f, A %f) == %f, want %f",
				c.inPd, c.inN, c.inA, got, c.want)
		}
	}
}

func TestClearence(t *testing.T) {
	cases := []testCase{
		{100, 10, 30, 0.3},
		{200, 8, 25, 0.25},
	}
	for _, c := range cases {
		g := Gear{
			Pd: c.inPd,
			N:  c.inN,
			A:  c.inA,
		}
		got := g.GetClearence()
		if got != c.want {
			t.Errorf("GetClearence(Pd %f, N %f, A %f) == %f, want %f",
				c.inPd, c.inN, c.inA, got, c.want)
		}
	}
}

func TestAddendum(t *testing.T) {
	cases := []testCase{
		{100, 10, 30, 10},
		{200, 8, 25, 25},
	}
	for _, c := range cases {
		g := Gear{
			Pd: c.inPd,
			N:  c.inN,
			A:  c.inA,
		}
		got := g.GetAddendum()
		if got != c.want {
			t.Errorf("GetAddendum(Pd %f, N %f, A %f) == %f, want %f",
				c.inPd, c.inN, c.inA, got, c.want)
		}
	}
}

func TestDedendum(t *testing.T) {
	cases := []testCase{
		{100, 10, 30, 13},
		{200, 8, 25, 31.25},
	}
	for _, c := range cases {
		g := Gear{
			Pd: c.inPd,
			N:  c.inN,
			A:  c.inA,
		}
		got := g.GetDedendum()
		if got != c.want {
			t.Errorf("GetDedendum(Pd %f, N %f, A %f) == %f, want %f",
				c.inPd, c.inN, c.inA, got, c.want)
		}
	}
}

func TestOutsideDia(t *testing.T) {
	cases := []testCase{
		{100, 10, 30, 120},
		{200, 8, 25, 250},
	}
	for _, c := range cases {
		g := Gear{
			Pd: c.inPd,
			N:  c.inN,
			A:  c.inA,
		}
		got := g.GetOutsideDia()
		if got != c.want {
			t.Errorf("GetOutsideDia(Pd %f, N %f, A %f) == %f, want %f",
				c.inPd, c.inN, c.inA, got, c.want)
		}
	}
}

func TestBaseCircleDia(t *testing.T) {
	cases := []testCase{
		{100, 10, 30, 86.603},
		{200, 8, 25, 181.262},
	}
	for _, c := range cases {
		g := Gear{
			Pd: c.inPd,
			N:  c.inN,
			A:  c.inA,
		}
		got := RoundPlus(g.GetBaseCircleDia(), 3)
		if got != c.want {
			t.Errorf("GetBaseCircleDia(Pd %.3f, N %.3f, A %.3f) == %.3f, want %.3f",
				c.inPd, c.inN, c.inA, got, c.want)
		}
	}
}

func TestChordalToothThickness(t *testing.T) {
	cases := []testCase{
		{100, 10, 30, 15.643},
		{200, 8, 25, 39.018},
	}
	for _, c := range cases {
		g := Gear{
			Pd: c.inPd,
			N:  c.inN,
			A:  c.inA,
		}
		got := RoundPlus(g.GetChordalToothThickness(), 3)
		if got != c.want {
			t.Errorf("GetChordalToothThickness(Pd %.3f, N %.3f, A %.3f) == %.3f, want %.3f",
				c.inPd, c.inN, c.inA, got, c.want)
		}
	}
}

func TestAngularToothThickness(t *testing.T) {
	cases := []testCase{
		{100, 10, 30, 18.000},
		{200, 8, 25, 22.500},
	}
	for _, c := range cases {
		g := Gear{
			Pd: c.inPd,
			N:  c.inN,
			A:  c.inA,
		}
		got := RoundPlus(g.GetAngularToothThickness(), 3)
		if got != c.want {
			t.Errorf("GetAngularToothThickness(Pd %.3f, N %.3f, A %.3f) == %.3f, want %.3f",
				c.inPd, c.inN, c.inA, got, c.want)
		}
	}
}

func TestRootCircleDia(t *testing.T) {
	cases := []testCase{
		{100, 10, 30, 74.000},
		{200, 8, 25, 137.500},
	}
	for _, c := range cases {
		g := Gear{
			Pd: c.inPd,
			N:  c.inN,
			A:  c.inA,
		}
		got := RoundPlus(g.GetRootCircleDia(), 3)
		if got != c.want {
			t.Errorf("GetRootCircleDiameter(Pd %.3f, N %.3f, A %.3f) == %.3f, want %.3f",
				c.inPd, c.inN, c.inA, got, c.want)
		}
	}
}

func TestAlphaAngle(t *testing.T) {
	cases := []testCase{
		{100, 10, 30, 3.080},
		{200, 8, 25, 1.717},
	}
	for _, c := range cases {
		g := Gear{
			Pd: c.inPd,
			N:  c.inN,
			A:  c.inA,
		}
		got := RoundPlus(g.GetAlphaAngle(), 3)
		if got != c.want {
			t.Errorf("GetAlphaAngle(Pd %.3f, N %.3f, A %.3f) == %.3f, want %.3f",
				c.inPd, c.inN, c.inA, got, c.want)
		}
	}
}
