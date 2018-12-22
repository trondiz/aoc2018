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

func main() {
	t_input := time.Now()
	b, err := ioutil.ReadFile("d16.data")
	check(err)
	funcs := []func([]int, []int) []int{addr, addi, mulr, muli, banr, bani, borr, bori, setr, seti, gtir, gtri, gtrr, eqir, eqri, eqrr}
	foo := strings.Split(string(b), "\n")
	tcntr := 0
	before := make([]int, 4)
	after := make([]int, 4)
	op := make([]int, 4)

	for _, f := range foo {
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

		cntr := 0
		for _, f := range funcs {
			bb := make([]int, 4)
			copy(bb, before)
			tmp := f(op, bb)
			log.Println(tmp)
			if cmp(tmp, after) {
				cntr++
			}
		}
		if cntr >= 3 {
			tcntr++
		}
		log.Println(before, op, after, cntr)

	}

	input_t_elapsed := time.Since(t_input)
	log.Println("P1:", input_t_elapsed)
	log.Println("Solution:", tcntr)

}
