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
		log.Println("Executing", row)
		for i := 0; i < count; i++ {
			pop := fromStack.Pop()
			log.Println("Popping", pop, "from", (*c)[from])
			(*toStack).Push(pop)
			log.Println("Pushing", pop, "to", (*c)[to])
		}
		log.Println("Stacks after execution:", (*c)[from], (*c)[to])
	}
}

type Stack[T any] []T

func (s *Stack[T]) Push(v T) {
	s.Prepend(v)
}

func (s *Stack[T]) Prepend(v T) {
	*s = append(Stack[T]{v}, *s...)
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
				stack := (i % ((len(row) / 4) + 1))
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
