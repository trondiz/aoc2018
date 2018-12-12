package main

import (
	"io/ioutil"
	"log"
	"os"
	"regexp"
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

func main() {
	b, err := ioutil.ReadFile("d12.data")
	check(err)
	foo := strings.Split(string(b), "\n")
	foo = foo[:len(foo)-1]
	r := regexp.MustCompile(`initial state: ([\.\#]*)`)
	in := r.FindStringSubmatch(foo[0])
	state := strings.Split(in[1], "")
	rules := make([][]string, 0)
	for _, f := range foo {
		r := regexp.MustCompile(`([\.\#]{5}) => (#?.?)`)
		res := r.FindStringSubmatch(f)
		if len(res) > 2 {
			matcher := res[1:]
			rules = append(rules, matcher)
		}
	}
	//nullpos := 0
	// Add padding
	nullpos := 122
	for i := 0; i < nullpos; i++ {
		state = append(state, ".")
		state = append(state, ".")
		copy(state[1:], state[0:])
		state[0] = "."
	}
	//diff := 0
	presult := 0
	diff := 0
	prediff := 0
	i := 0
	for {
		tmp := make([]string, len(state))
		copy(tmp, state)
		for si, _ := range state {
			if si < 2 {
				continue
			}
			if si >= len(state)-2 {
				break
			}
			try := strings.Join(state[si-2:si+3], "")
			match := false
			for _, r := range rules {

				if string(try) == r[0] {
					tmp[si] = r[1]
					match = true
				}
			}
			if !match {
				tmp[si] = "."
			}
		}
		state = tmp
		//log.Println(strings.Join(state, ""))
		result := 0
		for si, s := range state {
			if s == "#" {
				result += si - nullpos
			}
		}
		if result-presult == diff {
			if diff == prediff {
				log.Print(result + (diff * (50000000000 - (i + 1))))
				os.Exit(0)
			} else {
				prediff = diff
			}
		} else {
			diff = result - presult
		}
		presult = result
		i++
	}

}
