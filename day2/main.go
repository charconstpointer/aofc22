package main

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("day2.txt")
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	line, _, err := r.ReadLine()
	if err != nil {
		log.Fatalf("failed to read line: %v", err)
	}
	rounds := make([]Round, 0)
	for line != nil {
		rounds = append(rounds, parse(string(line)))
		line, _, err = r.ReadLine()
		if err != nil {
			if !errors.Is(err, io.EOF) {
				log.Fatalf("failed to read line: %v", err)
			}
		}
	}
	score := 0
	for _, r := range rounds {
		score += r.Counter2()
	}
	log.Printf("score: %d", score)
}

func parse(l string) Round {
	splitted := strings.Split(l, " ")
	var (
		op      Op
		counter Counter
	)
	switch splitted[0] {
	case "A":
		op = A
	case "B":
		op = B
	case "C":
		op = C
	}
	switch splitted[1] {
	case "X":
		counter = X
	case "Y":
		counter = Y
	case "Z":
		counter = Z
	}

	return Round{
		Op:      Op(op),
		Counter: Counter(counter),
	}
}

type Op int

const (
	A Op = iota // rock
	B           // paper
	C           // scissors
)

type Counter int

const (
	X Counter = iota // rock
	Y                //paper
	Z                //scissors
)

type Res int

const (
	Win Res = iota
	Draw
	Loss
)

var values map[Counter]int = map[Counter]int{
	Y: 2,
	X: 1,
	Z: 3,
}

var outcomeValues map[Res]int = map[Res]int{
	Win:  6,
	Draw: 3,
	Loss: 0,
}

type Round struct {
	Op      Op
	Counter Counter
}

func (r *Round) Result() Res {
	score := int(r.Op) - int(r.Counter)
	if score == -1 || score == 2 {
		return Win
	}
	if score == 0 {
		return Draw
	}
	return Loss
}

func (r *Round) Counter2() int {
	var res Res
	switch r.Counter {
	case X:
		res = Loss
	case Y:
		res = Draw
	case Z:
		res = Win
	}

	if res == Win {
		next := int(r.Op) + 1
		counter := Counter(next % 3)
		return values[counter] + outcomeValues[res]
	}
	if res == Draw {
		counter := Counter(r.Op)
		return values[counter] + outcomeValues[res]
	}
	next := int(r.Op) - 1
	counter := Counter((next + 3) % 3)
	return values[counter] + outcomeValues[res]
}

func calcScore(r Round) int {
	result := r.Result()
	return outcomeValues[result] + values[Counter(r.Counter)]
}
