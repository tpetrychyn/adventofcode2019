package main

import (
	"log"
)

func main() {
	//log.Printf("Part one answer: %d", partOne())

	noun, verb := partTwo()
	log.Printf("Part two answer: %d %d", noun, verb)
}

func partOne() int {
	var pc int // program counter

programLoop:
	for {
		opcode := input[pc]

		switch opcode {
		case 99:
			break programLoop
		case 1:
			input[input[pc+3]] = input[input[pc+1]] + input[input[pc+2]]
		case 2:
			input[input[pc+3]] = input[input[pc+1]] * input[input[pc+2]]
		}
		log.Printf("read opcode %d with params %d %d, stored result in %d", opcode, input[pc+1], input[pc+2], input[pc+3])
		pc += 4

	}

	log.Printf("result: %v", input)
	return input[0]
}

func partTwo() (int, int) {
	var target = 19690720

	inputCopy := append(make([]int, 0), input[:]...)

	// horrible brute force answer
	for n:=0;n<99;n++ {
		for v:=0;v<99;v++ {
			var pc int // program counter
			input = append(make([]int, 0), inputCopy[:]...)
			input[1] = n
			input[2] = v

		programLoop:
			for {
				opcode := input[pc]

				switch opcode {
				case 99:
					break programLoop
				case 1:
					input[input[pc+3]] = input[input[pc+1]] + input[input[pc+2]]
				case 2:
					input[input[pc+3]] = input[input[pc+1]] * input[input[pc+2]]
				}
				pc += 4

			}

			if input[0] == target {
				return n, v
			}
		}
	}
	return 0,0
}


var input = []int{1, 12, 2, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 1, 9, 19, 1, 19, 5, 23, 2, 6, 23, 27, 1, 6, 27, 31, 2, 31, 9, 35, 1, 35, 6, 39, 1, 10, 39, 43, 2, 9, 43, 47, 1, 5, 47, 51, 2, 51, 6, 55, 1, 5, 55, 59, 2, 13, 59, 63, 1, 63, 5, 67, 2, 67, 13, 71, 1, 71, 9, 75, 1, 75, 6, 79, 2, 79, 6, 83, 1, 83, 5, 87, 2, 87, 9, 91, 2, 9, 91, 95, 1, 5, 95, 99, 2, 99, 13, 103, 1, 103, 5, 107, 1, 2, 107, 111, 1, 111, 5, 0, 99, 2, 14, 0, 0}
