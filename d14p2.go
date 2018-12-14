package main

import (
	"log"
	"os"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func diff(a []byte, b []byte) bool {
	for ci, c := range a {
		if (b[ci] ^ c) > 0 {
			return true
		}
	}
	return false
}

func main() {
	t_input := time.Now()

	input := []byte{7, 0, 4, 3, 2, 1}
	scoreboard := []byte{3, 7}
	e1pos := 0
	e2pos := 1
	for {
		nr := scoreboard[e1pos] + scoreboard[e2pos]
		if nr > 9 {
			e1r := (nr / 10) % 10
			e2r := nr % 10
			scoreboard = append(scoreboard, e1r)
			if len(scoreboard) > len(input) {
				diff := false
				for ini, in := range input {
					if in != scoreboard[len(scoreboard)-len(input)+ini] {
						diff = true
						break
					}
				}
				if !diff {
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
		e1pos = (e1pos + 1 + int(scoreboard[e1pos])) % len(scoreboard)
		e2pos = (e2pos + 1 + int(scoreboard[e2pos])) % len(scoreboard)
		if len(scoreboard) > len(input) {
			diff := false
			for ini, in := range input {
				if in != scoreboard[len(scoreboard)-len(input)+ini] {
					diff = true
					break
				}
			}
			if !diff {
				input_t_elapsed := time.Since(t_input)
				log.Println(input_t_elapsed)
				log.Println(len(scoreboard) - len(input))
				os.Exit(0)
			}
		}
	}

}
