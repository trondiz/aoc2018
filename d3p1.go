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
	vsm := make([]int, len(vs))
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
		//r := regexp.MustCompile(`(?P<Id>#\d+) @ (?P<x>\d+),(?P<y>\d+): (?P<sx>\d+)x(?P<sy>\d+)`)
		r := regexp.MustCompile(`(?P<Id>#\d+) @ (?P<x>\d+),(?P<y>\d+): (?P<sx>\d+)x(?P<sy>\d+)`)
		res := r.FindStringSubmatch(f)
		claims = append(claims, res)
	}
	claimsreal := make([][]int, len(claims))
	for i, claim := range claims {
		claim = append(claim[:0], claim[1:]...)
		claim = append(claim[:0], claim[1:]...)
		claimsreal[i] = ClaimAtoi(claim)
	}
	collcounter := 0

	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			claimcounter := 0
			for _, claim := range claimsreal {
				if x > claim[0] && x <= claim[0]+claim[2] && y > claim[1] && y <= claim[1]+claim[3] {
					//log.Println("X:", x, "Y:", y, "Claim matches:", claim, "Pic:", claim[0], claim[1], claim[0]+claim[2], claim[1]+claim[3])
					claimcounter++
				}
				if claimcounter == 2 {
					//log.Println("Breaking off")
					collcounter++
					//log.Println(collcounter)
					break
				}
			}
		}
	}
	log.Println(collcounter)
}
