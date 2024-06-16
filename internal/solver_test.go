package solver

import "testing"

func TestSolveGrid(t *testing.T) {
	input := [9][9]int{
		{1, 0, 0, 0, 0, 0, 3, 0, 0},
		{1, 0, 0, 0, 0, 3, 5, 0, 4},
		{0, 0, 0, 5, 0, 6, 0, 0, 0},
		{2, 0, 5, 7, 0, 0, 0, 0, 1},
		{0, 0, 6, 0, 0, 0, 8, 0, 0},
		{3, 0, 0, 0, 0, 1, 4, 0, 6},
		{0, 0, 0, 8, 0, 2, 0, 0, 0},
		{4, 0, 7, 3, 0, 0, 0, 0, 2},
		{0, 0, 8, 0, 0, 0, 0, 0, 3},
	}

	expectedOutOne := [9][9]int{
		{6, 5, 4, 1, 2, 7, 3, 9, 8},
		{1, 7, 2, 9, 8, 3, 5, 6, 4},
		{8, 9, 3, 5, 4, 6, 2, 1, 7},
		{2, 4, 5, 7, 6, 8, 9, 3, 1},
		{7, 1, 6, 4, 3, 9, 8, 2, 5},
		{3, 8, 9, 2, 5, 1, 4, 7, 6},
		{5, 3, 1, 8, 7, 2, 6, 4, 9},
		{4, 6, 7, 3, 9, 5, 1, 8, 2},
		{9, 2, 8, 6, 1, 4, 7, 5, 3},
	}

	expectedOutTwo := true

	actualOne, actualTwo := SolveGrid(input)

	if actualOne != expectedOutOne {
		t.Errorf("error got %v, wanted %v", actualOne, expectedOutOne)
	}
	if actualTwo != expectedOutTwo {
		t.Errorf("error got %v, wanted %v", actualTwo, expectedOutTwo)
	}
}
