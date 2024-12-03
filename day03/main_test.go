package main

import "testing"

func TestGetAnswerNoFilter(t *testing.T) {
	got := part2(&[]string{"do()do()don't()mul(1,1)don't()"})
	want := 0

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
