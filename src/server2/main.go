package main

import (
	"fmt"
	"image"
	"image/color/palette"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		count++
		mu.Unlock()
		lissajous(w)
	})
	http.HandleFunc("/counter", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Count: %d\n", count)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation rames
		delay   = 8     // delay between frames in 10ms units
	)
	rand.Seed(time.Now().UTC().UnixNano())
	/*
		var palette []color.Color
		var red, green, blue uint8
		const halfHeight = 127
		for t := 0; t < 256; t++ {
			fmt.Fprintf(os.Stderr, "t: %d, t/halfHeight: %f\n", t, float64(t)/halfHeight)
			// math.Sin(float64(t/halfHeight)))
			red = uint8(halfHeight + halfHeight*math.Sin(float64(t)/halfHeight))
			green = uint8(halfHeight + halfHeight*math.Sin(float64(t)/halfHeight+5))
			blue = uint8(halfHeight + halfHeight*math.Sin(float64(t)/halfHeight+10))
			palette = append(palette, color.RGBA{red, green, blue, 0xff})
		}
	*/
	pal := palette.Plan9
	freq := rand.Float64() * 3.0 // relative frequency of y oscillations
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, pal)
		color := uint8(1)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), color%255)
			color++
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // ignoring encoding errors
}
