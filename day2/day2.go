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

    f, err := os.Open("2.txt")
    if err != nil  {
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
            m = append(m , ic)
        }
    }
    return m
}

func execute(mem []int) {
    for ip := 0; ip < len(mem); ip += 4 {
        opcode, srcl, srcr, dst := mem[ip], mem[ip+1], mem[ip+2], mem[ip+3]
        switch opcode {
            case 99:
                return
            case 1:
                mem[dst] = mem[srcl] + mem[srcr]
            case 2:
                mem[dst] = mem[srcl] * mem[srcr]
            default:
                log.Fatalf("unrecognized opcode %d", opcode)
        }
    }
}

func main() {
    // load the program into memory
    mem := load()
    // set the 1202 program alarm state
    mem[1] = 12
    mem[2] = 2
    // interpret updated memory
    execute(mem)
    fmt.Printf("Day 2 Part 1: %d", mem[0])
}
