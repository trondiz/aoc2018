package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	b, err := ioutil.ReadFile("d2.data")
	check(err)
	foo := strings.Split(string(b), "\n")
	twocnt := 0
	threecnt := 0
	for _, f := range foo {
		ltrs := strings.Split(string(f), "")
		m := make(map[string]int)
		for _, l := range ltrs {
			m[l] = m[l] + 1
		}
		twoseen := false
		threeseen := false
		for _, c := range m {
			if c == 2 && !twoseen {
				twocnt += 1
				twoseen = true
			}
			if c == 3 && !threeseen {
				threecnt += 1
				threeseen = true
			}
		}

	}
	log.Println(twocnt * threecnt)
}
