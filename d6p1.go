package main

import (
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func MyAtoi(v string) int {
	c, _ := strconv.Atoi(v)
	return c
}

func shortestId(coords [][]string, x int, y int) int {
	id := 0
	length := 999
	for ci, c := range coords {
		cx := MyAtoi(c[0])
		cy := MyAtoi(c[1])
		len := int(math.Abs(float64(cx - x + cy - y)))
		if len < length {
			length = len
			id = ci
		}
	}
	return id
}

func main() {
	b, err := ioutil.ReadFile("d6.data")
	check(err)
	foo := strings.Split(string(b), "\n")
	coords := make([][]string, len(foo)-1)
	for i, c := range foo {
		//log.Println(c, i)
		if c != "" {
			coords[i] = strings.Split(c, ", ")
		}
	}

	inflist := make([]int, 0)
	system := make([][]int, 500)
	for s := range system {
		system[s] = make([]int, 500)
	}

	// test and mark infinities
	for x := 0; x < 500; x++ {
		y := 0
		shortest := shortestId(coords, x, y)
		if _, !ok := inflist[shortest] {
			inflist[]
		}
		y = 500
		shortest = shortestId(coords, x, y)
	}

	for i := 0; i < 500; i++ {
		for y := 0; y < 500; y++ {

		}
	}
	//log.Println(system)

}
