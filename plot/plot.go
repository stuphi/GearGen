package plot

import (
  "github.com/ajstarks/svgo"
  "github.com/stuphi/GearGen/gear"
  "os"
  "fmt"
  "math"
)

const (
  RadToDeg = 180.0 / math.Pi
  DegToRad = math.Pi / 180.0
  RadToGrad = 200.0 / math.Pi
  GradToDeg = math.Pi / 200.0
  factor = 1000
)

func style(s string) string {
  switch s {
  case "dash":
    return fmt.Sprintf("fill:none; stroke-width:%d; stroke:black; stroke-dasharray:%d,%d,%d,%d",
      int(0.2 * factor), 3 * factor, 1 * factor, 1 * factor, 1 * factor)
  case "solid":
    return fmt.Sprintf("fill:none; stroke-width:%d; stroke:black",
      int(0.3 * factor))
  case "thin":
    return fmt.Sprintf("fill:none; stroke-width:%d; stroke:black",
      int(0.1 * factor))
  case "grid":
    return fmt.Sprintf("fill:none; stroke-width:%d; stroke:lightgrey",
      int(0.1 * factor))
  }
  return "fill:none; stroke:none"
}

func plotGrid(cx int, cy int, width int, height int, canvas *svg.SVG){
  spaceing := 5 * factor // 5mm

  gx := int(math.Mod(float64(cx), float64(spaceing)))
  gy := int(math.Mod(float64(cy), float64(spaceing)))

  gw := width - (2 * gx)
  gh := height - (2 * gy)

  canvas.Grid(gx, gy, gw, gh, spaceing, style("grid"))
}

func involute_intersect_angle(br, r float64) float64 {
  return math.Sqrt(math.Pow(r/br, 2) - 1)
} //= sqrt (pow (radius/base_radius, 2) - 1) * 180 / pi;

func xy_location(br, ang float64) (float64, float64) {
  x := br*(math.Cos(ang) + ang * math.Sin(ang))
  y := br*(math.Sin(ang) - ang * math.Cos(ang))
  return x, y
}

func plotInvCurve(g gear.Gear, canvas *svg.SVG){
  var px []int
  var py []int
  var pyi []int
  var x, y float64
  var r, sr, ang float64
  var offsetAng float64
  br := g.GetBaseCircleDia() * factor / 2 // Base Radius
  or := g.GetOutsideDia() * factor / 2 // Outside Radius
  rr := g.GetRootCircleDia() * factor / 2 // Root Radius
  pr := g.Pd * factor / 2 // Pitch Circle Radius
  ang = involute_intersect_angle(br, pr)
  x, y = xy_location(br, ang)
  offsetAng = math.Atan(y / x) * -1
  offsetAng += (math.Pi / (float64(g.N)/2.0))/-4.0
  if rr > br {
    sr = rr
  }else{
    sr = br
  }
  rinc := (or - sr) / 100
  for r = sr; r<=or; r += rinc {
    ang = involute_intersect_angle(br, r)
    x, y = xy_location(br, ang)
    px = append(px, int(x))
    py = append(py, int(y))
    pyi = append(pyi, int(y) * -1)
  }
  canvas.Gtransform(fmt.Sprintf("rotate(%0.4f)", offsetAng * RadToDeg))
  canvas.Polyline(px, py, style("solid"))
  canvas.Gend()
  canvas.Gtransform(fmt.Sprintf("rotate(%0.4f)", offsetAng * RadToDeg * -1.0))
  canvas.Polyline(px, pyi, style("solid"))
  canvas.Gend()

}

func plotGear(cx int, cy int, rot float64, g gear.Gear, canvas *svg.SVG){
  canvas.Gtransform(fmt.Sprintf("translate(%d, %d)", cx, cy))
  canvas.Circle(0, 0, int(g.GetOutsideDia() * factor / 2), style("solid"))
  canvas.Circle(0, 0, int(g.Pd * factor / 2), style("dash"))
  canvas.Circle(0, 0, int(g.GetRootCircleDia() * factor / 2), style("dash"))
  canvas.Circle(0, 0, int(g.GetBaseCircleDia() * factor / 2), style("thin"))
  cntrLen := int(g.GetOutsideDia() * factor / 8)
  canvas.Line( -cntrLen, 0, cntrLen, 0, style("solid"))
  canvas.Line(0, -cntrLen, 0, cntrLen, style("solid"))
  canvas.Gtransform(fmt.Sprintf("rotate(%0.3f)", rot))
  for i := 0; i<g.N; i++ {
//  for i := 0; i<1; i++ {
    canvas.Line(int((math.Cos((360/float64(g.N))*float64(i)*DegToRad)*g.GetRootCircleDia()*factor/2)),
                int((math.Sin((360/float64(g.N))*float64(i)*DegToRad)*g.GetRootCircleDia()*factor/2)),
                int((math.Cos((360/float64(g.N))*float64(i)*DegToRad)*g.GetOutsideDia()*factor/2)),
                int((math.Sin((360/float64(g.N))*float64(i)*DegToRad)*g.GetOutsideDia()*factor/2)),
                style("thin"))
    canvas.Gtransform(fmt.Sprintf("rotate(%f)", 360.0/float64(g.N)*float64(i)))
    plotInvCurve(g, canvas)
    canvas.Gend()
  }
  canvas.Gend()
  canvas.Gend()
}

func Plot(g1, g2 gear.Gear) {
  var width, height int

  border := 5.0

  if g1.Pd > g2.Pd {
    height = int(g1.GetOutsideDia() + (2 * border))
  }else{
    height = int(g2.GetOutsideDia() + (2 * border))
  }
  width = int((g1.Pd + g2.Pd + g1.GetAddendum() + g2.GetAddendum()) + (2 * border))


  centerDist := (g1.Pd + g2.Pd) / 2
  cx := int((border + (g1.GetOutsideDia() / 2.0)) * factor)
  cy := height * factor / 2
  f, err := os.Create("TestFile.svg")
  if err != nil {
    fmt.Println("Something failed creating file!")
  }
  canvas := svg.New(f)

  // Setup canvas so that each drawing unit is 0.01mm.
  canvas.StartviewUnit(width, height, "mm", 0, 0, width * factor, height * factor)
  plotGrid(cx, cy, width * factor, height * factor, canvas)
  rot := 0.0
  plotGear(cx, cy, rot, g1, canvas)
  cx = cx + int(centerDist * factor)
  if math.Mod(float64(g2.N), 2) == 0{
    rot = 180.0 / float64(g2.N)
  }
  plotGear(cx, cy, rot, g2, canvas)

  // canvas.Text(width *100 /2, height * 100 /2, "Hello, SVG", "text-anchor:middle;font-size:300;fill:black")
  canvas.End()
  f.Sync()
  f.Close()
}
