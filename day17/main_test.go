package main

import "testing"

func TestProgram(t *testing.T) {
	t.Run("C contains 9, [2, 6]", func(t *testing.T) {
		program := makeProgram(0, 0, 9, []int64{2, 6})
		program.run()

		if program.B != 1 {
			t.Errorf("got %d want %d", program.B, 2)
		}
	})

	t.Run("A contains 10, [5,0,5,1,5,4]", func(t *testing.T) {
		program := makeProgram(10, 0, 0, []int64{5, 0, 5, 1, 5, 4})
		result := program.run()
		want := "0,1,2"

		if result != want {
			t.Errorf("got %s want %s", result, want)
		}
	})

	t.Run("A contains 2024, [0,1,5,4,3,0]", func(t *testing.T) {
		program := makeProgram(2024, 0, 0, []int64{0, 1, 5, 4, 3, 0})
		result := program.run()
		want := "4,2,5,6,7,7,7,7,3,1,0"

		if result != want {
			t.Errorf("got %s want %s", result, want)
		}

		if program.A != 0 {
			t.Errorf("got %d want %d", program.A, 0)
		}
	})

	t.Run("B contains 29, [1,7]", func(t *testing.T) {
		program := makeProgram(0, 29, 0, []int64{1, 7})
		program.run()

		if program.B != 26 {
			t.Errorf("got %d want %d", program.B, 26)
		}
	})

	t.Run("B contains 2024, C contains 43690, [4, 0]", func(t *testing.T) {
		program := makeProgram(0, 2024, 43690, []int64{4, 0})
		program.run()

		if program.B != 44354 {
			t.Errorf("got %d want %d", program.B, 44354)
		}
	})

	t.Run("Sample input: A=729, B=0, C=0, [0,1,5,4,3,0]", func(t *testing.T) {
		program := makeProgram(729, 0, 0, []int64{0, 1, 5, 4, 3, 0})
		result := program.run()
		want := "4,6,3,5,6,3,5,2,1,0"

		if result != want {
			t.Errorf("got %s want %s", result, want)
		}
	})

}
