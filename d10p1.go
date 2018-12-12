package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type point struct {
	x  int
	y  int
	xv int
	yv int
}

func MyAtoi(v string) int {
	c, _ := strconv.Atoi(v)
	return c
}

func frameSize(a []point) []int {
	minx := 99999
	maxx := -99999
	miny := 99999
	maxy := -99999

	for _, p := range a {
		if p.x < minx {
			minx = p.x
		}
		if p.x > maxx {
			maxx = p.x
		}
		if p.y > maxy {
			maxy = p.y
		}
		if p.y < miny {
			miny = p.y
		}
	}
	xy := make([]int, 2)
	xy[0] = maxx - minx
	xy[1] = maxy - miny
	return xy
}

func main() {
	t_input := time.Now()
	b, err := ioutil.ReadFile("d10.data")
	check(err)
	foo := strings.Split(string(b), "\n")
	foo = foo[:len(foo)-1]

	points := make([]point, 0)
	for _, f := range foo {
		r := regexp.MustCompile(`position=<\s*(-?\d+),\s*(-?\d+)>\s*velocity=<\s*(-?\d+),\s*(-?\d+)>`)
		res := r.FindStringSubmatch(f)
		p := point{x: MyAtoi(res[1]), y: MyAtoi(res[2]), xv: MyAtoi(res[3]), yv: MyAtoi(res[4])}
		points = append(points, p)
	}
	input_t_elapsed := time.Since(t_input)
	t_total := time.Now()
	size := 999999999
	counter := 0
	for {
		for pi, p := range points {
			points[pi].x = p.x + p.xv
			points[pi].y = p.y + p.yv
		}
		counter++
		f := frameSize(points)
		newsize := f[0]
		if newsize < size {
			size = newsize
		} else {
			// We're there, time to render
			// step on second back
			for pi, p := range points {
				points[pi].x = p.x - p.xv
				points[pi].y = p.y - p.yv
			}
			counter--
			offsetx := 500
			offsety := 500
			for _, p := range points {
				if p.x < offsetx {
					offsetx = p.x
				}
				if p.y < offsety {
					offsety = p.y
				}
			}
			offsetx = int(math.Abs(float64(offsetx)))
			offsety = int(math.Abs(float64(offsety)))
			sx := 0
			sy := 0
			for pi, p := range points {
				points[pi].x = p.x - offsetx
				points[pi].y = p.y - offsety
			}
			for _, p := range points {
				if p.x > sx {
					sx = p.x + 1
				}
				if p.y > sy {
					sy = p.y + 1
				}
			}
			amap := make([][]string, sy)
			for ami, _ := range amap {
				amap[ami] = make([]string, sx)
			}
			for y := 0; y < sy; y++ {
				for x := 0; x < sx; x++ {
					amap[y][x] = "."
				}
			}
			for _, p := range points {
				amap[p.y][p.x] = "#"
			}
			for _, am := range amap {
				fmt.Println(am)
			}
			break
		}
	}
	log.Println("Runtime", counter)
	totaltime := time.Since(t_total)
	log.Println("Spent on input:", input_t_elapsed)
	log.Println("Spent on calc+draw:", totaltime)

}
