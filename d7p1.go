package main

import (
	"io/ioutil"
	"log"
	"regexp"
	"sort"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func AddUnique(v string, a []string) []string {
	for _, b := range a {
		if v == b {
			return a
		}
	}
	return append(a, v)
}
func Remove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func findStartingKey(letters map[string][]string) string {
	startingKeys := make([]string, 0)
	for k, v := range letters {
		if len(v) == 0 {
			startingKeys = append(startingKeys, k)
		}
	}
	sort.Strings(startingKeys)
	return startingKeys[0]
}

func main() {
	b, err := ioutil.ReadFile("d7.data")
	check(err)
	foo := strings.Split(string(b), "\n")
	foo = foo[:len(foo)-1]

	letters := make(map[string][]string)
	for _, f := range foo {
		r := regexp.MustCompile(`Step (?P<first>[A-Z]) must be finished before step (?P<then>[A-Z]) can begin.`)
		res := r.FindStringSubmatch(f)
		res = res[1:]
		_, ok := letters[res[0]]
		if !ok {
			letters[res[0]] = make([]string, 0)
		}
		letters[res[1]] = append(letters[res[1]], res[0])
	}

	result := ""
	for len(letters) > 0 {
		sKey := findStartingKey(letters)
		result += sKey
		//log.Println(letters, sKey)
		// Delete key from entries that are blocked by this
		for li, _ := range letters {
			letters[li] = Remove(letters[li], sKey)
		}
		delete(letters, sKey)
	}
	log.Println(result)
}
