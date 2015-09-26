package plot

import (
  "github.com/ajstarks/svgo"
  "github.com/stuphi/GearGen/gear"
  "os"
  "fmt"
)

func Plot(g gear.Gear) {
  width := int(g.GetOutsideDia() * 1.1)
  height := width
  cx := width * 50
  cy := height * 50
  f, err := os.Create("TestFile.svg")
  if err != nil {
    fmt.Println("Something failed creating file!")
  }
  canvas := svg.New(f)
  dashStyle := "fill:none; stroke-width:20; stroke:black; stroke-dasharray:300,100,100,100"
  solidStyle := "fill:none; stroke-width:30; stroke:black"
  thinStyle := "fill:none; stroke-width:20; stroke:black"
  
  // Stup canvas so that each drawing unit is 0.01mm.
  canvas.StartviewUnit(width, height, "mm", 0, 0, width * 100, height * 100)
  canvas.Circle(cx, cy, int(g.GetOutsideDia() * 100 / 2), solidStyle)
  canvas.Circle(cx, cy, int(g.Pd * 100 / 2), dashStyle)
  canvas.Circle(cx, cy, int(g.GetRootCircleDiameter() * 100 / 2), dashStyle)
  cntrLen := int(g.GetOutsideDia() * 100 / 8)
  canvas.Line(cx - cntrLen, cy, cx + cntrLen, cy, thinStyle)
  canvas.Line(cx, cy - cntrLen, cx, cy + cntrLen, thinStyle)
  for
  canvas.Text(width *100 /2, height * 100 /2, "Hello, SVG", "text-anchor:middle;font-size:300;fill:black")
  canvas.End()
  f.Sync()
  f.Close()
}
