package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Circuit struct {
	X         int
	Y         int
	MoveCount int
	Moves     []string
	CurrMove  int
	CurrIndex int
}


var grid [][]rune
func partTwo() {

	wireOne = "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51"
	wireTwo = "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"

	var size = 100000
	grid = make([][]rune, size)
	for k := range grid {
		grid[k] = make([]rune, size)
	}

	circuitOne := &Circuit{
		X:         size / 2,
		Y:         size / 2,
		Moves:     strings.Split(wireOne, ","),
	}

	circuitTwo := &Circuit{
		X:         size / 2,
		Y:         size / 2,
		Moves:     strings.Split(wireTwo, ","),
	}

	//log.Printf("circuittwo moves %+v", circuitTwo.Moves)
	for i:=0;i<10000;i++ {
		found := move(circuitOne)
		if found {
			log.Printf("moves %d %d", circuitOne.MoveCount, circuitTwo.MoveCount)
			break
		}
		found = move(circuitTwo)
		log.Printf("circuit2 %d %d", circuitTwo.X, circuitTwo.Y)
		if found {
			log.Printf("moves %d %d", circuitOne.MoveCount, circuitTwo.MoveCount)
			break
		}
	}
}

func move(circuit *Circuit) bool {
	direction := string(circuit.Moves[circuit.CurrMove][0])
	if grid[circuit.X][circuit.Y] == '-' && circuit.X != 50000 && circuit.Y	!= 50000 {
		//log.Printf("intersection found at %d %d", circuit.X, circuit.Y)
		return true
	}
	switch direction {
	case "R":
		grid[circuit.X][circuit.Y] = '-'
		circuit.X++
	case "L":
		grid[circuit.X][circuit.Y] = '-'
		circuit.X--
	case "D":
		grid[circuit.X][circuit.Y] = '-'
		circuit.Y++
	case "U":
		grid[circuit.X][circuit.Y] = '-'
		circuit.Y--
	}

	magnitude, err := strconv.Atoi(circuit.Moves[circuit.CurrMove][1:])
	if err != nil {
		panic(fmt.Sprintf("failed to parse move %s", circuit))
	}
	if circuit.CurrIndex == magnitude {
		circuit.CurrMove++
		circuit.CurrIndex = 0
	}

	circuit.CurrIndex++
	circuit.MoveCount++

	return false
}