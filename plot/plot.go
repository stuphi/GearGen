package plot

import (
  "github.com/ajstarks/svgo"
  "github.com/stuphi/GearGen/gear"
  "os"
)

func Plot(g gear.Gear) {
  width := int(g.GetOutsideDia() * 1.1)
  height := width
  canvas := svg.New(os.Stdout)
  canvas.StartviewUnit(width, height, "mm", 0, 0, width * 100, height * 100)
  canvas.Circle(width *100 / 2, height * 100 / 2, int(g.GetOutsideDia() * 100 / 2))
  canvas.Text(width/2, height/2, "Hello, SVG", "text-anchor:middle;font-size:30px;fill:white")
  canvas.End()
}
