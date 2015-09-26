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

func (g Gear) getDiametricPitch() float64 {
  return float64(g.N) / g.Pd
}

func (g Gear) getClearence() float64 {
  return g.A / 100.0
}

func (g Gear) getAddendum() float64 {
  return 1.0 / g.getDiametricPitch()
}

func (g Gear) getDedendum() float64 {
  return (1.0 + g.getClearence()) / g.getDiametricPitch()
}

func (g Gear) GetOutsideDia() float64 {
  return (float64(g.N) + 2.0) / g.getDiametricPitch()
}

func (g Gear) getBaseCircleDia() float64 {
  return g.Pd * math.Cos(g.A * DegToRad)
}

func (g Gear) getChordalToothThickness() float64 {
  return g.Pd * math.Sin((90 * DegToRad) / float64(g.N))
}

func (g Gear) getAngularToothThickness() float64 {
  return 360 / float64(g.N) / 2
}

func (g Gear) GetRootCircleDiameter() float64 {
  return g.Pd - (2 * g.getDedendum())
}


/**
 * Return the alpha angle from the root to the point the involute crosses the pitch circle.
 *
 * @return  The angle between the root and the point the involute crosses the pitch circle.
 */
func (g Gear) getAlphaAngle() float64 {
  return math.Sqrt(math.Pow(g.Pd, 2) - math.Pow(g.getBaseCircleDia(), 2))/g.getBaseCircleDia() * RadToDeg - g.A
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
  retval += fmt.Sprintf("Diametric Pitch:         %.3f\n", g.getDiametricPitch())
  retval += fmt.Sprintf("Clearance:               %.3f\n", g.getClearence())
  retval += fmt.Sprintf("Addendum:                %.3f\n", g.getAddendum())
  retval += fmt.Sprintf("Dedendum:                %.3f\n", g.getDedendum())
  retval += fmt.Sprintf("Base Circle Diameter:    %.3f\n", g.getBaseCircleDia())
  retval += fmt.Sprintf("Root Circle Diameter:    %.3f\n", g.GetRootCircleDiameter())
  retval += fmt.Sprintf("Chordal Tooth Thickness: %.3f\n", g.getChordalToothThickness())
  retval += fmt.Sprintf("Angular Tooth Thickness: %.3f\n", g.getAngularToothThickness())
  retval += fmt.Sprintf("Alpha Angle:             %.3f\n", g.getAlphaAngle())
  return retval
}
