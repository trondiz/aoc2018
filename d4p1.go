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
	b, err := ioutil.ReadFile("d4.data")
	check(err)
	foo := strings.Split(string(b), "\n")
	foo = foo[:len(foo)-1]
	records := make([][]string, 0)
	for _, f := range foo {
		//log.Println(f)
		//r := regexp.MustCompile(`(?P<Id>#\d+) @ (?P<x>\d+),(?P<y>\d+): (?P<sx>\d+)x(?P<sy>\d+)`)
		r := regexp.MustCompile(`\[(?P<Year>\d+)-(?P<Month>\d+)-(?P<Day>\d+) (?P<Hour>\d+):(?P<Minute>\d+)\] (?P<action>.*)`)
		res := r.FindStringSubmatch(f)
		//log.Println(res)
		records = append(records, res)
	}
	guards := make([][]int, 5000)
	for i := 0; i < len(guards); i++ {
		guards[i] = make([]int, 60)
	}
	var shiftStart = regexp.MustCompile(`Guard`)
	var guardsleep = regexp.MustCompile(`falls asleep`)
	var guardwakes = regexp.MustCompile(`wakes up`)
	curGuard := 0
	for _, r := range records {
		switch {
		case shiftStart.MatchString(r[6]):
			id := regexp.MustCompile(`Guard #(?P<Id>\d+) begins shift`)
			res := id.FindStringSubmatch(r[6])
			curGuard, err = strconv.Atoi(res[1])
		case guardsleep.MatchString(r[6]):
			min, err := strconv.Atoi(r[5])
			check(err)
			guards[curGuard][min]++
		case guardwakes.MatchString(r[6]):
			//noop
		}
	}

	log.Println(guards[419])

	//collcounter := 0
	//for x := 0; x < 1000; x++ {
	//	for y := 0; y < 1000; y++ {
	//		claimcounter := 0
	//		for _, claim := range claimsreal {
	//			if x > claim[0] && x <= claim[0]+claim[2] && y > claim[1] && y <= claim[1]+claim[3] {
	//				//log.Println("X:", x, "Y:", y, "Claim matches:", claim, "Pic:", claim[0], claim[1], claim[0]+claim[2], claim[1]+claim[3])
	//				claimcounter++
	//			}
	//			if claimcounter == 2 {
	//				//log.Println("Breaking off")
	//				collcounter++
	//				//log.Println(collcounter)
	//				break
	//			}
	//		}
	//	}
	//}
	//log.Println(collcounter)
}
