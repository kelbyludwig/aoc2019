package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	f, err := os.Open("1.txt")
	if err != nil {
		log.Fatal("unable to read file")
		return
	}
	sc := bufio.NewScanner(f)

	sum := 0
	modules := make([]int, 0)
	for sc.Scan() {
		var mass int
		n, err := fmt.Sscanf(sc.Text(), "%d", &mass)
		if err != nil {
			log.Fatal("error scanning file")
			return
		}
		if n != 1 {
			log.Fatal("incorrect number of tokens")
			return
		}
		modules = append(modules, mass)
		fuel := mass/3.0 - 2
		fmt.Printf("Read %v : %v ", sc.Text(), sum)
		sum += fuel
		fmt.Printf("-= %v = %v\n", fuel, sum)
	}

	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1 Answer: %v\n", sum)

	sum, fuel := 0, 0
	for _, module := range modules {
		fuel = module/3.0 - 2
		for fuel > 0 {
			sum += fuel
			fuel = fuel/3.0 - 2
		}
	}

	fmt.Printf("Part 2 Answer: %v\n", sum)
}
