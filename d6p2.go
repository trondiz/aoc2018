package main

import (
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
	"time"
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

func shortestId(coords [][]int, x int, y int) int {
	length := 0
	for _, c := range coords {
		length += int(math.Abs(float64(x-c[0])) + math.Abs(float64(y-c[1])))
		if length > 10000 {
			return 0
		}
	}
	return 1
}

func main() {
	start := time.Now()
	b, err := ioutil.ReadFile("d6.data")
	check(err)

	foo := strings.Split(string(b), "\n")
	coords := make([][]int, len(foo)-1)
	for i, c := range foo {
		if c != "" {
			strs := strings.Split(c, ", ")
			c := make([]int, len(strs))
			c[0] = MyAtoi(strs[0])
			c[1] = MyAtoi(strs[1])
			coords[i] = c
		}
	}
	counter := 0
	for y := 49; y < 352; y++ {
		for i := 49; i < 352; i++ {
			counter = counter + shortestId(coords, i, y)
		}
	}
	elapsed := time.Since(start)
	log.Println(counter)
	log.Printf("%s", elapsed)
}
