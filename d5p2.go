package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func removeUnit(g []string, u string) []string {
	r := make([]string, 0)
	for _, h := range g {
		if h != u && h != strings.ToLower(u) {
			r = append(r, h)
		}
	}
	return r
}

func colapsePolymer(polymers []string) int {
	plen := len(polymers)
	counter := 0
	for {
		for i := 1; i < len(polymers); i++ {
			j := i - 1
			a := strings.ToUpper(polymers[i])
			b := strings.ToUpper(polymers[j])
			if a == b && polymers[i] != polymers[j] {
				//log.Println(polymers[i], polymers[j], i, j)

				//log.Println(len(polymers))
				polymers = append(polymers[:i], polymers[i+1:]...)
				polymers = append(polymers[:j], polymers[j+1:]...)
				//log.Println(len(polymers))
			}
		}

		//.Println(len(polymers))
		//log.Println(counter)
		counter++
		if len(polymers) < plen {
			plen = len(polymers)
		} else {
			break
		}
	}

	return (len(polymers))
}

func main() {
	b, err := ioutil.ReadFile("d5.data")
	check(err)
	foo := strings.Split(string(b), "\n")
	foo = foo[:len(foo)-1]
	polymers := strings.Split(foo[0], "")
	al := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	shortest := 12000
	for _, n := range al {
		a := removeUnit(polymers, n)
		leng := colapsePolymer(a)
		if leng < shortest {
			shortest = leng
		}
	}
	log.Println(shortest)

}
