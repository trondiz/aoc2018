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
				//log.Println(string(chrt[yr][xr]), y, x, yr, xr)
			}

		}
	}
	return cntr
}

func AddUnique(v int, a []int) []int {
	for _, b := range a {
		if v == b {
			return a
		}
	}
	return append(a, v)
}

func main() {
	t_input := time.Now()
	b, err := ioutil.ReadFile("d18.data")
	check(err)

	foo := strings.Split(string(b), "\n")
	foo = foo[:len(foo)-1]
	chart := make([][]rune, len(foo))
	oldchart := make([][]rune, len(foo))
	for fi, f := range foo {
		chart[fi] = []rune(f)
		oldchart[fi] = []rune(f)
	}
	input_t_elapsed := time.Since(t_input)
	log.Println("Input:", input_t_elapsed)
	t_input = time.Now()
	duplist := make([]int, 0)
	for j := 0; j < 1000000000; j++ {
		for oi, o := range oldchart {
			for ooi, oo := range o {
				switch oo {
				case '.':
					treecnt := Counter(oldchart, ooi-1, oi-1, '|')
					if treecnt >= 3 {
						chart[oi][ooi] = '|'
					}
				case '|':
					lmbcnt := Counter(oldchart, ooi-1, oi-1, '#')
					if lmbcnt >= 3 {
						chart[oi][ooi] = '#'
					}
				case '#':
					lmbcnt := Counter(oldchart, ooi-1, oi-1, '#')
					treecnt := Counter(oldchart, ooi-1, oi-1, '|')
					//log.Println(lmbcnt, treecnt, ooi, oi)
					if lmbcnt < 1 || treecnt < 1 {
						chart[oi][ooi] = '.'
					}
				}
			}
		}
		for i, ii := range chart {
			for l, ll := range ii {
				oldchart[i][l] = ll
			}
			//log.Println(string(ii))
		}
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
		for _, d := range duplist {
			if d == res {
				if (1000000000-j-1)%(len(duplist)+1) == 0 {
					log.Println("one billion!:", res)
					os.Exit(0)
				}
			}
		}
		// Lets wait a while before checking for duplicates
		if j > 1000 {
			duplist = AddUnique(res, duplist)
		}

	}

}
