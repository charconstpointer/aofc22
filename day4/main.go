package main

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	sc := bufio.NewScanner(bytes.NewReader(b))
	overlapped := 0
	for sc.Scan() {
		row := sc.Text()

		secs := strings.Split(row, ",")
		l, r := secs[0], secs[1]
		var (
			lx, ly int
			rx, ry int
		)
		ls := strings.Split(l, "-")
		rs := strings.Split(r, "-")
		lx, _ = strconv.Atoi(string(ls[0]))
		ly, _ = strconv.Atoi(string(ls[1]))
		rx, _ = strconv.Atoi(string(rs[0]))
		ry, _ = strconv.Atoi(string(rs[1]))
		if rangePartiallyOverlaps(lx, ly, rx, ry) {
			overlapped++
		}
	}
	log.Println(overlapped)
}

func rangeFullyOverlaps(x1, y1, x2, y2 int) bool {
	ok := x1 <= y1 && x2 >= y2 || x1 >= y1 && x2 <= y2
	return ok
}

func rangePartiallyOverlaps(rx, ry, lx, ly int) bool {
	overlap := rx == lx || rx == ly || ry == lx || ry == ly || rx < lx && ry > lx || rx > lx && ry < ly || rx < ly && ry > ly
	return overlap
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
