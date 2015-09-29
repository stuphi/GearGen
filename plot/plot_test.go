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

package plot

import (
  "fmt"
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

func TestInvoluteIntersectAngle(t *testing.T) {
	cases := []struct { inbr, inr, want float64 }{
		{100, 110, 0.458258},
		{200, 220, 0.458258},
    {100, 100, 0.0},
    {100, 200, 1.732051},
    {100, 300, 2.828427},
	}
	for _, c := range cases {
  	got := involuteIntersectAngle(c.inbr, c.inr)
		if fmt.Sprintf("%0.6f", got) != fmt.Sprintf("%0.6f", c.want) {
			t.Errorf("involuteIntersectAngle(br %f, r %f) == %f, want %f", c.inbr, c.inr, got, c.want)
		}
	}
}

func TestXyLocation(t *testing.T) {
	cases := []struct { inbr, inang, wantX, wantY float64 }{
		{100, 0.458258, 109.955165, 3.140951},
		{200, 0.458258, 219.910331, 6.281902},
    {100, 0.0, 100,0},
    {100, 1.732051, 154.902371, 126.511906},
    {100, 2.828427, -8.000432, 299.893291},
	}
	for _, c := range cases {
  	gotX, gotY := xyLocation(c.inbr, c.inang)
		if fmt.Sprintf("%0.6f,%0.6f", gotX, gotY) !=
      fmt.Sprintf("%0.6f,%0.6f", c.wantX, c.wantY) {
			t.Errorf("xyLocation(br %f, ang %f) == %f,%f want %f,%f", c.inbr, c.inang,
        gotX, gotY, c.wantX, c.wantY)
		}
	}
}
