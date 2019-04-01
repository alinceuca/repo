package main

import(
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
	"math/rand"
	"net/http"
	"time"
)

type Point struct {
	coordX float64
	coordY float64
}
var palette = []color.Color{color.White, color.RGBA{255,0,0,255}}
var n int32 = 3
var width float64 = 700
var height float64 = 700

func SierpinskiHandler(w http.ResponseWriter, r *http.Request) {
	var iterations int = 50000
	var nframes = 64
	var delay = 10
	PlotPoints(w, GeneratePoints(width, height, iterations),
		int(width), int(height), nframes, delay, "Draw Triangle")
}

func main() {
	http.HandleFunc("/", SierpinskiHandler)
	s := &http.Server{
		Addr:           ":8081",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("listening on", s.Addr)
	log.Fatal(s.ListenAndServe())
}

func AddLabel(img *image.Paletted, x, y int, label string) {
	col := color.RGBA{200, 100, 0, 255}
	point := fixed.Point26_6{fixed.Int26_6(x * 64), fixed.Int26_6(y * 64)}

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(label)
}

func PlotPoints(out io.Writer, points []Point, w int, h int, nframes int, delay int, label string) {
	steps := make([]int, nframes, nframes)
	for k := 0; k < nframes; k++ {
		steps[k] = int(math.Min(float64(len(points)-1), math.Pow(math.SqrtPhi, float64(k))))
	}

	anim := gif.GIF{LoopCount: nframes}
	ymargin := 10
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, w, h)
		img := image.NewPaletted(rect, palette)
		for j := 0; j <= steps[i]; j++ {
			img.Set(int(points[j].coordX), int(points[j].coordY) + ymargin, color.Black)
		}

		AddLabel(img, w/3+50, h-45, label)

		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim)
}

func GeneratePoints(w float64, h float64, iter int) []Point {
	sin := math.Sin
	cos := math.Cos

	teta := 2*math.Pi/float64(n) + math.Pi/12

	xr := w/2
	yr := h/2

	cx := xr*(cos(teta) + 1)
	cy := yr*(sin(teta) + 1)

	points := []Point{}
	points = append(points, Point{coordX: cx, coordY: cy})

	for i := 0; i < iter; i++ {
		vn := rand.Int31() % n + 1
		vx := xr*(cos(teta*float64(vn)) + 1)
		vy := yr*(sin(teta*float64(vn)) + 1)
		cx += 0.5 * (vx - cx)
		cy += 0.5 * (vy - cy)
		points = append(points, Point{coordX: cx, coordY: cy})
	}

	return points
}