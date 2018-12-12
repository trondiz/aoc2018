package main

import (
	"log"
	"strconv"
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

func getHundreds(a int) int {
	stringTarget := strconv.Itoa(a)
	len := len(stringTarget)
	if len < 3 {
		return 0
	}
	b, err := strconv.Atoi(string(stringTarget[len-3]))
	check(err)
	return b
}

func getPower(x, y int) int {
	pid := x + 10
	pwr := pid * y
	pwr += input
	pwr = pwr * pid
	pwr = getHundreds(pwr)
	pwr -= 5
	return pwr
}

func checkSquare(cx, cy, size int, g [][]int) int {
	tpwr := 0
	if cx+size > 300 {
		return 0
	}
	if cy+size > 300 {
		return 0
	}
	for x := cx; x < size+cx; x++ {
		for y := cy; y < size+cy; y++ {
			tpwr += g[x][y]
		}
	}
	return tpwr
}

var input int

func main() {
	input = 1309

	grid := make([][]int, 301)
	for gi, _ := range grid {
		grid[gi] = make([]int, 301)
	}
	tpwr := 0
	tpwrid := make([]int, 3)

	for x := 1; x <= 300; x++ {
		for y := 1; y <= 300; y++ {
			grid[x][y] = getPower(x, y)
		}
	}
	for s := 1; s <= 300; s++ {
		for x := 1; x <= 300; x++ {
			for y := 1; y <= 300; y++ {
				pwr := checkSquare(x, y, s, grid)
				if pwr > tpwr {
					tpwr = pwr
					tpwrid[0] = x
					tpwrid[1] = y
					tpwrid[2] = s
				}
				//log.Println(s)
			}
		}
	}

	log.Println(tpwr, tpwrid)

}
