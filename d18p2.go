package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Counter(chrt [][]rune, x, y int, tp rune) int {
	cntr := 0
	for yr := 0 + y; yr < 3+y; yr++ {
		for xr := 0 + x; xr < 3+x; xr++ {
			if yr < 0 || yr >= len(chrt) {
				continue
			}
			if xr < 0 || xr >= len(chrt) {
				continue
			}
			if xr == 1+x && yr == 1+y {
				continue
			}
			if chrt[yr][xr] == tp {
				cntr++
			}
			if cntr == 3 {
				return 3
			}
		}
	}
	return cntr
}

func AddUnique(v string, a []string) []string {
	for i := len(a) - 1; i >= 0; i-- {
		if v == a[i] {
			return a
		}
	}
	return append(a, v)
}

type change struct {
	x int
	y int
	r rune
}

func main() {
	t_input := time.Now()
	b, err := ioutil.ReadFile("d18.data")
	check(err)

	foo := strings.Split(string(b), "\n")
	foo = foo[:len(foo)-1]
	chart := make([][]rune, len(foo))
	duplist := make([]string, 0)
	realduplist := make([]string, 0)
	changeLog := make([]change, 0)
	for fi, f := range foo {
		chart[fi] = []rune(f)
	}
	input_t_elapsed := time.Since(t_input)
	log.Println("Input:", input_t_elapsed)
	t_input = time.Now()
	duplen := 0

	for j := 0; j < 1000000000; j++ {
		for oi, o := range chart {
			for ooi, oo := range o {
				switch oo {
				case '.':
					treecnt := Counter(chart, ooi-1, oi-1, '|')
					if treecnt >= 3 {
						changeLog = append(changeLog, change{x: ooi, y: oi, r: '|'})
					}
				case '|':
					lmbcnt := Counter(chart, ooi-1, oi-1, '#')
					if lmbcnt >= 3 {
						changeLog = append(changeLog, change{x: ooi, y: oi, r: '#'})
					}
				case '#':
					lmbcnt := Counter(chart, ooi-1, oi-1, '#')
					if lmbcnt < 1 {
						changeLog = append(changeLog, change{x: ooi, y: oi, r: '.'})
					} else {
						treecnt := Counter(chart, ooi-1, oi-1, '|')
						if treecnt < 1 {
							changeLog = append(changeLog, change{x: ooi, y: oi, r: '.'})
						}
					}
				}
			}
		}

		for _, i := range changeLog {
			chart[i.y][i.x] = i.r
		}
		changeLog = changeLog[:0]

		// Compare entire thing as a string
		var tmp strings.Builder
		for _, k := range chart {
			tmp.WriteString(string(k))
		}
		df := false
		for _, d := range duplist {
			if d == tmp.String() {
				realduplist = AddUnique(tmp.String(), realduplist)
				if duplen < len(realduplist) {
					duplen = len(realduplist)
					df = true
				}
			}
		}

		if !df && len(realduplist) > 0 {
			if (1000000000-j-1)%len(realduplist) == 0 {
				treecnt := 0
				lmbcnt := 0
				for _, c := range chart {
					for _, a := range c {
						switch a {
						case '|':
							treecnt++
						case '#':
							lmbcnt++
						}
					}
				}
				res := treecnt * lmbcnt
				input_t_elapsed = time.Since(t_input)
				log.Println("P2:", input_t_elapsed)
				log.Println("J", j)
				log.Println("Result:", res)
				os.Exit(0)
			}
		}

		duplist = AddUnique(tmp.String(), duplist)
	}
}
