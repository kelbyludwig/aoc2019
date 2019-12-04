package main

import (
    "fmt"
    "log"
    "math"
    "os"
    "bufio"
    "strings"
    "strconv"
)

type Coord struct {
    X int
    Y int
}

type Line struct {
    Traces []Trace
}

type Trace struct {
    Direction string
    Amount int
}

func loadLines() []Line {
    f, err := os.Open("3.txt")
    if err != nil {
        log.Fatal(err)
    }
    sc := bufio.NewScanner(f)
    lines := make([]Line, 0)
    for sc.Scan() {
        ts := strings.Split(sc.Text(), ",")
        traces := make([]Trace, 0)
        for _, t := range ts {
            direction := t[0:1]
            amount, err := strconv.Atoi(t[1:])
            if err != nil {
                log.Fatal(err)
            }
            traces = append(traces, Trace{direction, amount})
        }
        lines = append(lines, Line{traces})
    }
    return lines
}

func countCoords(l Line) map[Coord]int {
    currX, currY :=  0, 0
    coordCounter := make(map[Coord]int)
    for _, trace := range l.Traces {
        counter := trace.Amount
        for counter > 0 {
            switch trace.Direction {
            case "U":
                currY += 1
            case "D":
                currY -= 1
            case "L":
                currX -= 1
            case "R":
                currX += 1
            default:
                log.Fatalf("unrecoginized direction %s", trace.Direction)
            }
            counter -= 1
            // we do not care about self crosses so do not increment
            coordCounter[Coord{currX, currY}] = 1
        }
    }
    return coordCounter
}

func intersection(left, right map[Coord]int) []Coord {

    intersectors := make([]Coord, 0)
    for coord, _ := range left {
        if right[coord] > 0 {
            intersectors = append(intersectors, coord)
        }
    }
    return intersectors
}

func distance(c Coord) int {
    return int(math.Abs(float64(c.X)) + math.Abs(float64(c.Y)))
}

func main() {
    lines := loadLines()
    c0 := countCoords(lines[0])
    c1 := countCoords(lines[1])
    is := intersection(c0, c1)
    mindist := math.MaxUint32
    for _, i := range is {
        d := distance(i)
        if d < mindist {
            mindist = d
        }
    }
    fmt.Printf("Day 3 Part 1: %d\n", mindist)
}
