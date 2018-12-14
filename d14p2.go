package main

import (
	"log"
	"os"
	"strconv"
	"time"
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
func diff(a []int, b []int) bool {
	for ci, c := range a {
		if b[ci] != c {
			return true
		}
	}
	return false
}

func main() {
	t_input := time.Now()

	input := []int{7, 0, 4, 3, 2, 1}
	scoreboard := []int{3, 7}
	e1pos := 0
	e2pos := 1
	for {
		nr := scoreboard[e1pos] + scoreboard[e2pos]
		if nr > 9 {
			e1r := (nr / 10) % 10
			e2r := nr % 10
			scoreboard = append(scoreboard, e1r)
			if len(scoreboard) > len(input) {
				cmp := scoreboard[len(scoreboard)-len(input):]
				diff := diff(cmp, input)
				if !diff {
					log.Println(cmp, input)
					log.Println(len(scoreboard) - len(input))
					input_t_elapsed := time.Since(t_input)
					log.Println(input_t_elapsed)
					os.Exit(0)
				}
			}
			scoreboard = append(scoreboard, e2r)
		} else {
			scoreboard = append(scoreboard, nr)
		}
		e1pos = (e1pos + 1 + scoreboard[e1pos]) % len(scoreboard)
		e2pos = (e2pos + 1 + scoreboard[e2pos]) % len(scoreboard)
		if len(scoreboard) > len(input) {
			cmp := scoreboard[len(scoreboard)-len(input):]
			diff := diff(cmp, input)
			if !diff {
				input_t_elapsed := time.Since(t_input)
				log.Println(input_t_elapsed)
				log.Println(cmp, input)
				log.Println(len(scoreboard) - len(input))
				os.Exit(0)
			}
		}
	}

}
