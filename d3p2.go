package main

import (
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ClaimAtoi(vs []string) []int {
	vsm := make([]int, len(vs)+1)
	for i, v := range vs {
		var e error
		vsm[i], e = strconv.Atoi(v)
		check(e)
	}
	return vsm
}

func main() {
	b, err := ioutil.ReadFile("d3.data")
	check(err)
	foo := strings.Split(string(b), "\n")
	foo = foo[:len(foo)-1]
	claims := make([][]string, 0)
	for _, f := range foo {
		r := regexp.MustCompile(`#(?P<Id>\d+) @ (?P<x>\d+),(?P<y>\d+): (?P<sx>\d+)x(?P<sy>\d+)`)
		res := r.FindStringSubmatch(f)
		claims = append(claims, res)
	}

	claimsreal := make([][]int, len(claims))
	for i, claim := range claims {
		claim = append(claim[:0], claim[1:]...)
		claimsreal[i] = ClaimAtoi(claim)
	}

	for _, cur := range claimsreal {
		for x := cur[1]; x <= cur[1]+cur[3]; x++ {
			for y := cur[2]; y <= cur[2]+cur[4]; y++ {
				for _, claim := range claimsreal {
					if cur[0] == claim[0] {
						// we always overlap ourselves
						continue
					}
					if x > claim[1] && x <= claim[1]+claim[3] && y > claim[2] && y <= claim[2]+claim[4] {
						//log.Println("X:", x, "Y:", y, "Claim overlaps:", claim, "Pic:", claim[1], claim[2], claim[1]+claim[3], claim[2]+claim[4])
						// store that this claim overlaps
						cur[5] = 1
					}
				}
			}
		}
	}

	for _, cur := range claimsreal {
		if cur[5] != 1 {
			log.Println(cur[0])
		}
	}
}
