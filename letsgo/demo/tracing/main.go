// This program generates Mandelbrot fractals using different
// concurrency patterns.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"runtime/pprof"
	"sync"
	"time"
)

const (
	output = "out.png"
	width  = 2048
	height = 2048
)

func main() {
	defer trackTime(time.Now(), "MAIN")
	// uncomment these lines to generate traces into stderr.
	pprof.StartCPUProfile(os.Stderr)
	defer pprof.StopCPUProfile()

	// uncomment these lines to generate profile data into stderr.
	//pprof.StartCPUProfile(os.Stderr)
	//defer pprof.StopCPUProfile()

	f, err := os.Create(output)
	if err != nil {
		log.Fatal(err)
	}

	img := createWorkers(width, height, true)
	// img := createSeq(width, height)

	if err = png.Encode(f, img); err != nil {
		log.Fatal(err)
	}
}

// track the time for execution of a function
func trackTime(s time.Time, msg string) {
	elapsed := time.Since(s)
	fmt.Println(msg, ":", elapsed)
}

// createSeq fills one pixel at a time.
func createSeq(width, height int) image.Image {
	m := image.NewGray(image.Rect(0, 0, width, height))
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			m.Set(i, j, pixel(i, j, width, height))
		}
	}
	return m
}

// createPixel creates one goroutine per pixel.
func createPixel(width, height int) image.Image {
	m := image.NewGray(image.Rect(0, 0, width, height))
	var w sync.WaitGroup
	w.Add(width * height)
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			go func(i, j int) {
				m.Set(i, j, pixel(i, j, width, height))
				w.Done()
			}(i, j)
		}
	}
	w.Wait()
	return m
}

// createRow creates one goroutine per row.
func createRow(width, height int) image.Image {
	m := image.NewGray(image.Rect(0, 0, width, height))
	var w sync.WaitGroup
	w.Add(width)
	for i := 0; i < width; i++ {
		go func(i int) {
			for j := 0; j < height; j++ {
				m.Set(i, j, pixel(i, j, width, height))
			}
			w.Done()
		}(i)
	}
	w.Wait()
	return m
}

// createWorkers creates 8 workers and uses a channel to pass each pixel.
func createWorkers(width, height int, buffered bool) image.Image {
	m := image.NewGray(image.Rect(0, 0, width, height))

	type px struct{ x, y int }

	cap := 0
	if buffered {
		cap = width * height
	}
	c := make(chan px, cap)

	var w sync.WaitGroup
	w.Add(8)
	for i := 0; i < 8; i++ {
		go func() {
			for px := range c {
				m.Set(px.x, px.y, pixel(px.x, px.y, width, height))
			}
			w.Done()
		}()
	}

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			c <- px{i, j}
		}
	}
	close(c)
	w.Wait()
	return m

}

// createRowWorkers creates 8 workers and uses a channel to pass each row.
func createRowWorkers(width, height int, buffered bool) image.Image {
	m := image.NewGray(image.Rect(0, 0, width, height))

	cap := 0
	if buffered {
		cap = width
	}
	c := make(chan int, cap)

	var w sync.WaitGroup
	w.Add(8)
	for i := 0; i < 8; i++ {
		go func() {
			for i := range c {
				for j := 0; j < height; j++ {
					m.Set(i, j, pixel(i, j, width, height))
				}
			}
			w.Done()
		}()
	}

	for i := 0; i < width; i++ {
		c <- i
	}

	close(c)
	w.Wait()
	return m
}

// pixel returns the color of a Mandelbrot fractal at the given point.
func pixel(i, j, width, height int) color.Color {
	// Play with this constant to increase the complexity of the fractal.
	// In the justforfunc.com video this was set to 4.
	const complexity = 1024

	xi := norm(i, width, -1.0, 2)
	yi := norm(j, height, -1, 1)

	const maxI = 1000
	x, y := 0., 0.

	for i := 0; (x*x+y*y < complexity) && i < maxI; i++ {
		x, y = x*x-y*y+xi, 2*x*y+yi
	}

	return color.Gray{uint8(x)}
}

func norm(x, total int, min, max float64) float64 {
	return (max-min)*float64(x)/float64(total) - max
}
