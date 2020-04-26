// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package surface

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func Surface() string {
	out := ""
	out += fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az := corner(i+1, j)
			bx, by, bz := corner(i, j)
			cx, cy, cz := corner(i, j+1)
			dx, dy, dz := corner(i+1, j+1)
			for _, num := range []float64{ax, ay, bx, by, cx, cy, dx, dy} {
				if math.IsInf(num, 0) {
					fmt.Println("Skipping infinity")
					continue
				}
				if math.IsNaN(num) {
					fmt.Println("Skipping NAN")
					continue
				}
			}
			avgz := zscale * (az + bz + cz + dz) / 4
			zrange := 3.0
			red := 0
			blue := 0
			h := int(math.Abs(avgz/zrange) * 255)
			if h > 255 {
				h = 255
			}
			if avgz >= 0 {
				red = h
			} else {
				blue = h
			}
			fillColor := fmt.Sprintf("#%02x00%02x", red, blue)
			out += fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='%s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, fillColor)
		}
	}
	out += fmt.Sprintln("</svg>")

	return out
}

func corner(i, j int) (float64, float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z * zscale / height
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	//return math.Sin(r) / r
	return (math.Sin(x) + math.Sin(y)) / r
}

//!-
