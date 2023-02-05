// Lissajous generate Lissajous random GIF animations
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first palette color
	blackIndex = 1 // second palette color
)

// func main() {
// 	rand.Seed(time.Now().UTC().UnixNano()) // TODO: what happen here?
// 	lissajous(os.Stdout)
// }

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // x oscilation
		res     = 0.001 // angle
		size    = 100   // canvas image
		nframes = 64    // number of frames
		delay   = 8     // time between frames in unit of 10ms
	)

	freq := rand.Float64() * 3.0 // frequency relative of oscilation y
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // diff phases

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim) // NOTE: ignored err
}
