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

func TestDiametricPitch(t *testing.T) {
	cases := []struct {
		inPd, inN, inA, want float64
	}{
		{100, 10, 30, 0.1},
		{200, 8, 25,  0.04},
	}
	for _, c := range cases {
    g := Gear{
      Pd: c.inPd,
      N: c.inN,
      A: c.inA,
    }
		got := g.getDiametricPitch()
		if got != c.want {
			t.Errorf("getDiametricPitch(Pd %f, N %f, A %f) == %f, want %f", c.inPd, c.inN, c.inA, got, c.want)
		}
	}
}

func TestClearence(t *testing.T) {
	cases := []struct {
		inPd, inN, inA, want float64
	}{
		{100, 10, 30, 0.3},
		{200, 8, 25,  0.25},
	}
	for _, c := range cases {
    g := Gear{
			Pd: c.inPd,
      N: c.inN,
      A: c.inA,
    }
		got := g.getClearence()
		if got != c.want {
			t.Errorf("getClearence(Pd %f, N %f, A %f) == %f, want %f", c.inPd, c.inN, c.inA, got, c.want)
		}
	}
}

func TestAddendum(t *testing.T) {
	cases := []struct {
		inPd, inN, inA, want float64
	}{
		{100, 10, 30, 10},
		{200, 8, 25,  25},
	}
	for _, c := range cases {
    g := Gear{
			Pd: c.inPd,
      N: c.inN,
      A: c.inA,
		}
		got := g.getAddendum()
		if got != c.want {
			t.Errorf("getAddendum(Pd %f, N %f, A %f) == %f, want %f", c.inPd, c.inN, c.inA, got, c.want)
		}
	}
}

func TestDedendum(t *testing.T) {
	cases := []struct {
		inPd, inN, inA, want float64
	}{
		{100, 10, 30, 13},
		{200, 8, 25,  31.25},
	}
	for _, c := range cases {
    g := Gear{
			Pd: c.inPd,
      N: c.inN,
      A: c.inA,
    }
		got := g.getDedendum()
		if got != c.want {
			t.Errorf("getDedendum(Pd %f, N %f, A %f) == %f, want %f", c.inPd, c.inN, c.inA, got, c.want)
		}
	}
}

func TestOutsideDia(t *testing.T) {
	cases := []struct {
		inPd, inN, inA, want float64
	}{
		{100, 10, 30, 120},
		{200, 8, 25,  250},
	}
	for _, c := range cases {
    g := Gear{
			Pd: c.inPd,
      N: c.inN,
      A: c.inA,
    }
		got := g.getOutsideDia()
		if got != c.want {
			t.Errorf("getOutsideDia(Pd %f, N %f, A %f) == %f, want %f", c.inPd, c.inN, c.inA, got, c.want)
		}
	}
}

func TestBaseCircleDia(t *testing.T) {
	cases := []struct {
		inPd, inN, inA, want float64
	}{
		{100, 10, 30, 86.603},
		{200, 8, 25,  181.262},
	}
	for _, c := range cases {
    g := Gear{
			Pd: c.inPd,
      N: c.inN,
      A: c.inA,
    }
		got := RoundPlus(g.getBaseCircleDia(), 3)
		if got != c.want {
			t.Errorf("getBaseCircleDia(Pd %.3f, N %.3f, A %.3f) == %.3f, want %.3f", c.inPd, c.inN, c.inA, got, c.want)
		}
	}
}

func TestChordalToothThickness(t *testing.T) {
	cases := []struct {
		inPd, inN, inA, want float64
	}{
		{100, 10, 30, 15.643},
		{200, 8, 25,  39.018},
	}
	for _, c := range cases {
    g := Gear{
			Pd: c.inPd,
      N: c.inN,
      A: c.inA,
    }
		got := RoundPlus(g.getChordalToothThickness(), 3)
		if got != c.want {
			t.Errorf("getChordalToothThickness(Pd %.3f, N %.3f, A %.3f) == %.3f, want %.3f", c.inPd, c.inN, c.inA, got, c.want)
		}
	}
}

func TestAngularToothThickness(t *testing.T) {
	cases := []struct {
		inPd, inN, inA, want float64
	}{
		{100, 10, 30, 18.000},
		{200, 8, 25,  22.500},
	}
	for _, c := range cases {
    g := Gear{
			Pd: c.inPd,
      N: c.inN,
      A: c.inA,
    }
		got := RoundPlus(g.getAngularToothThickness(), 3)
		if got != c.want {
			t.Errorf("getAngularToothThickness(Pd %.3f, N %.3f, A %.3f) == %.3f, want %.3f", c.inPd, c.inN, c.inA, got, c.want)
		}
	}
}

func TestRootCircleDia(t *testing.T) {
	cases := []struct {
		inPd, inN, inA, want float64
	}{
		{100, 10, 30, 74.000},
		{200, 8, 25,  137.500},
	}
	for _, c := range cases {
    g := Gear{
			Pd: c.inPd,
      N: c.inN,
      A: c.inA,
    }
		got := RoundPlus(g.getRootCircleDiameter(), 3)
		if got != c.want {
			t.Errorf("getRootCircleDiameter(Pd %.3f, N %.3f, A %.3f) == %.3f, want %.3f", c.inPd, c.inN, c.inA, got, c.want)
		}
	}
}

func TestAlphaAngle(t *testing.T) {
	cases := []struct {
		inPd, inN, inA, want float64
	}{
		{100, 10, 30, 3.080},
		{200, 8, 25,  1.717},
	}
	for _, c := range cases {
    g := Gear{
			Pd: c.inPd,
      N: c.inN,
      A: c.inA,
    }
		got := RoundPlus(g.getAlphaAngle(), 3)
		if got != c.want {
			t.Errorf("getAlphaAngle(Pd %.3f, N %.3f, A %.3f) == %.3f, want %.3f", c.inPd, c.inN, c.inA, got, c.want)
		}
	}
}
