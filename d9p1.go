package main

import "log"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func insertAtPos(m int, pos int, circle []int) []int {
	circle = append(circle, 0)
	copy(circle[pos+1:], circle[pos:])
	circle[pos] = m
	return circle
}
func removeAtPos(pos int, circle []int) []int {
	circle = append(circle[:pos], circle[pos+1:]...)
	return circle
}

func main() {
	//var err error

	// init game state
	currentMarble := 0
	currentMarblePos := 0
	circle := make([]int, 0)
	currentPlayer := 1
	playerMax := 403
	marbleMax := 71920

	// create game pieces
	marbles := make([]int, 0)
	for i := 0; i <= marbleMax; i++ {
		marbles = append(marbles, i)
	}
	players := make([]int, playerMax)
	//log.Println(players)
	// game loop
	for len(marbles) > 0 {
		//Get next marbles
		currentMarble, marbles = marbles[0], marbles[1:]
		// 0 is special
		if currentMarble == 0 {
			circle = append(circle, currentMarble)
			continue
		}

		nextPos := currentMarblePos + 2
		if nextPos > len(circle) {
			nextPos -= len(circle)
		}

		if currentMarble%23 == 0 {
			//log.Println(currentPlayer, currentMarblePos, len(circle))
			players[currentPlayer] += currentMarble
			otherone := currentMarblePos - 7
			if otherone < 0 {
				otherone += len(circle)
			}
			nextPos = otherone
			players[currentPlayer] += circle[otherone]
			circle = removeAtPos(otherone, circle)
		} else {
			circle = insertAtPos(currentMarble, nextPos, circle)
		}

		currentMarblePos = nextPos
		currentPlayer++
		if currentPlayer >= len(players) {
			currentPlayer -= len(players)
		}
	}
	highscore := 0
	for _, p := range players {
		if p > highscore {
			highscore = p
		}
	}
	log.Println(highscore)
}
