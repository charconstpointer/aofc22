package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Controller map[int]*Stack[string]

func (c *Controller) Exec(r io.Reader) {
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		row := sc.Text()
		log.Println("Exec", row)
		tokens := strings.Split(row, " ")
		count, _ := strconv.Atoi(string(tokens[1]))
		from, _ := strconv.Atoi(string(tokens[3]))
		to, _ := strconv.Atoi(string(tokens[5]))
		fromStack, ok := (*c)[from]
		if !ok {
			log.Fatalf("No stack for tower %v", from)
		}
		toStack, ok := (*c)[to]
		if !ok {
			log.Fatalf("No stack for tower %v", to)
		}
		log.Printf("Move %v from %v to %v", count, from, to)
		toStack.Put(fromStack.Take(count))
	}
}

type Stack[T any] []T

func (s *Stack[T]) Push(v T) {
	s.Prepend(v)
}

func (s *Stack[T]) Prepend(v T) {
	*s = append(Stack[T]{v}, *s...)
}

func (s *Stack[T]) Take(n int) Stack[T] {
	var min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	n = min(len(*s), n)
	take := make(Stack[T], n)
	copy(take, (*s)[:n])
	log.Println("Take", n, "from", *s, "got", take)
	*s = (*s)[n:]
	return take
}

func (s *Stack[T]) Put(v Stack[T]) {
	test := append(v, *s...)
	log.Println("Put", v, "into", *s, "got", test)
	*s = append(v, *s...)
}

func (s *Stack[T]) Pop() T {
	v := (*s)[0]
	*s = (*s)[1:]
	return v
}

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	sc := bufio.NewScanner(bytes.NewReader(b))
	c := make(Controller)
	for sc.Scan() {
		row := sc.Text()
		for i := 0; i < len(row); i++ {
			if row[i] == ' ' {
				continue
			}
			if row[i] == '[' {
				val := row[i : i+3]
				stack := (i / 4) + 1
				s, ok := c[stack]
				if !ok {
					s = &Stack[string]{}
					c[stack] = s
				}
				*s = append(*s, val)
			}
		}
		if strings.TrimSpace(row) == "" {
			break
		}
	}

	for sc.Scan() {
		c.Exec(strings.NewReader(sc.Text()))
	}

	for k, v := range c {
		log.Println("Stack", k, ":", *v)
	}
}
