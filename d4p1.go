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

func MyAtoi(v string) int {
	c, _ := strconv.Atoi(v)
	return c
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
	curGuardId := 0
	curGuardSleeping := false
	//1518-02-10
	for m := 2; m <= 12; m++ {
		for d := 1; d <= 31; d++ {
			for h := 0; h <= 23; h++ {
				for mi := 0; mi <= 59; mi++ {
					// Find log record
					for _, r := range records {
						if MyAtoi(r[2]) == m && MyAtoi(r[3]) == d && MyAtoi(r[4]) == h && MyAtoi(r[5]) == mi {
							switch {
							case shiftStart.MatchString(r[6]):
								id := regexp.MustCompile(`Guard #(?P<Id>\d+) begins shift`)
								res := id.FindStringSubmatch(r[6])
								curGuardId = MyAtoi(res[1])
								curGuardSleeping = false
							case guardsleep.MatchString(r[6]):
								curGuardSleeping = true
							case guardwakes.MatchString(r[6]):
								curGuardSleeping = false
							}
							//log.Println(r, "time is now", m, d, h, mi)
							break
						}
					}
					if curGuardSleeping && h == 0 {
						//log.Println(curGuardId, "is sleeping at", m, d, h, mi, curGuardSleeping)
						guards[curGuardId][mi]++
					}
				}
			}
		}
	}

	commonSleeperId := 0
	commonSleeperMinute := 0
	commonSleeperMins := 0
	commonSleeperSum := 0
	for gi, g := range guards {
		sum := 0
		for _, m := range g {
			sum += m
		}
		if sum > commonSleeperSum {
			commonSleeperId = gi
			commonSleeperSum = sum
		}
	}
	for bi, b := range guards[commonSleeperId] {
		if b > commonSleeperMins {
			commonSleeperMins = b
			commonSleeperMinute = bi
		}
	}

	log.Println(commonSleeperId, commonSleeperMinute, commonSleeperMins)
	log.Println(commonSleeperId * commonSleeperMinute)
}
