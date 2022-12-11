package main

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"unicode"
)

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	sc := bufio.NewScanner(bytes.NewReader(b))
	var rucksacks []*Rucksack
	for sc.Scan() {
		rucksacks = append(rucksacks, NewRucksack(sc.Text()))
	}
	var prio int
	for _, r := range rucksacks {
		prio += r.Prio()
	}
	log.Printf("Prio: %d", prio)
}

type Comp string

type Rucksack struct {
	c []*Comp
}

func NewRucksack(items string) *Rucksack {
	l, r := items[0:len(items)/2], items[len(items)/2:]
	lc, rc := Comp(l), Comp(r)
	return &Rucksack{
		c: []*Comp{&lc, &rc},
	}
}

func (r *Rucksack) Prio() int {
	common := intersect([]byte(*r.c[0]), []byte(*r.c[1]))
	log.Printf("Common: %s", common)
	var prio int
	for _, c := range common {
		prio += asciiToPrio(c)
	}
	return prio
}

func asciiToPrio(c byte) int {
	upper := unicode.IsUpper(rune(c))
	base := int(c - 'A')
	if upper {
		return base + 27
	}
	return base - 31
}

func intersect(a, b []byte) []byte {
	m := make(map[byte]bool)
	for _, r := range a {
		m[r] = true
	}
	var c []byte
	for _, r := range b {
		if m[r] {
			c = append(c, r)
			delete(m, r)
		}
	}
	return c
}
