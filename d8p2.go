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

func MyAtoi(v string) int {
	c, _ := strconv.Atoi(v)
	return c
}

type node struct {
	numChild    int
	numMeta     int
	Child       []node
	metaEntries []int
}

var result int
var nds []int

func newNode() node {
	n := node{}
	n.numChild, nds = nds[0], nds[1:]
	n.numMeta, nds = nds[0], nds[1:]
	for i := 0; i < n.numChild; i++ {
		n.Child = append(n.Child, newNode())
	}
	for i := 0; i < n.numMeta; i++ {
		var a int
		a, nds = nds[0], nds[1:]
		n.metaEntries = append(n.metaEntries, a)
	}
	return n
}

func getValue(node node) int {
	if node.numChild == 0 {
		r := 0
		for _, m := range node.metaEntries {
			r += m
		}
		return r
	} else {
		s := 0
		for _, m := range node.metaEntries {
			if m > len(node.Child) {
				continue
			}
			s += getValue(node.Child[m-1])
		}
		return s
	}
}

func main() {
	result = 0
	b, err := ioutil.ReadFile("d8.data")
	check(err)
	foo := strings.Split(string(b), "\n")
	bar := foo[0]
	numbers := strings.Split(bar, " ")
	nds = make([]int, 0)
	for _, n := range numbers {
		nds = append(nds, MyAtoi(n))
	}
	n := newNode()
	log.Println(getValue(n))

}
