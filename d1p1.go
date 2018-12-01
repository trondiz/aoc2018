package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	b, err := ioutil.ReadFile("d1.data")
	check(err)
	freqs := strings.Split(string(b), "\n")
	sum := 0
	var seen []int
	seen = append(seen, 0)
	for _, f := range freqs {
		if f != "" {
			i, err := strconv.Atoi(f)
			check(err)
			//log.Println(i)
			sum += i
		}
	}
	log.Println(sum)
}
