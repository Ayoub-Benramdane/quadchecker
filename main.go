package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func Check(input string, x int, sli []byte) bool {
	y := 1
	c := 0
	for i := 0; i < len(input); i++ {
		c++
		if i < x || i >= len(input)-(x+1) {
			if i == 0 || i == len(input)-(x+1) {
				if i == 0 && input[i] == sli[0] {
					continue
				} else if i == len(input)-(x+1) && input[i] == sli[6] {
					continue
				} else {
					break
				}
			} else if i < ((y-1)*(x+1))+(x-1) && input[i] == sli[1] {
				continue
			} else if i == x-1 || i == len(input)-1 {
				if i == x-1 && input[i] == sli[2] {
					continue
				} else if i == len(input)-1 && input[i] == sli[7] {
					continue
				} else {
					break
				}
			}
		} else if input[i] == '\n' {
			y++
		} else {
			if i == ((y-1)*(x+1)) && input[i] == sli[3] {
				continue
			} else if i < ((y-1)*(x+1))+(x-1) && input[i] == sli[4] {
				continue
			} else if i == ((y-1)*(x+1))+(x-1) && input[i] == sli[5] {
				continue
			} else {
				break
			}
		}

	}
	if c == len(input) {
		return true
	}
	return false
}
func main() {
	var args []byte
	args, _ = ioutil.ReadAll(os.Stdin)
	input := string(args)
	x := 0
	y := 0
	check := 0
	for i := 0; i < len(input); i++ {
		if input[i] == '\n' {
			if y == 0 {
				check = x
			} else {
				if check != x {
					fmt.Println("Not a quad function")
					return
				}
			}
			x = 0
			y++
		} else {
			x++
		}
	}
	x = check
	count := 0
	str1 := []byte{'o', '-', 'o', '|', ' ', '|', 'o', '-', 'o'}
	str2 := []byte{'/', '*', '\\', '*', ' ', '*', '\\', '*', '/'}
	str3 := []byte{'A', 'B', 'A', 'B', ' ', 'B', 'C', 'B', 'C'}
	str4 := []byte{'A', 'B', 'C', 'B', ' ', 'B', 'A', 'B', 'C'}
	str5 := []byte{'A', 'B', 'C', 'B', ' ', 'B', 'C', 'B', 'A'}
	if len(input) >= 1 {
		if input[0] == 'o' {
			if Check(input, x, str1) {
				fmt.Printf("[quadA] [%d] [%d]", x, y)
			} else {
				fmt.Print("Not a quad function")
			}
		} else if input[0] == '/' {
			if Check(input, x, str2) {
				fmt.Printf("[quadB] [%d] [%d]", x, y)
			} else {
				fmt.Print("Not a quad function")
			}
		} else if input[0] == 'A' {
			if Check(input, x, str3) {

				fmt.Printf("[quadC] [%d] [%d]", x, y)
				count++
			}
			if Check(input, x, str4) {
				if count != 0 {
					fmt.Print(" || ")
				}
				fmt.Printf("[quadD] [%d] [%d]", x, y)
				count++
			}
			if Check(input, x, str5) {
				if count != 0 {
					fmt.Print(" || ")
				}
				fmt.Printf("[quadE] [%d] [%d]", x, y)
				count++
			}
			if count == 0 {
				fmt.Print("Not a quad function")
			}
		} else {
			fmt.Printf("Not a quad function")
		}
	} else {
		fmt.Printf("Not a quad function")
	}
	fmt.Printf("\n")
}
