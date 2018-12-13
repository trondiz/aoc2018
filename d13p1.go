package main

import (
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type cart struct {
	x                  int
	y                  int
	mydirection        rune // <>v^
	lastinterdirection int  // 0 1 2, 0 == left, 1 == straight, 2 == rigth
}
type cartSlice []cart

func (c cartSlice) Len() int           { return len(c) }
func (c cartSlice) Less(i, j int) bool { return c[i].y < c[j].y }
func (c cartSlice) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }

func main() {
	t_input := time.Now()
	b, err := ioutil.ReadFile("d13.data")
	check(err)

	foo := strings.Split(string(b), "\n")
	foo = foo[:len(foo)-1]
	chart := make([][]rune, len(foo))
	for fi, f := range foo {
		chart[fi] = []rune(f)
	}
	carts := make([]cart, 0)
	for yi, c := range chart {
		for xi, r := range c {
			if r == '<' || r == '>' || r == 'v' || r == '^' {
				carts = append(carts, cart{x: xi, y: yi, mydirection: r})
				if r == '<' || r == '>' {
					chart[yi][xi] = '-'
				}
				if r == '^' || r == 'v' {
					chart[yi][xi] = '|'
				}
			}

		}
	}
	input_t_elapsed := time.Since(t_input)
	log.Println("Input:", input_t_elapsed)
	t_input = time.Now()
	// ticks
	//for k := 0; k < 10000; k++ {
	firstcoll := false
	for {
		sort.Sort(cartSlice(carts))
		deletionlist := make([]int, 0)
		for ci, c := range carts {
			switch chart[c.y][c.x] {
			case '|':
				if c.mydirection == '^' {
					carts[ci].y--
				}
				if c.mydirection == 'v' {
					carts[ci].y++
				}
			case '+':
				switch c.mydirection {
				case '<':
					switch c.lastinterdirection {
					case 0:
						carts[ci].mydirection = 'v'
						carts[ci].y++
						carts[ci].lastinterdirection++
					case 1:
						carts[ci].x--
						carts[ci].lastinterdirection++
					case 2:
						carts[ci].mydirection = '^'
						carts[ci].y--
						carts[ci].lastinterdirection = 0
					}
				case '^':
					switch c.lastinterdirection {
					case 0:
						carts[ci].mydirection = '<'
						carts[ci].x--
						carts[ci].lastinterdirection++
					case 1:
						carts[ci].y--
						carts[ci].lastinterdirection++
					case 2:
						carts[ci].mydirection = '>'
						carts[ci].x++
						carts[ci].lastinterdirection = 0
					}
				case '>':
					switch c.lastinterdirection {
					case 0:
						carts[ci].mydirection = '^'
						carts[ci].y--
						carts[ci].lastinterdirection++
					case 1:
						carts[ci].x++
						carts[ci].lastinterdirection++
					case 2:
						carts[ci].mydirection = 'v'
						carts[ci].y++
						carts[ci].lastinterdirection = 0
					}
				case 'v':
					switch c.lastinterdirection {
					case 0:
						carts[ci].mydirection = '>'
						carts[ci].x++
						carts[ci].lastinterdirection++
					case 1:
						carts[ci].y++
						carts[ci].lastinterdirection++
					case 2:
						carts[ci].mydirection = '<'
						carts[ci].x--
						carts[ci].lastinterdirection = 0
					}
				}
			case '/':
				switch c.mydirection {
				case '<':
					carts[ci].mydirection = 'v'
					carts[ci].y++
				case '^':
					carts[ci].mydirection = '>'
					carts[ci].x++
				case '>':
					carts[ci].mydirection = '^'
					carts[ci].y--
				case 'v':
					carts[ci].mydirection = '<'
					carts[ci].x--
				}
			case '\\':
				switch c.mydirection {
				case '<':
					carts[ci].mydirection = '^'
					carts[ci].y--
				case '^':
					carts[ci].mydirection = '<'
					carts[ci].x--
				case '>':
					carts[ci].mydirection = 'v'
					carts[ci].y++
				case 'v':
					carts[ci].mydirection = '>'
					carts[ci].x++
				}
			case '-':
				if c.mydirection == '<' {
					carts[ci].x--
				}
				if c.mydirection == '>' {
					carts[ci].x++
				}
			}
			// Collision detection, have this cart collided
			for ai, a := range carts {
				if ci != ai {
					if c.x == a.x && c.y == a.y {
						if !firstcoll {
							input_t_elapsed := time.Since(t_input)
							log.Println("P1 Collision at:", c.x, c.y, "between", c, a, "Finished", input_t_elapsed, "after input")
							t_input = time.Now()
							firstcoll = true
						}
						deletionlist = append(deletionlist, ai)
						deletionlist = append(deletionlist, ci)
					}
				}
			}
		}

		// Deletion
		if len(deletionlist) > 0 {
			ox := carts[deletionlist[0]].x
			oy := carts[deletionlist[0]].y
			tx := carts[deletionlist[1]].x
			ty := carts[deletionlist[1]].y
			for ci, c := range carts {
				if c.x == ox && c.y == oy {
					carts = append(carts[:ci], carts[ci+1:]...)
				}
			}
			for ci, c := range carts {
				if c.x == tx && c.y == ty {
					carts = append(carts[:ci], carts[ci+1:]...)
				}
			}
			deletionlist = deletionlist[:0]
		}

		if len(carts) == 1 {
			input_t_elapsed := time.Since(t_input)
			log.Println("last cart is at", carts[0].x, carts[0].y, "Time:", input_t_elapsed)
			os.Exit(0)
		}

	}

}
