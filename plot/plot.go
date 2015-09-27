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
    factor = 100
)

const (
  dashStyle = "fill:none; stroke-width:20; stroke:black; stroke-dasharray:300,100,100,100"
  solidStyle = "fill:none; stroke-width:30; stroke:black"
  thinStyle = "fill:none; stroke-width:10; stroke:black"
  gridStyle = "fill:none; stroke-width:10; stroke:lightgrey"
)

func plotGrid(cx int, cy int, width int, height int, canvas *svg.SVG){
  spaceing := 5 * factor // 5mm

  gx := int(math.Mod(float64(cx), float64(spaceing)))
  gy := int(math.Mod(float64(cy), float64(spaceing)))

  gw := width - (2 * gx)
  gh := height - (2 * gy)

  canvas.Grid(gx, gy, gw, gh, spaceing, gridStyle)
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
  var x, y float64
  var r, ang float64
  br := g.GetBaseCircleDia() * factor / 2
  or := g.GetOutsideDia() * factor /2
  rinc := (or - br) / 10
  for r = br; r<=or; r += rinc {
    ang = involute_intersect_angle(br, r)
    x, y = xy_location(br, ang)
    px = append(px, int(x))
    py = append(py, int(y))
  }
  canvas.Polyline(px, py, solidStyle)

}

/*func plotInvCurve(g gear.Gear, canvas *svg.SVG){
  var px []int
  var py []int
  var xc, yc, rc float64
  var x, y int
  var ang float64
  var s float64
  r := g.GetBaseCircleDia() * factor / 2
  for i:=0.0; i<1; i = i + 0.01 {
    ang = i * 90 * DegToRad
    s = (math.Pi * r * i)/2
    xc = r * math.Cos(ang)
    yc = r * math.Sin(ang)
    x = int(xc+(s * math.Sin(ang)))
    y = int(yc-(s * math.Cos(ang)))
    rc = math.Sqrt(math.Pow(float64(x), 2) + math.Pow(float64(y), 2))
    if rc < (g.GetOutsideDia() * factor / 2) {
      px = append(px, x)
      py = append(py, y)
    } else {
      i = 1
    }
  }
  canvas.Polyline(px, py, solidStyle)
}*/

func plotGear(cx int, cy int, rot float64, g gear.Gear, canvas *svg.SVG){
  canvas.Gtransform(fmt.Sprintf("translate(%d, %d)", cx, cy))
  canvas.Circle(0, 0, int(g.GetOutsideDia() * factor / 2), solidStyle)
  canvas.Circle(0, 0, int(g.Pd * factor / 2), dashStyle)
  canvas.Circle(0, 0, int(g.GetRootCircleDia() * factor / 2), dashStyle)
  canvas.Circle(0, 0, int(g.GetBaseCircleDia() * factor / 2), thinStyle)
  cntrLen := int(g.GetOutsideDia() * factor / 8)
  canvas.Line( -cntrLen, 0, cntrLen, 0, solidStyle)
  canvas.Line(0, -cntrLen, 0, cntrLen, solidStyle)
  canvas.Gtransform(fmt.Sprintf("rotate(%0.3f)", rot))
  for i := 0; i<g.N; i++ {
    canvas.Line(int((math.Cos((360/float64(g.N))*float64(i)*DegToRad)*g.GetRootCircleDia()*factor/2)),
                int((math.Sin((360/float64(g.N))*float64(i)*DegToRad)*g.GetRootCircleDia()*factor/2)),
                int((math.Cos((360/float64(g.N))*float64(i)*DegToRad)*g.GetOutsideDia()*factor/2)),
                int((math.Sin((360/float64(g.N))*float64(i)*DegToRad)*g.GetOutsideDia()*factor/2)),
                thinStyle)
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
