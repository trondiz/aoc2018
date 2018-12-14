package main

import (
	"log"
	"os"
	"strconv"
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
	input := []int{9, 2, 0, 8, 3, 1}
	scoreboard := []int{3, 7}
	e1pos := 0
	e2pos := 1
	for {
		nr := scoreboard[e1pos] + scoreboard[e2pos]
		if nr > 9 {
			e1r := (nr / 10) % 10
			e2r := nr % 10
			scoreboard = append(scoreboard, e1r)
			scoreboard = append(scoreboard, e2r)
		} else {
			scoreboard = append(scoreboard, nr)
		}
		e1pos = (e1pos + 1 + scoreboard[e1pos]) % len(scoreboard)
		e2pos = (e2pos + 1 + scoreboard[e2pos]) % len(scoreboard)
		if len(scoreboard) > len(input) {
			cmp := scoreboard[len(scoreboard)-len(input):]
			diff := false
			for ci, c := range cmp {
				if input[ci] != c {
					diff = true
					break
				}
			}
			if !diff {
				log.Println(len(scoreboard) - len(input))
				os.Exit(0)
			}
		}
	}

}
