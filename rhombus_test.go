package main

import "testing"

type rhombusTest struct {
	n   int
	out string
}

var rhombusTests = []rhombusTest{
	{1, "*"},
	{2, " * \n***\n * "},
	{3, "  *  \n *** \n*****\n *** \n  *  "},
}

func TestRhombus(t *testing.T) {
	for _, test := range rhombusTests {
		actual := Rhombus(test.n)
		if actual != test.out {
			t.Errorf("got %v, expect %v", actual, test.out)
		}
	}
}
