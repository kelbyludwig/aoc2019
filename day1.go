package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

func main() {

    f, err := os.Open("inputs/1.txt")
    if err != nil {
        log.Fatal("unable to read file")
        return
    }
    sc := bufio.NewScanner(f)

    sum := 0
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
        fuel := mass/3.0 - 2
        fmt.Printf("Read %v : %v ", sc.Text(), sum)
        sum += fuel
        fmt.Printf("-= %v = %v\n", fuel, sum)
    }

    if err := sc.Err(); err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Answer: %v\n" , sum)
}
