package main

import (
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Fruit struct {
	Callories int
}

type Deer struct {
	Fruits []*Fruit
}

func (d *Deer) Callories() int {
	if d == nil || d.Fruits == nil {
		return 0
	}
	c := 0
	for _, f := range d.Fruits {
		c += f.Callories
	}
	return c
}

func main() {
	f, err := os.ReadFile("day1.txt")
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}
	content := string(f)
	fruitsTokens := strings.Split(content, "\n\n")
	deers := make([]*Deer, 0, len(fruitsTokens))
	for _, fruitRow := range fruitsTokens {
		lines := strings.Split(fruitRow, "\n")
		fruits := make([]*Fruit, 0)
		for _, callories := range lines {
			c, err := strconv.Atoi(callories)
			if err != nil {
				log.Fatalf("failed to parse callories: %v", err)
			}
			fruits = append(fruits, &Fruit{Callories: c})
		}
		deers = append(deers, &Deer{Fruits: fruits})
	}

	var beefiest *Deer
	for _, deer := range deers {
		if beefiest.Callories() < deer.Callories() {
			beefiest = deer
		}
	}
	log.Printf("beefiest deer has %d callories", beefiest.Callories())

	sort.Slice(deers, func(i, j int) bool {
		return deers[i].Callories() < deers[j].Callories()
	})

	//sum of callories of top 3 beefiest deers
	sum := 0
	for i := len(deers) - 1; i >= len(deers)-3; i-- {
		sum += deers[i].Callories()
	}
	log.Printf("sum of callories of top 3 beefiest deers is %d", sum)
}
