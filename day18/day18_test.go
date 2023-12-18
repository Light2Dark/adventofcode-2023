package day18

import "testing"

type AreaTest struct {
	vertices []Position
	expected int
}

func TestGetVol(t *testing.T) {
	vertices1 := []Position{
		{4, 0}, {4, 2}, {2, 2}, {2, 4}, {0, 4}, {0, 0},
	}

	vertices2 := []Position{
		{0, 0}, {5, 0}, {5, 2}, {1, 2}, {1, 1}, {0, 1},
	}

	var tests = []AreaTest{
		{vertices1, 21},
		{vertices2, 17},
	}

	for _, test := range tests {
		area := getArea(test.vertices)
		perimeter := getPerimeter(test.vertices)
		interior := getInteriorPoints(area, perimeter)
		if output := interior + perimeter; output != test.expected {
			t.Errorf("got %v, expected %v", output, test.expected)
		}
	}
}
