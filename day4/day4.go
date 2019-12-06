package main

import (
    "log"
)

const LB int = 272091
const UB int = 815432

func digits(i int) [6]int {
    return  [6]int{(i/100000) % 10, (i/10000) % 10, (i/1000) % 10, (i/100) % 10, (i/10) % 10, i % 10}
}

func adjacent(ds [6]int) bool {
    var r bool
    for i := 0; i < 5; i++ {
        r = r || ds[i] == ds[i+1]
    }
    return r
}

func decreasing(ds [6]int) bool {
    r := true
    p := -1
    for _, d := range ds {
       r = r && d >= p
       p = d
    }
    return r
}

func valid(i int) bool {
    ds := digits(i)
    return adjacent(ds) && decreasing(ds)
}

func main() {

    lb := digits(LB)
    ub := digits(UB)

    if lb != [6]int{2,7,2,0,9,1} {
        log.Fatalf("lb digits mismatch: %v", lb)
    }

    if ub != [6]int{8,1,5,4,3,2} {
        log.Fatalf("ub digits mismatch: %v", lb)
    }

    if !adjacent(digits(111111)) || adjacent(digits(123456)) {
        log.Fatalf("adjacent is broken")
    }

    if !decreasing(digits(111123)) || !decreasing(digits(135679)) || decreasing(digits(111011)) {
        log.Fatalf("decreasing is broken")
    }

    if !valid(111111) || valid(223450) || valid(123789) {
        log.Fatalf("valid is broken")
    }

    count := 0
    for i := LB; i <= UB; i++ {
        if valid(i) {
            count += 1
        }
    }
    log.Printf("Day 4 Part 1: %v", count)

}
