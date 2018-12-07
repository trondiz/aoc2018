package main

import (
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func AddUnique(v rune, a []rune) []rune {
	for _, b := range a {
		if v == b {
			return a
		}
	}
	return append(a, v)
}

type RuneSlice []rune

func (p RuneSlice) Len() int           { return len(p) }
func (p RuneSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p RuneSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func Remove(s []rune, r rune) []rune {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func availJobs(letters map[rune][]rune) []rune {
	startingKeys := make([]rune, 0)
	for k, v := range letters {
		if len(v) == 0 {
			startingKeys = append(startingKeys, k)
		}
	}
	sort.Sort(RuneSlice(startingKeys))
	//log.Println(startingKeys)
	return startingKeys
}

type worker struct {
	part     rune
	timeleft int
}

func main() {
	b, err := ioutil.ReadFile("d7.data")
	check(err)
	foo := strings.Split(string(b), "\n")
	foo = foo[:len(foo)-1]

	letters := make(map[rune][]rune)
	for _, f := range foo {
		r := regexp.MustCompile(`Step (?P<first>[A-Z]) must be finished before step (?P<then>[A-Z]) can begin.`)
		res := r.FindStringSubmatch(f)
		res = res[1:]
		a := []rune(res[0])
		b := []rune(res[1])
		_, ok := letters[a[0]]
		if !ok {
			letters[a[0]] = make([]rune, 0)
		}
		letters[b[0]] = append(letters[b[0]], a[0])
	}

	result := 0
	workers := make([]worker, 5)
	for len(letters) >= 0 {
		for wi, w := range workers {
			if w.part != 0 {
				workers[wi].timeleft--
			}
			if w.timeleft == 0 && w.part != 0 {
				for li, _ := range letters {
					letters[li] = Remove(letters[li], w.part)
				}
				if len(letters) == 0 {
					log.Println(result)
					os.Exit(0)
				}
				workers[wi].part = 0
			}
		}
		sKey := availJobs(letters)
		for _, s := range sKey {
			for wi, w := range workers {
				if w.part == 0 {
					workers[wi].part = s
					delete(letters, s)
					workers[wi].timeleft = 60 + (int(s) - 65)
					break
				}
			}
		}
		result++
	}
}
