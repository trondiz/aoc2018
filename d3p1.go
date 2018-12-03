package main

import (
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	b, err := ioutil.ReadFile("d3.data")
	check(err)
	foo := strings.Split(string(b), "\n")
	foo = foo[:len(foo)-1]
	claims := make([][]string, 0)
	for _, f := range foo {
		r := regexp.MustCompile(`(?P<Id>#\d+) @ (?P<x>\d+),(?P<y>\d+): (?P<sx>\d+)x(?P<sy>\d+)`)
		res := r.FindStringSubmatch(f)
		claims = append(claims, res)
	}

	for len(claims) > 0 {
		var cur []string
		cur, claims = claims[0], claims[1:]
		collcounter := 0
		for _, ba := range claims {
			if compare(cur, ba) {
				//fmt.Println("Collision:", cur, ba)
				collcounter++
				break
			}
		}
		//os.Exit(0)
	}

	log.Println(collcounter)
}
