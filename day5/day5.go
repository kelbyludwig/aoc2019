package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func load() []int {
	m := make([]int, 0)

	f, err := os.Open("5.txt")
	if err != nil {
		log.Fatal("error opening file")
	}
	sc := bufio.NewScanner(f)

	for sc.Scan() {
		icStrings := strings.Split(sc.Text(), ",")
		for _, icString := range icStrings {
			ic, err := strconv.Atoi(icString)
			if err != nil {
				log.Fatal(err)
			}
			m = append(m, ic)
		}
	}
	return m
}

func value(memory []int, positionMode int, arg int) int {
	if positionMode == 0 {
		return memory[arg]
	} else if positionMode == 1 {
		return arg
	}
	log.Fatalf("unknown positionMode %d", positionMode)
	return -1
}

func digits(i int) [4]int {
	return [4]int{i % 100, (i / 100) % 10, (i / 1000) % 10, (i / 10000) % 10}
}

func nextInstruction(memory []int, ip int) (opcode int, modes []int, args []int, newIp int) {

	// parse first part of instruction into opcode and mode digits
	ds := digits(memory[ip])
	opcode = ds[0]
	modes = []int{ds[1], ds[2], ds[3]}

	// setup args and bump ip
	switch opcode {
	case 1, 2:
		args = []int{memory[ip+1], memory[ip+2], memory[ip+3]}
		newIp = ip + 4
	case 3, 4:
		args = []int{memory[ip+1]}
		newIp = ip + 2
	case 99:
		args = []int{}
		newIp = ip + 1
	default:
		log.Fatalf("cannot parse opcode %d", opcode)
	}
	return
}

func execute(mem []int, input []int) (output []int) {
	output = make([]int, 0)
	ip := 0
	inpp := 0
	for ip < len(mem) {
		code, modes, args, newip := nextInstruction(mem, ip)
		ip = newip
		switch code {
		case 99:
			return output
		case 1:
			left := value(mem, modes[0], args[0])
			right := value(mem, modes[1], args[1])
			destination := args[2]
			log.Printf("mem[%d] = %d + %d", destination, left, right)
			mem[destination] = left + right
		case 2:
			left := value(mem, modes[0], args[0])
			right := value(mem, modes[1], args[1])
			destination := args[2]
			log.Printf("mem[%d] = %d * %d", destination, left, right)
			mem[destination] = left * right
		case 3:
			log.Printf("mem[%d] = %d", args[0], input[inpp])
			mem[args[0]] = input[inpp]
			inpp++
		case 4:
			log.Printf("writing output %d", mem[args[0]])
			output = append(output, mem[args[0]])
		default:
			log.Fatalf("unrecognized instruction %d", code)
		}
	}
	log.Fatalf("no more instructions to execute")
	return nil
}

func main() {
	if digits(1002) != [4]int{2, 0, 1, 0} {
		log.Fatalf("digits result is incorrect: %+v", digits(1002))
	}
	memory := load()
	input := []int{1}
	output := execute(memory, input)
	fmt.Printf("Day 5 Part 1: %v", output)
}
