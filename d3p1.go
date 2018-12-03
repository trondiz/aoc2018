package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	b, err := ioutil.ReadFile("d3.data")
	check(err)
	foo := strings.Split(string(b), "\n")

	for _, f := range foo {
		data := strings.Split(string(f), " ")
		claims := make(map[string]map[string]string)

		//		claims[data[0]] = make(map[string]int)
		log.Println(data)
		xy := strings.Split(string(data[2]), ",")
		y := strings.Split(string(xy[1]), ":")
		rx, err := strconv.Atoi(xy[0])
		ry, err := strconv.Atoi(y[0])
		claims[data[0]]["x"] = data[rx]
		claims[data[0]]["y"] = data[ry]
		sxy := strings.Split(string(data[3]), "x")
		rsx, err := strconv.Atoi(sxy[0])
		rsy, err := strconv.Atoi(sxy[1])
		claims[data[0]]["sx"] = data[rsx]
		claims[data[0]]["sy"] = data[rsy]
		log.Println(claims)
		check(err)
	}

}
