package main

import (
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func addr(op, regs []int) []int {
	regs[op[3]] = regs[op[1]] + regs[op[2]]
	return regs
}

func addi(op, regs []int) []int {
	regs[op[3]] = regs[op[1]] + op[2]
	return regs
}

func mulr(op, regs []int) []int {
	regs[op[3]] = regs[op[1]] * regs[op[2]]
	return regs
}

func muli(op, regs []int) []int {
	regs[op[3]] = regs[op[1]] * op[2]
	return regs
}

func banr(op, regs []int) []int {
	regs[op[3]] = regs[op[1]] & regs[op[2]]
	return regs
}

func bani(op, regs []int) []int {
	regs[op[3]] = regs[op[1]] & op[2]
	return regs
}

func borr(op, regs []int) []int {
	regs[op[3]] = regs[op[1]] | regs[op[2]]
	return regs
}

func bori(op, regs []int) []int {
	regs[op[3]] = regs[op[1]] | op[2]
	return regs
}

func setr(op, regs []int) []int {
	regs[op[3]] = regs[op[1]]
	return regs
}

func seti(op, regs []int) []int {
	regs[op[3]] = op[1]
	return regs
}

func gtir(op, regs []int) []int {
	if op[1] > regs[op[2]] {
		regs[op[3]] = 1
	} else {
		regs[op[3]] = 0
	}
	return regs
}

func gtri(op, regs []int) []int {
	if regs[op[1]] > op[2] {
		regs[op[3]] = 1
	} else {
		regs[op[3]] = 0
	}
	return regs
}

func gtrr(op, regs []int) []int {
	if regs[op[1]] > regs[op[2]] {
		regs[op[3]] = 1
	} else {
		regs[op[3]] = 0
	}
	return regs
}

func eqir(op, regs []int) []int {
	if op[1] == regs[op[2]] {
		regs[op[3]] = 1
	} else {
		regs[op[3]] = 0
	}
	return regs
}

func eqri(op, regs []int) []int {
	if regs[op[1]] == op[2] {
		regs[op[3]] = 1
	} else {
		regs[op[3]] = 0
	}
	return regs
}

func eqrr(op, regs []int) []int {
	if regs[op[1]] == regs[op[2]] {
		regs[op[3]] = 1
	} else {
		regs[op[3]] = 0
	}
	return regs
}

func MyAtoi(v []string) []int {
	r := make([]int, 4)
	for vi, vv := range v {
		r[vi], _ = strconv.Atoi(vv)
	}
	return r
}

func cmp(a, b []int) bool {
	for ai, av := range a {
		if av != b[ai] {
			return false
		}
	}
	return true
}

func Intersection(a, b []int) (c []int) {
	m := make(map[int]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; ok {
			c = append(c, item)
		}
	}
	return
}

func main() {
	t_input := time.Now()
	b, err := ioutil.ReadFile("d16.data")
	check(err)
	funcs := []func([]int, []int) []int{addr, addi, mulr, muli, banr, bani, borr, bori, setr, seti, gtir, gtri, gtrr, eqir, eqri, eqrr}
	opmap := make([][][]int, 16)
	foo := strings.Split(string(b), "\n")
	before := make([]int, 4)
	after := make([]int, 4)
	op := make([]int, 4)
	testprogindex := 0
	for fi, f := range foo {
		// Before: [0, 1, 2, 1]
		// 14 1 3 3
		// After:  [0, 1, 2, 1]
		b := regexp.MustCompile(`Before: \[(\d+), (\d+), (\d+), (\d+)\]`)
		btmp := b.FindStringSubmatch(f)
		if len(btmp) > 0 {
			btmp = append(btmp[:0], btmp[1:]...)
			before = MyAtoi(btmp)
			continue
		}
		o := regexp.MustCompile(`(\d+) (\d+) (\d+) (\d+)`)
		otmp := o.FindStringSubmatch(f)
		if len(otmp) > 0 {
			otmp = append(otmp[:0], otmp[1:]...)
			op = MyAtoi(otmp)
			continue
		}
		a := regexp.MustCompile(`After:  \[(\d+), (\d+), (\d+), (\d+)\]`)
		atmp := a.FindStringSubmatch(f)
		if len(atmp) > 0 {
			atmp = append(atmp[:0], atmp[1:]...)
			after = MyAtoi(atmp)
			continue
		}

		mf := make([]int, 0)
		for fi, f := range funcs {
			bb := make([]int, 4)
			copy(bb, before)
			tmp := f(op, bb)

			if cmp(tmp, after) {
				//mf = AddUnique
				mf = append(mf, fi)
			}
		}
		opmap[op[0]] = append(opmap[op[0]], mf)
		//log.Println(before, op, after)
		if foo[fi+1] == "" {
			testprogindex = fi + 3
			break
		}
	}
	//log.Println(foo[testprogindex])

	// get intersection of function arrays
	opmapr := make([][]int, 16)
	for oi, _ := range opmapr {
		t := opmap[oi][0]
		for _, mr := range opmap[oi] {
			t = Intersection(t, mr)
		}
		//log.Println(oi, t)
		opmapr[oi] = t
	}
	//opmaps := make([]int, 16)

	for i := 0; i < 20; i++ {
		for oi, o := range opmapr {
			if len(o) == 1 {
				for ooi, ooo := range opmapr {
					for ni, n := range ooo {
						if n == o[0] && ooi != oi {
							//log.Println("Fond", n, "and", o[0], "in", ni)
							opmapr[ooi] = append(opmapr[ooi][:ni], opmapr[ooi][ni+1:]...)
						}
					}
				}
			}
		}
	}
	//log.Println(opmapr)

	registers := make([]int, 4)
	for i := testprogindex; i < len(foo); i++ {
		o := regexp.MustCompile(`(\d+) (\d+) (\d+) (\d+)`)
		otmp := o.FindStringSubmatch(foo[i])
		if len(otmp) > 0 {
			otmp = append(otmp[:0], otmp[1:]...)
			op = MyAtoi(otmp)
			funcs[opmapr[op[0]][0]](op, registers)
		}
	}
	//log.Println(registers)

	input_t_elapsed := time.Since(t_input)
	log.Println("P2:", input_t_elapsed)
	log.Println("Solution:", registers[0])

}
