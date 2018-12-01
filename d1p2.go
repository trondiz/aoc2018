package main

import (
	"io/ioutil"
	"log"
	"os"
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
	for true {
		for _, f := range freqs {
			if f != "" {
				i, err := strconv.Atoi(f)
				check(err)
				sum += i
				for _, e := range seen {
					if e == sum {
						log.Println("Seen before:", e)
						os.Exit(0)
					}
				}
				seen = append(seen, sum)
			}
		}
	}
}
