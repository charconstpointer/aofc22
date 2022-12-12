package main

import (
	"log"
	"os"
)

func main() {
	b, err := os.ReadFile("sample.txt")
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	msgSize := 14
	offset := 14

	for !isUnique(b[offset-msgSize : offset]) {
		offset++
	}
	log.Println(string(b[offset-msgSize:offset]), offset)
}

func isUnique(b []byte) bool {
	charMap := map[int]bool{}
	for _, v := range b {
		if charMap[int(v)] {
			return false
		}
		charMap[int(v)] = true
	}
	return true
}
