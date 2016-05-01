package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	SPACE = " "
	AST   = "*"
)

func Rhombus(n int) string {
	result := ""

	var rhombus = make([][]string, 2*n+1)
	for i := range rhombus {
		rhombus[i] = make([]string, 2*n+1)
	}

	var (
		height   = n*2 - 1
		width    = height
		halfline = (width + 1) / 2
	)

	for i := 1; i <= height; i++ {
		numAst := 0
		if i <= halfline {
			numAst = 2*i - 1
		} else {
			numAst = width - 2*(i-halfline)
		}
		spaces := strings.Repeat(SPACE, (width-numAst)/2)
		asts := strings.Repeat(AST, numAst)
		result += spaces + asts + spaces
		if i != height {
			result += "\n"
		}
	}
	return result
}

func main() {
	var n int

	if len(os.Args) == 2 {
		n, _ = strconv.Atoi(os.Args[1])
	} else {
		fmt.Scan(&n)
	}

	fmt.Println(Rhombus(n))
}
