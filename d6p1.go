package main

import (
	"io/ioutil"
	"log"
	"math"
	"sort"
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
	dup := 0
	length := 999
	for ci, c := range coords {
		cx := MyAtoi(c[0])
		cy := MyAtoi(c[1])
		xr := math.Abs(float64(x - cx))
		yr := math.Abs(float64(y - cy))
		len := int(xr + yr)
		if len <= length {
			if len == length {
				dup = id
			} else {
				dup = 0
			}
			length = len
			id = ci
		}
	}
	if dup != 0 {
		return 99
	} else {
		return id
	}
}

func Add(v int, a []int) []int {
	for _, b := range a {
		if v == b {
			return a
		}
	}
	return append(a, v)
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
		inflist = Add(shortest, inflist)
		y = 500
		shortest = shortestId(coords, x, y)
		inflist = Add(shortest, inflist)
	}
	sort.Ints(inflist)
	log.Println(inflist)
	counter := make([]int, 100)
	for i := 0; i < 500; i++ {
		for y := 0; y < 500; y++ {
			system[i][y] = shortestId(coords, i, y)
			counter[system[i][y]]++
		}
	}
	for hi, h := range counter {
		found := false
		for _, g := range inflist {
			if hi == g {
				found = true
			}
		}
		if !found {
			log.Println(hi, h)
		}
	}
	log.Println(counter)

}
