package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func compare(a []string, b []string) bool {
	count := 0
	for i := 0; i < 26; i++ {
		if a[i] == b[i] {
			count++
		}
	}
	if count == 25 {
		return true
	} else {
		return false
	}
}

func main() {
	b, err := ioutil.ReadFile("d2.data")
	check(err)
	foo := strings.Split(string(b), "\n")
	data := make([][]string, 0)
	for _, f := range foo {
		ltrs := strings.Split(string(f), "")
		if len(ltrs) > 10 {
			data = append(data, ltrs)
		}
	}
	for true {
		var cur []string
		cur, data = data[0], data[1:]
		for _, ba := range data {
			if compare(cur, ba) {
				fmt.Println("FOUND:", cur, ba)
				os.Exit(0)
			}
		}
	}
}
