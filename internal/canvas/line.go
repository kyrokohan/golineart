package canvas

import (
	"image"
	"image/color"
	"image/draw"
	"math"
	"math/rand/v2"
	"runtime"
	"sync"
	"time"

	"github.com/kyrokohan/golineart/internal/img"
	"github.com/kyrokohan/golineart/internal/types"
	"github.com/kyrokohan/golineart/internal/utils"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func bresenham(line *types.Line) {
	x0, y0, x1, y1 := line.X0, line.Y0, line.X1, line.Y1

	dx := abs(x1 - x0)

	sx := -1
	if x0 < x1 {
		sx = 1
	}

	dy := -abs(y1 - y0)
	sy := -1
	if y0 < y1 {
		sy = 1
	}

	err := dx + dy

	for {
		line.Pixels = append(line.Pixels, image.Point{x0, y0})
		e2 := 2 * err

		if e2 >= dy {
			if x0 == x1 {
				break
			}
			err = err + dy
			x0 = x0 + sx
		}

		if e2 <= dx {
			if y0 == y1 {
				break
			}
			err = err + dx
			y0 = y0 + sy
		}
	}
}

func GenerateRandomLineCoordinates(r *rand.Rand, b image.Rectangle) types.Line {
	var line types.Line
	// get widths
	dx := b.Dx()
	dy := b.Dy()

	// get random edges while making sure same edge never gets chosen
	e0 := r.IntN(4)
	e1 := r.IntN(3)
	if e1 >= e0 {
		e1++
	}

	// helper function to get point on an edge
	pointOnEdge := func(edge int) (int, int) {
		switch edge {
		case 0: // top
			return b.Min.X + r.IntN(dx), b.Min.Y
		case 1: // right
			return b.Max.X - 1, b.Min.Y + r.IntN(dy)
		case 2: // bottom
			return b.Min.X + r.IntN(dx), b.Max.Y - 1
		default: // left
			return b.Min.X, b.Min.Y + r.IntN(dy)
		}
	}

	line.X0, line.Y0 = pointOnEdge(e0)
	line.X1, line.Y1 = pointOnEdge(e1)

	bresenham(&line)

	return line
}

func DrawLineOld(dst *image.RGBA, line types.Line, c color.RGBA, alpha uint) {
	src := image.NewUniform(c)
	mask := image.NewUniform(color.Alpha{uint8(alpha)})

	for i := range len(line.Pixels) {
		p1 := line.Pixels[i]
		p2 := line.Pixels[i]
		p2.X += 1
		p2.Y += 1
		draw.DrawMask(dst, image.Rectangle{p1, p2}, src, image.Point{}, mask, image.Point{}, draw.Over)
	}
}

func DrawLine(dst *image.RGBA, line types.Line, c color.RGBA, alpha uint) {
	for i := range len(line.Pixels) {
		p := line.Pixels[i]
		idx := dst.PixOffset(p.X, p.Y)

		r1, g1, b1 := utils.ApplyAlpha(dst, c, alpha, idx)

		dst.Pix[idx] = uint8(r1)
		dst.Pix[idx+1] = uint8(g1)
		dst.Pix[idx+2] = uint8(b1)
	}
}

func DrawBestOfNLines(dst, src *image.RGBA, n int, c color.RGBA, alpha uint, workers int) uint64 {
	if n <= 0 {
		return 0
	}

	if workers <= 0 {
		workers = max(runtime.GOMAXPROCS(0), 1)
	}

	type result struct {
		line  types.Line
		after uint64
		delta int64
		err   error
	}

	jobs := make(chan int)
	results := make(chan result, workers)

	var wg sync.WaitGroup

	for w := 0; w < workers; w++ {
		wg.Add(1)
		rnd := rand.New(rand.NewPCG(uint64(w), uint64(time.Now().UnixNano())))

		go func(r *rand.Rand) {
			defer wg.Done()

			for range jobs {
				line := GenerateRandomLineCoordinates(r, dst.Bounds())

				base, err := img.LineDiff(dst, src, line, c, alpha, false)
				if err != nil {
					results <- result{err: err}
					continue
				}

				after, err := img.LineDiff(dst, src, line, c, alpha, true)
				if err != nil {
					results <- result{err: err}
					continue
				}

				results <- result{line, after, int64(base) - int64(after), err}
			}
		}(rnd)
	}

	go func() {
		for i := range n {
			jobs <- i
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	best := result{delta: math.MinInt64}
	for res := range results {
		if res.err != nil {
			continue
		}
		if res.delta > best.delta {
			best = res
		}
	}

	if best.delta > 0 {
		DrawLine(dst, best.line, c, alpha)
		return best.after
	}

	return 0
}
