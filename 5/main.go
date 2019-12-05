package main

import (
	"log"
	"strconv"
)

func main() {
	log.Printf("Part one answer: %d", partOne(5))
}

func parseInstruction(opcode int) []string {
	opS := strconv.Itoa(opcode)
	in := make([]string, 5)
	for i := len(opS) - 1; i >= 0; i-- {
		in[len(opS)-i-1] = string(opS[i])
	}
	return in
}

func partOne(input int) int {
	var pc int // program counter

programLoop:
	for {
		opcode := program[pc]

		in := parseInstruction(opcode)
		var code string
		if len(in) == 1 {
			code = in[0]
		} else {
			c, _ := strconv.Atoi(in[1] + in[0])
			code = strconv.Itoa(c)
		}

		switch code {
		case "99":
			break programLoop
		case "1":
			program[program[pc+3]] = getVal(in[2], pc+1) + getVal(in[3], pc+2)
			log.Printf("pc %d read instruction %v with params %d %d, added result in %d", pc, in, program[pc+1], program[pc+2], program[pc+3])
			pc += 4
		case "2":
			program[program[pc+3]] = getVal(in[2], pc+1) * getVal(in[3], pc+2)
			log.Printf("pc %d read instruction %v with params %d %d, multiplied result in %d", pc, in, program[pc+1], program[pc+2], program[pc+3])
			pc += 4
		case "3":
			program[program[pc+1]] = input
			log.Printf("pc %d read instruction %v with params %d and input %d", pc, in, program[pc+1], input)
			pc += 2
		case "4":
			output := getVal(in[2], pc+1)
			log.Printf("output %d", output)
			pc += 2
		case "5": //jmpT
			if getVal(in[2], pc+1) != 0 {
				log.Printf("jmp to %d", getVal(in[3], pc+2))
				pc = getVal(in[3], pc+2)
			} else {
				pc += 3
			}
		case "6": //jmpT
			if getVal(in[2], pc+1) == 0 {
				pc = getVal(in[3], pc+2)
				log.Printf("jmp to %d", getVal(in[3], pc+2))
			} else {
				pc += 3
			}
		case "7": // storeLT
			if getVal(in[2], pc+1) < getVal(in[3], pc+2) {
				program[program[pc+3]] = 1
			} else {
				program[program[pc+3]] = 0
			}
			pc += 4
		case "8": // storeEq
			if getVal(in[2], pc+1) == getVal(in[3], pc+2) {
				program[program[pc+3]] = 1
			} else {
				program[program[pc+3]] = 0
			}
			pc += 4
		default:
			log.Printf("unknown opcode %d", code)
		}
	}

	return program[0]
}

func getVal(mode string, pc int) int {
	switch mode {
	case "1": // immediate
		return program[pc]
	default: // position
		return program[program[pc]]
	}
}

var program = []int{3, 225, 1, 225, 6, 6, 1100, 1, 238, 225, 104, 0, 1102, 91, 92, 225, 1102, 85, 13, 225, 1, 47, 17, 224, 101, -176, 224, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 7, 224, 1, 223, 224, 223, 1102, 79, 43, 225, 1102, 91, 79, 225, 1101, 94, 61, 225, 1002, 99, 42, 224, 1001, 224, -1890, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 6, 224, 1, 224, 223, 223, 102, 77, 52, 224, 1001, 224, -4697, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 7, 224, 1, 224, 223, 223, 1101, 45, 47, 225, 1001, 43, 93, 224, 1001, 224, -172, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 1, 224, 1, 224, 223, 223, 1102, 53, 88, 225, 1101, 64, 75, 225, 2, 14, 129, 224, 101, -5888, 224, 224, 4, 224, 102, 8, 223, 223, 101, 6, 224, 224, 1, 223, 224, 223, 101, 60, 126, 224, 101, -148, 224, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 2, 224, 1, 224, 223, 223, 1102, 82, 56, 224, 1001, 224, -4592, 224, 4, 224, 1002, 223, 8, 223, 101, 4, 224, 224, 1, 224, 223, 223, 1101, 22, 82, 224, 1001, 224, -104, 224, 4, 224, 1002, 223, 8, 223, 101, 4, 224, 224, 1, 223, 224, 223, 4, 223, 99, 0, 0, 0, 677, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1105, 0, 99999, 1105, 227, 247, 1105, 1, 99999, 1005, 227, 99999, 1005, 0, 256, 1105, 1, 99999, 1106, 227, 99999, 1106, 0, 265, 1105, 1, 99999, 1006, 0, 99999, 1006, 227, 274, 1105, 1, 99999, 1105, 1, 280, 1105, 1, 99999, 1, 225, 225, 225, 1101, 294, 0, 0, 105, 1, 0, 1105, 1, 99999, 1106, 0, 300, 1105, 1, 99999, 1, 225, 225, 225, 1101, 314, 0, 0, 106, 0, 0, 1105, 1, 99999, 8, 226, 677, 224, 102, 2, 223, 223, 1005, 224, 329, 1001, 223, 1, 223, 1007, 226, 226, 224, 1002, 223, 2, 223, 1006, 224, 344, 101, 1, 223, 223, 108, 226, 226, 224, 1002, 223, 2, 223, 1006, 224, 359, 1001, 223, 1, 223, 107, 226, 677, 224, 102, 2, 223, 223, 1006, 224, 374, 101, 1, 223, 223, 8, 677, 677, 224, 102, 2, 223, 223, 1006, 224, 389, 1001, 223, 1, 223, 1008, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 404, 101, 1, 223, 223, 7, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 419, 101, 1, 223, 223, 1108, 226, 677, 224, 1002, 223, 2, 223, 1005, 224, 434, 101, 1, 223, 223, 1108, 226, 226, 224, 102, 2, 223, 223, 1005, 224, 449, 1001, 223, 1, 223, 107, 226, 226, 224, 102, 2, 223, 223, 1005, 224, 464, 101, 1, 223, 223, 1007, 677, 677, 224, 102, 2, 223, 223, 1006, 224, 479, 101, 1, 223, 223, 1007, 226, 677, 224, 102, 2, 223, 223, 1005, 224, 494, 1001, 223, 1, 223, 1008, 226, 226, 224, 1002, 223, 2, 223, 1005, 224, 509, 1001, 223, 1, 223, 1108, 677, 226, 224, 1002, 223, 2, 223, 1006, 224, 524, 1001, 223, 1, 223, 108, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 539, 101, 1, 223, 223, 108, 226, 677, 224, 1002, 223, 2, 223, 1005, 224, 554, 101, 1, 223, 223, 1008, 677, 677, 224, 1002, 223, 2, 223, 1006, 224, 569, 1001, 223, 1, 223, 1107, 677, 677, 224, 102, 2, 223, 223, 1005, 224, 584, 1001, 223, 1, 223, 7, 677, 226, 224, 102, 2, 223, 223, 1005, 224, 599, 1001, 223, 1, 223, 8, 677, 226, 224, 1002, 223, 2, 223, 1005, 224, 614, 1001, 223, 1, 223, 7, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 629, 101, 1, 223, 223, 1107, 677, 226, 224, 1002, 223, 2, 223, 1005, 224, 644, 1001, 223, 1, 223, 1107, 226, 677, 224, 102, 2, 223, 223, 1006, 224, 659, 1001, 223, 1, 223, 107, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 674, 101, 1, 223, 223, 4, 223, 99, 226}
