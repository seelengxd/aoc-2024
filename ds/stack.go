package ds

// https://stackoverflow.com/questions/28541609/looking-for-reasonable-stack-implementation-in-golang

import (
	"fmt"
	"log"
)

// Generic stack
// One should implement better constraints
type Stack[V any] []V

func (s *Stack[V]) Push(v V) int {
	*s = append(*s, v)
	return len(*s)
}
func (s *Stack[V]) Last() V {
	l := len(*s)

	// Upto the developer to handle an empty stack
	if l == 0 {
		log.Fatal("Empty Stack")
	}

	last := (*s)[l-1]
	return last
}

func (s *Stack[V]) Pop() V {
	removed := (*s).Last()
	*s = (*s)[:len(*s)-1]

	return removed
}

// Pointer not needed because read-only operation
func (s Stack[V]) Values() {
	for i := len(s) - 1; i >= 0; i-- {
		fmt.Printf("%v ", s[i])
	}
}
