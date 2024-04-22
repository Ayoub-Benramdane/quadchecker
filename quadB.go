package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Print("Not a quad function")
	}
	x, _ := strconv.Atoi(os.Args[1])
	y, _ := strconv.Atoi(os.Args[2])

	QuadB(x, y)
}

func QuadB(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if (j == 0 && i == 0) || (j == x-1 && i == y-1 && i != 0 && j != 0) {
				z01.PrintRune('/')
			} else if (j == 0 && i == y-1) || (j == x-1 && i == 0) {
				z01.PrintRune('\\')
			} else if j == 0 || j == x-1 {
				z01.PrintRune('*')
			} else if i != 0 && i != y-1 {
				z01.PrintRune(' ')
			} else {
				z01.PrintRune('*')
			}
		}
		z01.PrintRune('\n')
	}
}
