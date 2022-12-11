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
	ok := true
	prio := 0
	for ok {
		var rows []string
		for i := 0; i < 3; i++ {
			ok = sc.Scan()
			rows = append(rows, sc.Text())
		}
		common := intersectMulti([]byte(rows[0]), []byte(rows[1]), []byte(rows[2]))
		log.Printf("Common: %s", common)
		for _, c := range common {
			prio += asciiToPrio(c)
		}
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
	var common []byte
	if len(r.c) == 2 {
		common = intersect([]byte(*r.c[0]), []byte(*r.c[1]))
	} else if len(r.c) == 3 {
		common = intersectMulti([]byte(*r.c[0]), []byte(*r.c[1]), []byte(*r.c[2]))
	}
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

func intersectMulti(a, b, c []byte) []byte {
	return intersect(intersect(a, b), c)
}
