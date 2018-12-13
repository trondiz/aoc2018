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

type cart struct {
	x                  int
	y                  int
	mydirection        rune // <>v^
	lastinterdirection int  // 0 1 2, 0 == left, 1 == straight, 2 == rigth
}

func main() {
	b, err := ioutil.ReadFile("d13.data")
	check(err)
	foo := strings.Split(string(b), "\n")
	foo = foo[:len(foo)-1]
	chart := make([][]rune, len(foo))
	for fi, f := range foo {
		chart[fi] = []rune(f)
	}
	carts := make([]cart, 0)
	for xi, c := range chart {
		for yi, r := range c {
			if r == '<' || r == '>' || r == 'v' || r == '^' {
				carts = append(carts, cart{x: xi, y: yi, mydirection: r})
			}
		}
	}
	// ticks
	for ci, c := range carts {
		switch chart[c.x][c.y] {
		case '|':
			log.Println("continue same dir")
			if c.mydirection == '^' {
				chart[c.x][c.y].y--
			}
			if c.mydirection == 'v' {
				chart[c.x][c.y].y++
			}

		case '+':
			log.Println("turn")
		case '/':
			log.Println("cart turns right")
		case '\':
			log.Println("cart turns left")
		case '-':
			log.Println("Continue same dir")
		}
		// Collision detection, have this cart collided

		for ai, a := range carts {
			if c.x == a.x && c.y == a.y {
				log.Println("Collision at:", c.x, c.y)
				os.Exit(0)
			}
		}
	}
	log.Println(carts)
}
