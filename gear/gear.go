package gear

import (
  "fmt"
  "math"
)

const (
    RadToDeg = 180.0 / math.Pi
    DegToRad = math.Pi / 180.0
    RadToGrad = 200.0 / math.Pi
    GradToDeg = math.Pi / 200.0
)

type Gear struct {
  Pd float64 // Pitch Diameter
  N int // Number of teeth
  A float64 // pressure angle
}

func (g Gear) GetDiametricPitch() float64 {
  return float64(g.N) / g.Pd
}

func (g Gear) GetClearence() float64 {
  return g.A / 100.0
}

func (g Gear) GetAddendum() float64 {
  return 1.0 / g.GetDiametricPitch()
}

func (g Gear) GetDedendum() float64 {
  return (1.0 + g.GetClearence()) / g.GetDiametricPitch()
}

func (g Gear) GetOutsideDia() float64 {
  return (float64(g.N) + 2.0) / g.GetDiametricPitch()
}

func (g Gear) GetBaseCircleDia() float64 {
  return g.Pd * math.Cos(g.A * DegToRad)
}

func (g Gear) GetChordalToothThickness() float64 {
  return g.Pd * math.Sin((90 * DegToRad) / float64(g.N))
}

func (g Gear) GetAngularToothThickness() float64 {
  return 360 / float64(g.N) / 2
}

func (g Gear) GetRootCircleDia() float64 {
  return g.Pd - (2 * g.GetDedendum())
}


/**
 * Return the alpha angle from the root to the point the involute crosses the pitch circle.
 *
 * @return  The angle between the root and the point the involute crosses the pitch circle.
 */
func (g Gear) GetAlphaAngle() float64 {
  return math.Sqrt(math.Pow(g.Pd, 2) - math.Pow(g.GetBaseCircleDia(), 2))/g.GetBaseCircleDia() * RadToDeg - g.A
}

/**
 * Spit out a load of text that describes this gear.
 *
 * @return  Return a text description of this gears parameters
 */

func (g Gear) String() string {
  var retval string
  retval += fmt.Sprintf("Pitch Diameter:          %.3f\n", g.Pd)
  retval += fmt.Sprintf("Outside Diameter:        %.3f\n", g.GetOutsideDia())
  retval += fmt.Sprintf("Diametric Pitch:         %.3f\n", g.GetDiametricPitch())
  retval += fmt.Sprintf("Clearance:               %.3f\n", g.GetClearence())
  retval += fmt.Sprintf("Addendum:                %.3f\n", g.GetAddendum())
  retval += fmt.Sprintf("Dedendum:                %.3f\n", g.GetDedendum())
  retval += fmt.Sprintf("Base Circle Diameter:    %.3f\n", g.GetBaseCircleDia())
  retval += fmt.Sprintf("Root Circle Diameter:    %.3f\n", g.GetRootCircleDia())
  retval += fmt.Sprintf("Chordal Tooth Thickness: %.3f\n", g.GetChordalToothThickness())
  retval += fmt.Sprintf("Angular Tooth Thickness: %.3f\n", g.GetAngularToothThickness())
  retval += fmt.Sprintf("Alpha Angle:             %.3f\n", g.GetAlphaAngle())
  return retval
}
