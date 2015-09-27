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
)

const (
  dashStyle = "fill:none; stroke-width:20; stroke:black; stroke-dasharray:300,100,100,100"
  solidStyle = "fill:none; stroke-width:30; stroke:black"
  thinStyle = "fill:none; stroke-width:10; stroke:black"
  gridStyle = "fill:none; stroke-width:8; stroke:grey"
)

func plotGrid(cx int, cy int, width int, height int, canvas *svg.SVG){
  spaceing := 500 // 5mm

  canvas.Group()
  for i:=0; i < cy; i += spaceing {
    canvas.Line(0, cy + i, width, cy + i, gridStyle)
    canvas.Line(0, cy - i, width, cy - i, gridStyle)
  }
  for i := 0; i < width - cx; i += spaceing {
    canvas.Line(cx+i, 0, cx+i, height, gridStyle)
    canvas.Line(cx-i, 0, cx-i, height, gridStyle)
  }
  canvas.Gend()
}

func plotGear(cx int, cy int, rot float64, g gear.Gear, canvas *svg.SVG){
  canvas.Gtransform(fmt.Sprintf("translate(%d, %d)", cx, cy))
  canvas.Circle(0, 0, int(g.GetOutsideDia() * 100 / 2), solidStyle)
  canvas.Circle(0, 0, int(g.Pd * 100 / 2), dashStyle)
  canvas.Circle(0, 0, int(g.GetRootCircleDia() * 100 / 2), dashStyle)
  cntrLen := int(g.GetOutsideDia() * 100 / 8)
  canvas.Line( -cntrLen, 0, cntrLen, 0, thinStyle)
  canvas.Line(0, -cntrLen, 0, cntrLen, thinStyle)
  canvas.Gtransform(fmt.Sprintf("rotate(%0.3f)", rot))
  for i := 0; i<g.N; i++ {
    canvas.Line(int((math.Cos((360/float64(g.N))*float64(i)*DegToRad)*g.GetRootCircleDia()*50)),
                int((math.Sin((360/float64(g.N))*float64(i)*DegToRad)*g.GetRootCircleDia()*50)),
                int((math.Cos((360/float64(g.N))*float64(i)*DegToRad)*g.GetOutsideDia()*50)),
                int((math.Sin((360/float64(g.N))*float64(i)*DegToRad)*g.GetOutsideDia()*50)),
                thinStyle)
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
  cx := int((border + (g1.GetOutsideDia() / 2.0)) * 100)
  cy := height * 50
  f, err := os.Create("TestFile.svg")
  if err != nil {
    fmt.Println("Something failed creating file!")
  }
  canvas := svg.New(f)

  // Setup canvas so that each drawing unit is 0.01mm.
  canvas.StartviewUnit(width, height, "mm", 0, 0, width * 100, height * 100)
  plotGrid(cx, cy, width * 100, height * 100, canvas)
  rot := 0.0
  plotGear(cx, cy, rot, g1, canvas)
  cx = cx + int(centerDist * 100)
  if math.Mod(float64(g2.N), 2) == 0{
    rot = 180.0 / float64(g2.N)
  }
  plotGear(cx, cy, rot, g2, canvas)

  // canvas.Text(width *100 /2, height * 100 /2, "Hello, SVG", "text-anchor:middle;font-size:300;fill:black")
  canvas.End()
  f.Sync()
  f.Close()
}
