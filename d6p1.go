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
	dup := 99
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
				dup = 99
			}
			length = len
			id = ci
		}
	}
	if dup != 99 {
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
	system := make([][]int, 1000)
	for s := range system {
		system[s] = make([]int, 1000)
	}

	// test and mark infinities
	for x := 0; x < 1000; x++ {
		y := 0
		shortest := shortestId(coords, y, x)
		inflist = Add(shortest, inflist)
		y = 1000
		shortest = shortestId(coords, y, x)
		inflist = Add(shortest, inflist)
	}
	for y := 0; y < 1000; y++ {
		x := 0
		shortest := shortestId(coords, y, x)
		inflist = Add(shortest, inflist)
		x = 1000
		shortest = shortestId(coords, y, x)
		inflist = Add(shortest, inflist)
	}
	sort.Ints(inflist)
	counter := make([]int, 100)
	for y := 0; y < 1000; y++ {
		for i := 0; i < 1000; i++ {
			system[y][i] = shortestId(coords, i, y)
			counter[system[y][i]]++
		}
	}

	lis := 0
	for hi, h := range counter {
		found := false
		for _, g := range inflist {
			if hi == g {
				found = true
			}
		}
		if !found {
			if h > lis {
				lis = h
			}
		}
	}
	log.Println("Answer is:", lis)
}
